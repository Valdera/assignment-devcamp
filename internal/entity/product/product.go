package product

import (
	"fmt"
	"time"
)

type Product struct {
	ID          int64     `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Price       float64   `json:"price" db:"price"`
	Description string    `json:"description" db:"description"`
	Variant     string    `json:"string" db:"variant"`
	Discount    float64   `json:"discount" db:"discount"`
	CreatedAt   time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

func (p *Product) Validate() error {
	switch {
	case p.Name == ``:
		return fmt.Errorf(`Name cannot be empty`)
	case p.IsVariant():
		if p.Price == 0 {
			return fmt.Errorf(`Price cannot be 0 or empty`)
		}
	case p.Discount > 100 || p.Discount < 0:
		return fmt.Errorf("Discount must be in range 0 - 100")
	default:
		return nil
	}
	return nil
}

func (p *Product) IsVariant() bool {
	return p.Variant == ``
}
