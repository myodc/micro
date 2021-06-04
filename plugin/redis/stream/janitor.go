package stream

import (
	"context"
	"time"

	"github.com/micro/micro/v3/service/logger"
)

var (
	janitorConsumerTimeout = 24 * time.Hour // threshold for an "old" consumer
	janitorFrequency       = 4 * time.Hour  // how often do we run the janitor
)

func (r *redisStream) runJanitor() {
	// Some times it's possible that a consumer group has old consumers that have failed to be deleted.
	// Janitor will clean up any consumers that haven't been seen for X duration
	go func() {
		for {
			if err := r.cleanupConsumers(); err != nil {
				logger.Errorf("Error cleaning up consumers")
			}
			time.Sleep(janitorFrequency)
		}
	}()

}

func (r *redisStream) cleanupConsumers() error {
	now := time.Now()
	keys, err := r.redisClient.Keys(context.Background(), "stream-*").Result()
	if err != nil {
		return err
	}
	for _, streamName := range keys {
		logger.Infof("Cleaning up stream %s", streamName)
		s, err := r.redisClient.XInfoStreamFull(context.Background(), streamName, 1).Result()
		if err != nil {
			logger.Errorf("Error getting info on groups for %s: %s", streamName, err)
			continue
		}
		for _, g := range s.Groups {
			logger.Infof("Cleaning up stream %s group %s", streamName, g.Name)
			for _, c := range g.Consumers {
				// Seen time is the last time this consumer read a message successfully.
				// This means if the stream is low volume you could delete currently connected consumers
				// This isn't a massive problem because the clients should reconnect with a new consumer
				if c.SeenTime.Add(janitorConsumerTimeout).After(now) {
					continue
				}
				logger.Infof("Cleaning up consumer %s, it is %s old", c.Name, time.Since(c.SeenTime))
				if err := r.redisClient.XGroupDelConsumer(context.Background(), streamName, g.Name, c.Name).Err(); err != nil {
					logger.Errorf("Error deleting consumer %s %s %s: %s", streamName, g.Name, c.Name, err)
					continue
				}
			}

		}
	}
	return nil
}
