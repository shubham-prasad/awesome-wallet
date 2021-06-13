// Code generated by sqlc. DO NOT EDIT.
// source: user.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  name, pwd
) VALUES (
  $1, $2
) RETURNING id, name, pwd, created_at, updated_at
`

type CreateUserParams struct {
	Name string `json:"name"`
	Pwd  string `json:"pwd"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.queryRow(ctx, q.createUserStmt, createUser, arg.Name, arg.Pwd)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Pwd,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int64) error {
	_, err := q.exec(ctx, q.deleteUserStmt, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, name, pwd, created_at, updated_at FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id int64) (User, error) {
	row := q.queryRow(ctx, q.getUserStmt, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Pwd,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listUser = `-- name: ListUser :many
SELECT id, name, pwd, created_at, updated_at FROM users
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListUserParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListUser(ctx context.Context, arg ListUserParams) ([]User, error) {
	rows, err := q.query(ctx, q.listUserStmt, listUser, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Pwd,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUserName = `-- name: UpdateUserName :one
Update users
SET name = $2, updated_at = now()
WHERE id = $1
RETURNING id, name, pwd, created_at, updated_at
`

type UpdateUserNameParams struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (q *Queries) UpdateUserName(ctx context.Context, arg UpdateUserNameParams) (User, error) {
	row := q.queryRow(ctx, q.updateUserNameStmt, updateUserName, arg.ID, arg.Name)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Pwd,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUserPassword = `-- name: UpdateUserPassword :one
Update users
SET pwd = $2, updated_at = now()
WHERE id = $1
RETURNING id, name, pwd, created_at, updated_at
`

type UpdateUserPasswordParams struct {
	ID  int64  `json:"id"`
	Pwd string `json:"pwd"`
}

func (q *Queries) UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) (User, error) {
	row := q.queryRow(ctx, q.updateUserPasswordStmt, updateUserPassword, arg.ID, arg.Pwd)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Pwd,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
