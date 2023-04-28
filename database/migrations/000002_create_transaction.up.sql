create table transaction
(
    id               serial,
    telegram_user_id int8 references telegram_user (id) on delete cascade,
    purchase_date    date        not null,
    asset            varchar(10) not null,
    price            float8      not null,
    amount           float8      not null,
    constraint transaction_id primary key (id)
);