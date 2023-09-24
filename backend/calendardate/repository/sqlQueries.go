package repository

const (
	getDateByDayOfMonth = `
	SELECT 
		calendar_dates.day_of_month,
		calendar_dates.reveal_date,
		calendar_dates.whisky_id
	FROM calendar_dates
	WHERE calendar_dates.day_of_month = (?)
	`
	getAllCalendarDates = `
	SELECT 
		calendar_dates.day_of_month,
		calendar_dates.reveal_date,
		calendar_dates.whisky_id
	FROM calendar_dates
	`
)
