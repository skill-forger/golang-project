CREATE TABLE users (
    id       serial primary key,
    username varchar not null,
    password varchar not null,
    email    varchar not null
);