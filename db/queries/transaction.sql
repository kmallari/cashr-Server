-- name: ListAllTransactions :many
SELECT *
FROM transaction
ORDER BY date DESC;

-- name: GetUserTransactions :many
select transaction.id                           as id,
       transaction.user_id                      as user_id,
       amount,
       date,
       category.id                              as category_id,
       description,
       transaction.is_expense                   as is_expense,
       transaction.created_at                   as created_at,
       transaction.updated_at                   as updated_at,
       trim(both from concat(icon, ' ', label)) as category_label
from transaction
         left join category on transaction.category_id = category.id
where transaction.user_id = $1;

-- name: CreateTransaction :one
insert into transaction (user_id, amount, date, is_expense, category_id, description)
values ($1, $2, $3, $4, $5, $6)
returning *;

-- name: UpdateTransaction :one
update transaction
set date        = coalesce($2, date),
    category_id = coalesce($3, category_id),
    description = coalesce($4, description),
    is_expense  = coalesce($5, is_expense),
    amount      = coalesce($6, amount),
    updated_at  = $7
where id = $1
returning *;

-- name: ReplaceTransactionCategory :exec
update transaction
set category_id = $2
where category_id = $1
  and user_id = $3;

-- name: CascadeCategoryDelete :exec
delete
from transaction
where category_id = $1
  and user_id = $2;

-- name: DeleteTransaction :exec
delete
from transaction
where id = $1
  and user_id = $2;

