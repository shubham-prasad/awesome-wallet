-- name: CreateWallet :one
INSERT INTO wallets (
  owner, currency, balance 
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetWallet :one
SELECT * FROM wallets
WHERE id = $1 LIMIT 1;

-- name: lockWalletsForUpdate :many
SELECT * FROM wallets
WHERE id = ANY(sqlc.arg(ids)::bigint[]) LIMIT 2
For NO KEY UPDATE;

-- name: ListWallet :many
SELECT * FROM wallets
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateWallet :exec
Update wallets
SET balance = balance + sqlc.arg(amount), updated_at = now()
WHERE id = sqlc.arg(id);

-- name: DeleteWallet :exec
DELETE FROM wallets
WHERE id = $1;