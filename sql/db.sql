create database if not exists devbook;
use devbook;

drop table if exists users;

create table users (
    id int auto_increment primary key,
    name varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    password varchar(100) not null,
    created_at timestamp default current_timestamp()
) engine=INNODB;