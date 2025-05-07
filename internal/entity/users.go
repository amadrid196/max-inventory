package entity

type Users struct {
	ID         int64  `db:"id"`
	NAME       string `db:"name"`
	EMAIL      string `db:"email"`
	PASSWORD   string `db:"password"`
	created_at string `db:"created_at"`
}
