package structs

type Product struct {
	Id    int     `json:"-" db:"id"`
	Name  string  `json:"name" db:"name" binding:"required"`
	Price float64 `json:"price" db:"price" binding:"required"`
}
