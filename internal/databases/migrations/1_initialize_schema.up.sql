CREATE TABLE IF NOT EXISTS "users"
(
    "id"         UUID        NOT NULL,
    "name"       TEXT        NOT NULL,
    "email"      TEXT        NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "created_by" VARCHAR(255),
    "updated_at" TIMESTAMPTZ,
    "updated_by" VARCHAR(255),
    CONSTRAINT "uq_users_id" UNIQUE ("id"),
    CONSTRAINT "pk_users" PRIMARY KEY ("id")
);
CREATE INDEX IF NOT EXISTS "ix_users_id" ON "users" USING btree("id");
CREATE INDEX IF NOT EXISTS "ix_email_email" ON "users" USING btree("email");

CREATE TABLE IF NOT EXISTS "user_point_events"
(
    "id"         UUID         NOT NULL,
    "user_id"    UUID         NOT NULL,
    "point"      NUMERIC      NOT NULL,
    "point_type" VARCHAR(100) NOT NULL,
    "notes"      JSONB        NOT NULL,
    "created_at" TIMESTAMPTZ  NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "created_by" VARCHAR(255),
    CONSTRAINT "uq_user_point_events_id" UNIQUE ("id"),
    CONSTRAINT "pk_user_point_events" PRIMARY KEY ("id")
);
CREATE INDEX IF NOT EXISTS "ix_user_point_events_id" ON "user_point_events" USING btree("id");
CREATE INDEX IF NOT EXISTS "ix_user_point_events_user_id" ON "user_point_events" USING btree("user_id");
CREATE INDEX IF NOT EXISTS "ix_user_point_events_balance_type" ON "user_point_events" USING btree("point_type");

CREATE TABLE IF NOT EXISTS "user_point_snapshots"
(
    "id"            UUID        NOT NULL,
    "user_id"       UUID        NOT NULL,
    "point"         NUMERIC     NOT NULL,
    "last_event_id" UUID        NOT NULL,
    "created_at"    TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "created_by"    VARCHAR(255),
    CONSTRAINT "uq_user_point_snapshots_id" UNIQUE ("id"),
    CONSTRAINT "pk_user_point_snapshots_id" PRIMARY KEY ("id")
);
CREATE INDEX IF NOT EXISTS "ix_user_point_snapshots_id" ON "user_point_snapshots" USING btree("id");
CREATE INDEX IF NOT EXISTS "ix_user_point_snapshots_user_id" ON "user_point_snapshots" USING btree("user_id");
CREATE INDEX IF NOT EXISTS "ix_user_point_snapshots_last_event_id" ON "user_point_snapshots" USING btree("last_event_id");
