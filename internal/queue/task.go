package queue

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

type Task struct {
	Args   []byte
	Result []byte

	id   uint64
	name string

	ctx context.Context

	tx   pgx.Tx
	base *pgxpool.Pool

	status string
}

func (task *Task) ID() uint64 {
	return task.id
}

func (task *Task) Name() string {
	return task.name
}

func (task *Task) Tx() pgx.Tx {
	return task.tx
}

func (task *Task) Succeed() {
	task.status = "FINISHED"
}

func (task *Task) Fail() {
	task.status = "FAILURE"
}

func (task *Task) Report(data string) error {
	return task.ReportWith(task.base, "REPORT", data)
}

func (task *Task) ReportWith(base txBase, tag, data string) error {
	return doTx(task.ctx, base, func(tx pgx.Tx) error {
		_, err := tx.Exec(task.ctx, `
		INSERT INTO "ocher_reports" (task_id, tag, message, created_at) VALUES ($1, $2, $3, STATEMENT_TIMESTAMP())
		`, task.id, tag, data)

		if err != nil {
			return errors.WithStack(err)
		}
		return nil
	})
}

func (task *Task) InsertStatus(status string) error {
	return doTx(task.ctx, task.base, func(tx pgx.Tx) error {
		_, err := tx.Exec(task.ctx, `
		INSERT INTO "ocher_statuses" (task_id, status, changed_at) VALUES ($1, $2, STATEMENT_TIMESTAMP())
		`, task.id, status)
		if err != nil {
			return errors.WithStack(err)
		}
		return nil
	})
}

func (task *Task) Commit() error {
	_, err := task.tx.Exec(task.ctx, `UPDATE "ocher_tasks" SET status=$1, result=$2 WHERE id=$3`, task.status, task.Result, task.ID())
	if err != nil {
		return err
	}

	if err := task.tx.Commit(task.ctx); err != nil {
		return errors.Wrap(err, "could not commit transaction")
	}
	return nil
}

func (task *Task) Close() {
	_ = task.tx.Rollback(task.ctx)
}
