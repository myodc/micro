package sql

import (
	"context"

	"github.com/micro/micro/v3/service/model"
)

type sqlModel struct {
	ctx context.Context
	opts model.Options
}

// Context sets the context for the model returning a new copy
func (s *sqlModel) Context(ctx context.Context) model.Model {
	return &sqlModel{
		ctx: ctx,
	}
}

// Register a new model eg. User struct, Order struct
func (s *sqlModel) Register(v interface{}) error {
	return nil
}

// Create a new object. (Maintains indexes set up)
func (s *sqlModel) Create(v interface{}) error {
	return nil
}

// Update will take an existing object and update it.
// TODO: Make use of "sync" interface to lock, read, write, unlock
func (s *sqlModel) Update(v interface{}) error {
	return nil
}

// Read accepts a pointer to a value and expects to fine one or more
// elements. Read throws an error if a value is not found or we can't
// find a matching index for a slice based query.
func (s *sqlModel) Read(query model.Query, resultPointer interface{}) error {
	return nil
}

// Deletes a record. Delete only support Equals("id", value) for now.
// @todo Delete only supports string keys for now.
func (s *sqlModel) Delete(query model.Query) error {
	return nil
}

// NewModel returns a new SQL model
func NewModel(opts ...model.Option) *sqlModel {
	var options model.Options

	for _, o := range opts {
		o(&options)
	}

	return &sqlModel{
		opts: options,
	}
}
