-- name: GetUserCategories :many
select *
from category
where user_id = $1
  and is_expense = $2
order by created_at desc;

-- name: CreateCategory :one
insert into category (user_id, is_expense, label, icon)
values ($1, $2, $3, $4)
returning *;

-- name: UpdateCategory :one
update category
set icon = coalesce($2, icon)
    and label = coalesce($3, label)
    and is_expense = coalesce($4, is_expense)
    and updated_at = $5
where id = $1
returning *;

-- name: DeleteCategory :exec
delete
from category
where id = $1
  and user_id = $2;