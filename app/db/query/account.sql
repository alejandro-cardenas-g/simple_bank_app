-- name: CreateAccount :one
INSERT INTO "public"."accounts" (
  owner, balance, currency
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetAccount :one
SELECT * FROM "public"."accounts"
WHERE id = $1 LIMIT 1;

-- name: GetAccountForUpdate :one
SELECT * FROM "public"."accounts"
WHERE id = $1 LIMIT 1 FOR NO KEY UPDATE;

-- name: ListAccounts :many
SELECT * FROM "public"."accounts"
WHERE owner = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateAccount :one
UPDATE "public"."accounts"
  set balance = $2
WHERE id = $1
RETURNING *;

-- name: AddAccountBalance :one
UPDATE "public"."accounts"
  set balance = balance + sqlc.arg(amount)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM "public"."accounts"
WHERE id = $1;