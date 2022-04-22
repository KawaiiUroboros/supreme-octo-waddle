--create table teams with city and string id
create table teams (
   id varchar(255) not null,
   city varchar(255) not null
);
-- create table with city and string external user id and address
create table users (
   external_user_id varchar(255) not null,
   city varchar(255) not null,
   address varchar(255) not null
);