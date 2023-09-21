package repository

const (
	createWhisky  = ``
	getWhiskyByID = `
	SELECT 
		whiskies.id,
		whiskies.name,
		whiskies.abv,
		whiskies.link,
		whiskies.description,
		whiskies.title,
		whiskies.distillery_id
	FROM whiskies
	WHERE id = (?)
	`
	getAllWhiskies = `
	SELECT 
		whiskies.id,
		whiskies.name,
		whiskies.abv,
		whiskies.link,
		whiskies.description,
		whiskies.title,
		whiskies.distillery_id
	FROM whiskies
	`
)
