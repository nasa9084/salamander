package context

import (
	"context"
	"database/sql"
	"errors"
)

var (
	Canceled         = context.Canceled
	DeadlineExceeded = context.DeadlineExceeded
)

type Context = context.Context

type CancelFunc = context.CancelFunc

var (
	Background   = context.Background
	TODO         = context.TODO
	WithCancel   = context.WithCancel
	WithDeadline = context.WithDeadline
	WithTimeout  = context.WithTimeout
	WithValue    = context.WithValue
)

type withTxContext struct {
	context.Context
	tx *sql.Tx
}

func (txctx *withTxContext) Tx() *sql.Tx {
	return txctx.tx
}

// WithTx returns a copy ov parent in which the database transaction
func WithTx(parent Context, tx *sql.Tx) Context {
	return &withTxContext{
		Context: parent,
		tx: tx,
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
