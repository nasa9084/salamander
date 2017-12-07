package context

import (
	"context"
	"database/sql"
	"errors"
)

var (
	// Canceled alias
	Canceled = context.Canceled
	// DeadlineExceeded alias
	DeadlineExceeded = context.DeadlineExceeded
)

// Context alias
type Context = context.Context

// CancelFunc alias
type CancelFunc = context.CancelFunc

var (
	// Background alias
	Background = context.Background
	// TODO alias
	TODO = context.TODO
	// WithCancel alias
	WithCancel = context.WithCancel
	// WithDeadline alias
	WithDeadline = context.WithDeadline
	// WithTimeout alias
	WithTimeout = context.WithTimeout
	// WithValue alias
	WithValue = context.WithValue
)

type withTxContext struct {
	context.Context
	tx *sql.Tx
}

// Tx returns transaction associated with the context
func (txctx *withTxContext) Tx() *sql.Tx {
	return txctx.tx
}

// WithTx returns a copy ov parent in which the database transaction
func WithTx(parent Context, tx *sql.Tx) Context {
	return &withTxContext{
		Context: parent,
		tx:      tx,
	}
}

// Tx returns the transaction associated with this context or nil
// if no transaction object is associated.
func Tx(ctx Context) (*sql.Tx, error) {
	if txctx, ok := ctx.(*withTxContext); ok {
		if tx := txctx.Tx(); tx != nil {
			return tx, nil
		}
	}
	return nil, errors.New(`no transaction associated`)
}
