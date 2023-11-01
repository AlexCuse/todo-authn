-- name: GetTODO :one
select * from todos where id = $1 and user_id = $2;

-- name: UpdateTODO :one
update todos
set title = $3,
    description = $4,
    completed = $5,
    updated_at = CURRENT_TIMESTAMP
where id = $1
    and user_id = $2
returning *;

-- name: GetUserTODOs :many
select * from todos where user_id = $1;

-- name: CreateTODO :one
insert into todos (user_id, title, description, completed, created_at)
values ($1, $2, $3, $4, CURRENT_TIMESTAMP)
returning *;

-- name: DeleteTODO :exec
delete from todos where id = $1 and user_id = $2;