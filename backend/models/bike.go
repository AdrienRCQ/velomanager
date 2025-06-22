// models/bike.go
package models

import "gorm.io/gorm"

type Bike struct {
	gorm.Model
	Brand     string
	Bikemodel string
	Status    string // ex: "disponible", "réparations", etc.
}
