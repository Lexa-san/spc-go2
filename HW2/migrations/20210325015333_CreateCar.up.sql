CREATE TABLE car (
    car_id bigserial not null primary key,
    mark varchar not null unique,
    max_speed int not null default 0,
    distance int not null default 0,
    handler varchar,
    stock varchar
);