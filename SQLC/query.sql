-- En este file estan las query que seran tansformadas en metodos:

-- Esto hara que SQLC genero un metodo llamada GetUserByUsername que resive como argumento username y retorna un solo registro.
-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = $1;

-- name: GetAllUsers :many
SELECT * FROM users;

-- name: AddUser :one
INSERT INTO users (
	username,
	password,
	language
) VALUES ($1, $2, $3)
RETURNING *;
-- Supongo que el returning este hacer que se retorne el user recien creado.

