-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE "votes" (
    "id" serial NOT NULL PRIMARY KEY,
    "option_id"  INTEGER REFERENCES options(id),
    "user_id"  INTEGER REFERENCES users(id),
    "question_id" INTEGER REFERENCES questions(id),
    
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE votes;
