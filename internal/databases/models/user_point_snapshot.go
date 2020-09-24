package models

type UserPointSnapshot struct {
	Id          string  `db:"id"`
	UserId      string  `db:"user_id"`
	Point       float64 `db:"point"`
	LastEventId string  `db:"last_event_id"`
	CreatedAt   string  `db:"created_at"`
	CreatedBy   string  `db:"created_by"`
	UpdatedAt   string  `db:"updated_at"`
	UpdatedBy   string  `db:"updated_by"`
}
