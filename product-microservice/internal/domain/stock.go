package domain

import "time"

type Stock struct {
	ID        int64
	ProductID int64
	Quantity  int64
	UpdatedAt time.Time
}

func (s *Stock) GenerateUpdatedAt() {
	s.UpdatedAt = time.Now()
}
