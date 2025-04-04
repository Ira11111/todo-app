CREATE TABLE users
(
    id       serial primary key,
    name     varchar(255) unique not null,
    email    varchar(255) unique not null,
    password varchar(255)        not null
);

CREATE TABLE tasks
(
    id          serial primary key,
    title       varchar(255) not null,
    description varchar(1000),
    start_date  date         not null,
    end_data    date         not null
);

CREATE TABLE lists
(
    id          serial primary key,
    title       varchar(255) not null,
    description varchar(255),
    user_id     int          not null,
    FOREIGN KEY (user_id) REFERENCES users (id) on delete cascade
);

CREATE TABLE lists_tasks
(
    id      serial primary key,
    list_id int not null,
    task_id int not null,
    FOREIGN KEY (list_id) REFERENCES lists (id) on delete cascade,
    FOREIGN KEY (task_id) REFERENCES tasks (id) on delete cascade
);