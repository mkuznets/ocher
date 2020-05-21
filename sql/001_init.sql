DROP TYPE IF EXISTS ocher_task_status CASCADE;
CREATE TYPE ocher_task_status AS ENUM (
    'ENQUEUED',
    'EXECUTING',
    'FAILURE',
    'EXPIRED',
    'FINISHED'
    );

DROP TABLE IF EXISTS ocher_tasks CASCADE;
CREATE TABLE IF NOT EXISTS ocher_tasks
(
    id          bigserial PRIMARY KEY,
    enqueued_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    started_at  timestamptz,
    finished_at timestamptz,
    retries     integer              DEFAULT 0,
    status      ocher_task_status    DEFAULT 'ENQUEUED',
    queue       text        NOT NULL,
    message     text                 DEFAULT NULL,
    args        bytea                DEFAULT NULL,
    result      bytea                DEFAULT NULL
);

CREATE INDEX "ocher_tasks__status_idx" ON ocher_tasks USING HASH (status) WHERE status = 'ENQUEUED';
CREATE INDEX "ocher_tasks__queue_idx" ON ocher_tasks USING HASH (queue);

-- -----------------------------------------------------------------------------

DROP TABLE IF EXISTS "ocher_retries" CASCADE;
CREATE TABLE IF NOT EXISTS "ocher_retries"
(
    id         bigserial PRIMARY KEY,
    task_id    bigint      NOT NULL,
    started_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX "ocher_retries__task_id_idx" ON "ocher_retries" (task_id);

-- -----------------------------------------------------------------------------

CREATE OR REPLACE FUNCTION ocher_tasks_notify()
    RETURNS trigger AS
$$
BEGIN
    PERFORM pg_notify(concat('ocher_', regexp_replace(NEW.queue, '[^a-zA-Z0-9_]+', '_', 'g')), NEW.id::text);
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS ocher_notifier ON ocher_tasks;
CREATE TRIGGER ocher_notifier
    AFTER INSERT
    ON ocher_tasks
    FOR EACH ROW
EXECUTE PROCEDURE ocher_tasks_notify();
