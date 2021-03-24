CREATE TABLE "user" (
    user_id bigserial not null primary key,
    login varchar not null unique,
    password varchar not null
);