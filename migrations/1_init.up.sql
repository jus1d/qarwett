create extension if not exists "uuid-ossp";

create table if not exists users (
    id uuid default uuid_generate_v4() not null unique,
    telegram_id bigint not null unique,
    username varchar(32) not null unique,
    firstname varchar(255),
    lastname varchar(255),
    linked_group_id bigint,
    language_code varchar(8),
    created_at timestamp DEFAULT now() NOT NULL
);