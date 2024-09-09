package constraints

import "github.com/google/uuid"

type ID interface {
	~uint | uuid.UUID
}
