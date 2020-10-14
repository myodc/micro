// Package ci is for continuous integration testing
package ci

import (
	"github.com/micro/go-micro/v3/runtime/local"
	"github.com/micro/micro/v3/service/store/file"
	"github.com/micro/micro/v3/profile"
	"github.com/micro/micro/v3/service/auth/jwt"
	"github.com/micro/micro/v3/service/broker/http"
	"github.com/micro/micro/v3/service/config"
	storeConfig "github.com/micro/micro/v3/service/config/store"
	evStore "github.com/micro/micro/v3/service/events/store"
	memStream "github.com/micro/micro/v3/service/events/stream/memory"
	"github.com/micro/micro/v3/service/logger"
	"github.com/urfave/cli/v2"

	"github.com/micro/micro/plugin/etcd/v3"
	microAuth "github.com/micro/micro/v3/service/auth"
	microEvents "github.com/micro/micro/v3/service/events"
	microRuntime "github.com/micro/micro/v3/service/runtime"
	microStore "github.com/micro/micro/v3/service/store"
)

func init() {
	profile.Register("ci", Profile)
}

// CI profile to use for CI tests
var Profile = &profile.Profile{
	Name: "ci",
	Setup: func(ctx *cli.Context) error {
		microAuth.DefaultAuth = jwt.NewAuth()
		microRuntime.DefaultRuntime = local.NewRuntime()
		microStore.DefaultStore = file.NewStore()
		config.DefaultConfig, _ = storeConfig.NewConfig(microStore.DefaultStore, "")
		microEvents.DefaultStream, _ = memStream.NewStream()
		microEvents.DefaultStore = evStore.NewStore(evStore.WithStore(microStore.DefaultStore))
		profile.SetupBroker(http.NewBroker())
		profile.SetupRegistry(etcd.NewRegistry())
		profile.SetupJWT(ctx)
		profile.SetupConfigSecretKey(ctx)

		var err error
		microStore.DefaultBlobStore, err = file.NewBlobStore()
		if err != nil {
			logger.Fatalf("Error configuring file blob store: %v", err)
		}

		return nil
	},
}
