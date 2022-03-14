package valueobject

import (
	"time"

	"github.com/gofrs/uuid"
)

// Transaction is a payment between two parties
type Transaction struct {
	// all values lowercase since they are immutable
	amount    int
	from      uuid.UUID
	to        uuid.UUID
	createdAt time.Time
}
