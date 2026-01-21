package domain

import "time"

type Stock struct {
	ID        int64
	Name      string
	ProductID int64
	Quantity  int64
	UpdatedAt time.Time
}
