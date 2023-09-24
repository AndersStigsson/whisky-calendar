package repository

const (
	getUserByID = `
	SELECT 
		users.id,
		users.username,
		users.password,
		users.name
	FROM users
	WHERE users.id = (?)
	`

	getUserByUsername = `
	SELECT 
		users.id,
		users.username,
		users.password,
	FROM users
	WHERE users.username = (?)
	`

	createUser = `
	INSERT INTO users (username, password, name)
	VALUES (?, ?, ?)
	`
)
