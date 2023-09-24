package repository

const (
	getCommentByID = `
	SELECT 
		comments.id,
		comments.user_id,
		comments.whisky_id,
		comments.content,
		comments.region_id,
		comments.rating
	FROM comments
	WHERE id = (?)
	`

	getCommentsByWhiskyID = `
	SELECT
		comments.id,
		comments.user_id,
		comments.whisky_id,
		comments.content,
		comments.region_id,
		comments.rating
	FROM comments
	WHERE whisky_id = (?)
	`

	storeComment = `
	INSERT INTO comments (user_id, whisky_id, content, region_id, rating)
	VALUES (?, ?, ?, ? ,?)
	`
)
