drop database if exists city_v2;
create database if not exists city_v2;
use city_v2;

create table if not exists users
(
    id         bigint unsigned not null primary key auto_increment,
    username   varchar(255)    not null unique,
    password   varchar(255)    not null,
    company_id bigint unsigned null default null,
    branch_id  bigint unsigned null default null,
    icon       varchar(255)         default '',
    fio        varchar(255)         default '',
    Rate       float                default 0,
    phone      varchar(255)         default '',
    pin_code   int                  default 0,
    active     boolean              default true,
    created_at timestamp       null default null,
    updated_at timestamp       null default null,
    deleted_at timestamp       null default null,
    unique (company_id, pin_code)
);

create table if not exists companies
(
    id            bigint unsigned not null primary key auto_increment,
    user_id       bigint unsigned null default null,
    name          varchar(255) unique,
    inn           int(9)               default 0,
    fio           varchar(255)         default '',
    legal_address varchar(255)         default '',
    address       varchar(255)         default '',
    description   varchar(255)         default '',
    icon          varchar(255)         default '',
    x             decimal(10, 9)       default 0,
    y             decimal(10, 9)       default 0,
    site          varchar(255)         default '',
    mail          varchar(255)         default '',
    phone1        varchar(255)         default '',
    phone2        varchar(255)         default '',
    active        boolean              default true,
    created_at    timestamp       null default null,
    updated_at    timestamp       null default null,
    deleted_at    timestamp       null default null
);

create table if not exists branches
(
    id            bigint unsigned not null primary key auto_increment,
    user_id       bigint unsigned null default null,
    company_id    bigint unsigned null default null,
    name          varchar(255)         default '',
    address       varchar(255)         default '',
    x             decimal(10, 9)       default 0,
    y             decimal(10, 9)       default 0,
    phone1        varchar(255)         default '',
    phone2        varchar(255)         default '',
    legal_address varchar(255)         default '',
    photo         varchar(255)         default '',
    active        boolean              default true,
    until_date    timestamp       null default null,
    created_at    timestamp       null default null,
    updated_at    timestamp       null default null,
    deleted_at    timestamp       null default null
);

create table if not exists roles
(
    id   bigint unsigned     not null primary key auto_increment,
    name varchar(255) unique not null
);

create table if not exists user_roles
(
    user_id bigint unsigned not null,
    role_id bigint unsigned not null,
    unique (user_id, role_id)
);

create table if not exists user_logs
(
    id         bigint unsigned not null primary key auto_increment,
    user_id    bigint          not null,
    data       text            not null,
    created_at timestamp       null default null
);

insert into roles (name)
values ('admin'),
       ('director'),
       ('manager'),
       ('hostess'),
       ('chef'),
       ('courier'),
       ('warehouse'),
       ('waiter'),
       ('user');


alter table users
    add foreign key (company_id) references companies (id),
    add foreign key (branch_id) references branches (id);

alter table companies
    add foreign key (user_id) references users (id);

alter table branches
    add foreign key (user_id) references users (id),
    add foreign key (company_id) references companies (id);

alter table user_roles
    add foreign key (user_id) references users (id),
    add foreign key (role_id) references roles (id);