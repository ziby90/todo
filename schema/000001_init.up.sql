CREATE TABLE users
(
    id            serial       not null unique,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null,
    created         timestamp  not null default NOW()
);

CREATE TABLE todo_lists
(
    id          serial       not null unique,
    title       varchar(255) not null,
    description varchar(255),
    user_id     int references users (id) on delete cascade      not null,
    created         timestamp  not null default NOW()
);
--
-- CREATE TABLE users_lists
-- (
--     id      serial                                           not null unique,
--     user_id int references users (id) on delete cascade      not null,
--     list_id int references todo_lists (id) on delete cascade not null,
--     created         timestamp  not null default NOW()
-- );

CREATE TABLE todo_items
(
    id          serial       not null unique,
    title       varchar(255) not null,
    description varchar(255),
    list_id int references todo_lists (id) on delete cascade not null,
    done        boolean      not null default false,
    created         timestamp  not null default NOW()
);

--
-- CREATE TABLE lists_items
-- (
--     id      serial                                           not null unique,
--     item_id int references todo_items (id) on delete cascade not null,
--     list_id int references todo_lists (id) on delete cascade not null,
--     created         timestamp  not null default NOW()
-- );