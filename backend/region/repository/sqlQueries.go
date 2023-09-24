package repository

const (
	getRegionByID = `
	SELECT 
		regions.id,
		regions.name
	FROM regions
	WHERE id = (?)
	`
)
