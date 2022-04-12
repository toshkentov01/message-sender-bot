package postgres

var (
	// GetMessagesByPrioritySQL ...
	GetMessagesByPrioritySQL = `
		SELECT
			id,
			content,
			priority
		FROM 
			messages
		ORDER BY priority
	`
)