package collection

// SortOrder determines if the collection should be sorted [Ascending] or [Descending].
type SortOrder bool

const (
	// Ascending will sort the collection from smallest to biggest elements.
	Ascending SortOrder = true
	// Descending will sort the collection from biggest to smallest elements.
	Descending SortOrder = false
)
