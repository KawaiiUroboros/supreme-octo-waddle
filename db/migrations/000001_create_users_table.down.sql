BEGIN;
--if exist drop table users
DROP TABLE IF EXISTS users CASCADE;
--if exist drop table schema_migrations
DROP TABLE IF EXISTS schema_migrations CASCADE;
COMMIT;