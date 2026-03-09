package client

// Element represents an ArchiMate element.
type Element struct {
	ID            string `json:"id"`
	Type          string `json:"type"`
	Name          string `json:"name"`
	Documentation string `json:"documentation,omitempty"`
}

// Relationship represents an ArchiMate relationship.
type Relationship struct {
	ID            string `json:"id"`
	Type          string `json:"type"`
	Name          string `json:"name,omitempty"`
	Documentation string `json:"documentation,omitempty"`
	SourceID      string `json:"source_id"`
	TargetID      string `json:"target_id"`
}

// View represents an ArchiMate view (diagram).
type View struct {
	ID            string `json:"id"`
	Type          string `json:"type"`
	Name          string `json:"name"`
	Documentation string `json:"documentation,omitempty"`
}

// ViewObject represents an element placed on a view.
type ViewObject struct {
	ID        string `json:"id"`
	ElementID string `json:"element_id,omitempty"`
	Type      string `json:"type"`
	X         int    `json:"x"`
	Y         int    `json:"y"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
}

// ModelInfo represents model metadata.
type ModelInfo struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Purpose string `json:"purpose,omitempty"`
}

// Property represents an element/relationship property.
type Property struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// ListResponse is a generic list response from the jArchi HTTP server.
type ListResponse[T any] struct {
	Items      []T `json:"items"`
	TotalCount int `json:"total_count"`
}

// StatusResponse represents a status-only response.
type StatusResponse struct {
	Status string `json:"status"`
	ID     string `json:"id,omitempty"`
	Note   string `json:"note,omitempty"`
}

// ViewImageExport represents an exported view image.
type ViewImageExport struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Format string `json:"format"`
	Data   string `json:"data"` // Base64-encoded PNG
}

// ErrorResponse represents an error from the jArchi HTTP server.
type ErrorResponse struct {
	Error string `json:"error"`
}
