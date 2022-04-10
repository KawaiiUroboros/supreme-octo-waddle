BEGIN;
create table users
(
    id                    serial primary key,
    external_user_id      varchar(255) not null,
    begin_date            timestamp not null,
    channel_id            varchar(60) not null,
    notification_interval integer not null ,
    last_confirmation     timestamp,
    is_deleted            boolean not null
);
--add unique constraint on channel_id
ALTER TABLE users ADD CONSTRAINT users_channel_id_key UNIQUE (channel_id);
COMMIT;