CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
    id uuid DEFAULT uuid_generate_v4() NOT NULL UNIQUE ,
    telegram_id bigint NOT NULL UNIQUE,
    username varchar(32) NOT NULL UNIQUE,
    firstname varchar(255) DEFAULT '' NOT NULL,
    lastname varchar(255) DEFAULT '' NOT NULL,
    stage int DEFAULT 0 NOT NULL,
    linked_group_id bigint DEFAULT 0 NOT NULL,
    linked_group_title varchar(255) DEFAULT '' NOT NULL,
    language_code varchar(8) NOT NULL,
    is_admin bool DEFAULT false not null,
    created_at timestamp DEFAULT now() NOT NULL
);