package transactions

import "time"

//namanya harusnya order
type Order struct {
	ID         int
	UserID     int
	ProductID  int
	Quantity   int
	TotalPrice int
	Status     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

//namanya harusnya transaction
type Payment struct {
	ID            int
	TransactionID int
	Status        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

const (
	Pocessing = "processing"
	Succes    = "succes"
	Failed    = "failed"
)
