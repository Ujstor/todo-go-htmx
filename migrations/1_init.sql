-- +goose Up
CREATE TABLE "todos" (
    "id" SERIAL NOT NULL,
    "todo" VARCHAR(255),
    "done" BOOLEAN,
    PRIMARY KEY ("id")
);

-- +goose Down
DROP TABLE "todos";

