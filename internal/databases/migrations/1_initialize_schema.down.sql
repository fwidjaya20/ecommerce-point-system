DROP TABLE IF EXISTS "users";
DROP INDEX IF EXISTS "ix_users_id";
DROP INDEX IF EXISTS "ix_users_email";

DROP TABLE IF EXISTS "user_point_events";
DROP INDEX IF EXISTS "ix_user_point_events_id";
DROP INDEX IF EXISTS "ix_user_point_events_user_id";

DROP TABLE IF EXISTS "user_point_snapshots";
DROP INDEX IF EXISTS "ix_user_point_snapshots_id";
DROP INDEX IF EXISTS "ix_user_point_snapshots_user_id";
DROP INDEX IF EXISTS "ix_user_point_snapshots_last_event_id";