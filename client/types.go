package client

// Element represents an ArchiMate element.
type Element struct {
	ID            string `json:"id"`
	Type          string `json:"type"`
	Name          string `json:"name"`
	Documentation string `json:"documentation,omitempty"`
	FolderID      string `json:"folder_id,omitempty"`
	FolderName    string `json:"folder_name,omitempty"`
	FolderPath    string `json:"folder_path,omitempty"`
}

// Relationship represents an ArchiMate relationship.
type Relationship struct {
	ID            string `json:"id"`
	Type          string `json:"type"`
	Name          string `json:"name,omitempty"`
	Documentation string `json:"documentation,omitempty"`
	SourceID      string `json:"source_id"`
	TargetID      string `json:"target_id"`
	FolderID      string `json:"folder_id,omitempty"`
	FolderName    string `json:"folder_name,omitempty"`
	FolderPath    string `json:"folder_path,omitempty"`
}

// View represents an ArchiMate view (diagram).
type View struct {
	ID              string `json:"id"`
	Type            string `json:"type"`
	Name            string `json:"name"`
	Documentation   string `json:"documentation,omitempty"`
	FolderID        string `json:"folder_id,omitempty"`
	FolderName      string `json:"folder_name,omitempty"`
	FolderPath      string `json:"folder_path,omitempty"`
	ObjectCount     int    `json:"object_count,omitempty"`
	ConnectionCount int    `json:"connection_count,omitempty"`
}

// Folder represents an ArchiMate model folder.
type Folder struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	ParentID       string `json:"parent_id,omitempty"`
	FolderPath     string `json:"folder_path,omitempty"`
	ElementCount   int    `json:"element_count,omitempty"`
	SubfolderCount int    `json:"subfolder_count,omitempty"`
}

// ObjectStyle represents visual style properties of a view object.
type ObjectStyle struct {
	FillColor     string `json:"fill_color,omitempty"`
	LineColor     string `json:"line_color,omitempty"`
	FontColor     string `json:"font_color,omitempty"`
	LineWidth     *int   `json:"line_width,omitempty"`
	Opacity       *int   `json:"opacity,omitempty"`
	TextAlignment *int   `json:"text_alignment,omitempty"`
}

// ConnectionStyle represents visual style properties of a view connection.
type ConnectionStyle struct {
	LineColor string `json:"line_color,omitempty"`
	LineWidth *int   `json:"line_width,omitempty"`
}

// ViewObject represents an element placed on a view.
type ViewObject struct {
	ID             string       `json:"id"`
	ElementID      string       `json:"element_id,omitempty"`
	Type           string       `json:"type"`
	X              int          `json:"x"`
	Y              int          `json:"y"`
	Width          int          `json:"width"`
	Height         int          `json:"height"`
	ParentObjectID string       `json:"parent_object_id,omitempty"`
	Depth          int          `json:"depth,omitempty"`
	ChildrenCount  int          `json:"children_count,omitempty"`
	Style          *ObjectStyle `json:"style,omitempty"`
}

// ViewConnection represents a visual connection on a view.
type ViewConnection struct {
	ID               string           `json:"id"`
	Type             string           `json:"type"`
	RelationshipID   string           `json:"relationship_id,omitempty"`
	RelationshipType string           `json:"relationship_type,omitempty"`
	SourceObjectID   string           `json:"source_object_id"`
	TargetObjectID   string           `json:"target_object_id"`
	Style            *ConnectionStyle `json:"style,omitempty"`
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
	ID       string `json:"id"`
	Name     string `json:"name"`
	Format   string `json:"format"`
	Width    int    `json:"width,omitempty"`
	Height   int    `json:"height,omitempty"`
	MimeType string `json:"mime_type,omitempty"`
	Data     string `json:"data"` // Base64-encoded image
}

// ErrorResponse represents an error from the jArchi HTTP server.
type ErrorResponse struct {
	Error string `json:"error"`
}
