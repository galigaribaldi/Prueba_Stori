package models

type Transaction struct {
	ID         int `gorm:"primary_key"`
	Month      int
	Day        int
	Date       string
	Transact   float64
	TypeCredit bool
}
