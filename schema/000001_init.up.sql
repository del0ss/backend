CREATE TABLE roles
(
    id   serial       not null unique,
    name varchar(255) not null unique
);

CREATE TABLE category -- change to categories
(
    id   serial       not null unique PRIMARY KEY,
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

CREATE TABLE pizzas
(
    id          serial       not null unique PRIMARY KEY,
    image_url   varchar(255) not null,
    name        varchar(255) not null,
    types       int[] not null,
    sizes       int[] not null,
    price       int          not null,
    category_id int          not null references category (id) ON DELETE CASCADE,
    rating      int          not null
);


CREATE TABLE refresh_tokens
(
    id        serial                                      not null unique,
    user_id   int references users (id) on delete cascade not null,
    token     varchar(255)                                not null,
    live_time varchar(255)                                not null
);

INSERT INTO roles
VALUES (1, 'admin'),
       (2, 'user');

INSERT INTO category
VALUES (1, 'Мясные'),
       (2, 'Вегетарианская'),
       (3, 'Гриль'),
       (4, 'Острые'),
       (5, 'Закрытые');