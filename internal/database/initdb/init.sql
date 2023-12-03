CREATE TABLE todos (
    id serial not null,
    todo varchar(255),
    done boolean,
    primary key(id)
);

INSERT INTO todos(todo, done) VALUES ('Hello', false);