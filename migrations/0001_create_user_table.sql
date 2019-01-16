-- +goose Up
-- SQL in this section is executed when the migration is applied.:wqw
CREATE TABLE "users"
(
    "id" serial NOT NULL PRIMARY KEY,
    "first_name" varchar (200) NOT NULL,
    "last_name" varchar (200) NOT NULL,
    "email" varchar(500) NOT NULL,
    "phone" varchar(50) NOT NULL,
    "password_digest" varchar(500) NOT NULL,

    "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP SCHEMA public CASCADE;
CREATE SCHEMA public;
