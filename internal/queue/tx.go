package queue

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

type txBase interface {
	Begin(ctx context.Context) (pgx.Tx, error)
}

func doTx(ctx context.Context, base txBase, op func(tx pgx.Tx) error) error {
	tx, err := base.Begin(ctx)

	if err != nil {
		return errors.Wrap(err, "could not start transaction")
	}
	//noinspection GoUnhandledErrorResult
	defer tx.Rollback(ctx) // nolint

	err = op(tx)
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return errors.Wrap(err, "could not commit transaction")
	}
	return nil
}
