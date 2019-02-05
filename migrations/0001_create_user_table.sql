-- +goose Up
-- SQL in this section is executed when the migration is applied.:wqw
CREATE TABLE "users"
(
    "id" serial NOT NULL PRIMARY KEY,
    "username" varchar(500) NOT NULL,
    "password" varchar(500) NOT NULL,
    "token" varchar(500),
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP SCHEMA public CASCADE;
CREATE SCHEMA public;
