package utils

// Paginator is used to do pagination.
type Paginator struct {
	Offset uint64
	Limit  uint64
}

// NewPaginator creates a new paginator.
func NewPaginator(offset, limit uint64) *Paginator {
	return &Paginator{Offset: offset, Limit: limit}
}
