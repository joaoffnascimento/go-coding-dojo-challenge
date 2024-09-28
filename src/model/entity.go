package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Product     string
	Acquired    int
	Avaliable   int
	Safety      int
	Price       float32
	Overbooking bool
}
