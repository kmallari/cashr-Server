drop table if exists transaction, category;

create table category
(
    id         char(36) primary key default gen_random_uuid(),
    user_id    char(36)    not null,
    label      varchar(24) not null,
    icon       char(1),
    is_expense bool        not null,
    created_at bigint      not null default date_part('epoch'::text, now()),
    updated_at bigint      not null default date_part('epoch'::text, now()),
    unique (is_expense, user_id, label),
    unique (is_expense, user_id, label, icon)
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

drop function if exists update_updated_at_column;
create function update_updated_at_column() returns trigger
    language plpgsql as
$$
begin
    new.updated_at = date_part('epoch'::text, now()); return new;
end;
$$;

drop trigger if exists update_category_updated_at on category;
create trigger update_category_updated_at
    after update
    on category
    for each row
execute procedure update_updated_at_column();

drop trigger if exists update_transaction_updated_at on transaction;
create trigger update_transaction_updated_at
    after update
    on transaction
    for each row
execute procedure update_updated_at_column();


drop index if exists user_id_and_date_idx, category_user_id;
create index transaction_user_id_and_date_idx on transaction (user_id, date);
create index transaction_user_id_and_category on transaction (user_id, category_id);
create index category_user_id on category (user_id);