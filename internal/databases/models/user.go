package models

type User struct {
	Id        string `db:"id"`
	Name      string `db:"name"`
	Email     string `db:"email"`
	CreatedAt string `db:"created_at"`
	CreatedBy string `db:"created_by"`
	UpdatedAt string `db:"updated_at"`
	UpdatedBy string `db:"updated_by"`
}