create table users(
    user_id serial not null unique,
    name varchar(256) not null,
    username varchar(256) not null unique,
    password_hash varchar(256) not null
);

create table todo_lists(
    todo_list_id serial not null unique,
    title varchar(256) not null,
    description varchar(256)
);

create table users_lists(
    users_list_id serial not null unique,
    user_id int references users (user_id) on delete cascade not null,
    todo_list_id int references todo_lists (todo_list_id) on delete cascade not null
);

create table todo_items(
    todo_item_id serial not null unique,
    title varchar(256) not null,
    description varchar(256),
    done boolean not null default false
);

create table lists_items(
    lists_item_id serial not null unique,
    todo_item_id int references todo_items (todo_item_id) on delete cascade not null,
    todo_list_id int references todo_lists (todo_list_id) on delete cascade not null
);
