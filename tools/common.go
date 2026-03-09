package tools

// ListResult wraps list response with metadata.
type ListResult struct {
	Items      []any `json:"items"`
	TotalCount int   `json:"total_count,omitempty"`
}
