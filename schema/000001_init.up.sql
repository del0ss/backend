CREATE TABLE roles
(
    id   serial       not null unique,
    name varchar(255) not null unique
);

CREATE TABLE users
(
    id            serial                                      not null unique,
    email         varchar(255)                                not null unique,
    login         varchar(255)                                not null unique,
    password_hash varchar(255)                                not null,
    role_id       int references roles (id) on delete cascade not null DEFAULT 2
);

CREATE TABLE posts
(
    id      serial                                      not null unique,
    user_id int references users (id) on delete cascade not null,
    title   varchar(255)                                not null,
    content varchar(255)                                not null
);


CREATE TABLE refresh_tokens
(
    id        serial                                      not null unique,
    user_id   int references users (id) on delete cascade not null,
    token     varchar(255)                                not null,
    live_time varchar(255)                                not null
);

INSERT INTO roles VALUES (1, 'admin'), (2, 'user')