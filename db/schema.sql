drop table if exists transaction, category;

create table category
(
    id         char(36)             default gen_random_uuid(),
    user_id    char(36)    not null,
    label      varchar(24) not null,
    icon       varchar(60),
    is_expense bool        not null,
    created_at bigint      not null default date_part('epoch'::text, now()),
    updated_at bigint      not null default date_part('epoch'::text, now()),
    last_used  bigint      not null default date_part('epoch'::text, now()),
    unique (is_expense, user_id, label),
    unique (is_expense, user_id, label, icon),
    primary key (id)
);

create table transaction
(
    id          char(36)          default gen_random_uuid(),
    user_id     char(36) not null,
    amount      decimal  not null,
    date        bigint   not null,
    category_id char(36) not null,
    description varchar(250),
    is_expense  bool     not null,
    created_at  bigint   not null default date_part('epoch'::text, now()),
    updated_at  bigint   not null default date_part('epoch'::text, now()),
    primary key (id),
    foreign key (category_id) references category (id)
);

drop index if exists user_id_and_date_idx, category_user_id;
create index transaction_user_id_and_date_idx on transaction (user_id, date);
create index transaction_user_id_and_category on transaction (user_id, category_id);
create index category_user_id on category (user_id);