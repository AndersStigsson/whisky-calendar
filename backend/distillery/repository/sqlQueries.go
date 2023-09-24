package repository

const (
	getDistilleryByID = `
	SELECT 
		distilleries.id,
		distilleries.name,
		distilleries.region_id
	FROM distilleries
	WHERE id = (?)
	`
)
