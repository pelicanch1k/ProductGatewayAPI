package product

type Product struct {
	Id    int     `json:"id" db:"id"`
	Name  string  `json:"name" db:"name" binding:"required"`
	Price float64 `json:"price" db:"price" binding:"required"`
}
