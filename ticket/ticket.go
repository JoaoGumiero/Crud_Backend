package ticket

import "time"

type Status int

// Go don't have the Enum so you use the iota instead
const (
	_          = iota // Initializes with 0
	Open              // 1
	InProgress        // 2
	Solved            // 3
	Closed            // 4
)

// A struct of a Ticket
type Ticket struct {
	ID            string    `json:"id"`
	Title         string    `json:"title"`
	Analysis_Date time.Time `json:"analysisDate"`
	Solving_Date  time.Time `json:"solvingDate"`
	Description   string    `json:"author"`
	SenderQueue   string    `json:"senderQueue"`
	RecieverQueue string    `json:"recieverQueue"`
	Status        Status    `json:"status"`
}
