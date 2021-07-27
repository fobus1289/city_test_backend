create table if not exists categories
(
    id         bigint primary key auto_increment not null,
    name       varchar(255)                      not null,
    company_id bigint unsigned default null,
    active     boolean                           not null
);

create table if not exists products
(
    id          bigint primary key auto_increment not null,
    name        varchar(255)                      not null,
    photo       varchar(255)    default 'no-img.png',
    price       float                             not null,
    description text            default '',
    where_run   enum ('bar','kitchen')            not null,
    company_id  bigint unsigned default null,
    branch_id   bigint unsigned default null,
    active      boolean,
    created_at  timestamp       default current_timestamp
);

create table if not exists category_components
(
    id          bigint primary key auto_increment not null,
    name        varchar(255)                      not null,
    sell_price  float                             not null,
    count       float                             not null,
    units       enum ('gram','piece')             not null,
    photo       varchar(255)    default 'no-img.png',
    description text            default '',
    category_id bigint unsigned default null,
    company_id  bigint unsigned default null,
    branch_id   bigint unsigned default null,
    active      boolean,
    created_at  timestamp       default current_timestamp
);

create table if not exists components
(
    id                    bigint primary key auto_increment not null,
    category_component_id bigint                            not null,
    name                  varchar(255)                      not null,
    units                 enum ('gram','piece')             not null,
    buy_price             float                             not null,
    company_id            bigint unsigned default null,
    branch_id             bigint unsigned default null,
    active                boolean,
    created_at            timestamp       default current_timestamp
);

create table if not exists product_components
(
    id                   bigint primary key auto_increment not null,
    product_id           bigint unsigned                   not null,
    component_id         bigint unsigned                   not null,
    can_add              boolean default false,
    can_remove           boolean default false,
    changed_price_add    boolean default false,
    changed_price_remove boolean default false,
    unique (product_id, component_id)
);

create table if not exists orders
(
    id              bigint primary key auto_increment                              not null,
    price           float                                                          not null,
    status          enum ('created','inProcessing','canceled','ready','delivered') not null,
    number          int                                                            not null,
    reason_canceled varchar(255) default '',
    description     varchar(255)                                                   not null,
    branch_id       bigint                                                         not null,
    customer_id     bigint                                                         not null,
    worker_id       bigint                                                         not null,
    created_at      timestamp    default current_timestamp
);

create table if not exists realizations
(
    id              bigint primary key auto_increment                              not null,
    order_id        bigint                                                         not null,
    price           float                                                          not null,
    product_name    varchar(255)                                                   not null,
    reason_canceled varchar(255) default '',
    branch_id       bigint                                                         not null,
    product_id      bigint                                                         not null,
    components      varchar(255)                                                   not null,
    where_run       enum ('bar','kitchen')                                         not null,
    status          enum ('created','inProcessing','canceled','ready','delivered') not null,
    created_at      timestamp    default current_timestamp
);

