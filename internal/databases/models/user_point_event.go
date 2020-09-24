package models

type UserPointEvent struct {
	Id        string  `db:"id"`
	UserId    string  `db:"user_id"`
	Point     float64 `db:"point"`
	PointType string  `db:"point_type"`
	Notes     string  `db:"notes"`
	CreatedAt string  `db:"created_at"`
	CreatedBy string  `db:"created_by"`
	UpdatedAt string  `db:"updated_at"`
	UpdatedBy string  `db:"updated_by"`
}
