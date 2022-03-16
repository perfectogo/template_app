create table users(
    id            serial       not null unique,
    name          varchar(256) not null,
    username      varchar(256) not null unique,
    password_hash varchar(256) not null
);

create table todo_lists(
    id          serial       not null unique,
    title       varchar(256) not null,
    description varchar(256)
);

create table users_lists(
    id      serial                                           not null unique,
    user_id int references users (id) on delete cascade      not null,
    list_id int references todo_lists (id) on delete cascade not null
);

create table todo_items(
    id          serial       not null unique,
    title       varchar(256) not null,
    description varchar(256),
    done        boolean      not null default false
);


create table lists_items(
    id      serial                                           not null unique,
    item_id int references todo_items (id) on delete cascade not null,
    list_id int references todo_lists (id) on delete cascade not null
);
