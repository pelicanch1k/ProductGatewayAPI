package structs

type User struct {
	Id       int    `json:"-"`
	Username string `json:"username" db:"username" binding:"required"`
	Password string `json:"password" db:"password_hash" binding:"required"`
}
