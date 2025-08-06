create database max_inventory;
use max_inventory;


-- Table: inventory

create table users (
    id int auto_increment not null,
    name varchar(50) not null,
    email varchar(50) not null unique,
    password varchar(255) not null,
    created_at timestamp default current_timestamp,
    primary key (id)
);

create table products (
    id int auto_increment not null,
    name varchar(50) not null,
    description text,
    price float not null,
    stock int not null,
    create_by int not null,
    created_at timestamp default current_timestamp,
    primary key (id)
    foreign key (create_by) references users(id) on delete cascade on update cascade
);

create table roles (
    id int auto_increment not null,
    name varchar(50) not null,
    created_at timestamp default current_timestamp,
    primary key (id)
);

create table user_roles (
    id int auto_increment not null,
    user_id int not null,
    role_id int not null,
    created_at timestamp default current_timestamp,
    primary key (id),
    foreign key (user_id) references users(id) on delete cascade on update cascade,
    foreign key (role_id) references roles(id) on delete cascade on update cascade
);

insert into roles (id, name) values (1, 'admin'), (2, 'customer'), (3, 'seller');

