package tools

import (
	"context"
	"fmt"

	"github.com/raoptimus/archimate-mcp/client"
)

// ViewConnectionTools provides MCP tool handlers for view connection operations.
type ViewConnectionTools struct {
	client *client.Client
}

// NewViewConnectionTools creates a new ViewConnectionTools instance.
func NewViewConnectionTools(c *client.Client) *ViewConnectionTools {
	return &ViewConnectionTools{client: c}
}

// ViewConnectionListInput represents input for archi_view_connection_list tool.
type ViewConnectionListInput struct {
	// View ID
	ViewID string `json:"view_id"`
}

// List returns all visual connections on a view.
func (t *ViewConnectionTools) List(ctx context.Context, input *ViewConnectionListInput) (*ListResult, error) {
	if input.ViewID == "" {
		return nil, fmt.Errorf("%w: view_id is required", ErrInvalidInput)
	}
	resp, err := t.client.ListViewConnections(ctx, input.ViewID)
	if err != nil {
		return nil, WrapError("view_connection_list", err)
	}

	items := make([]any, len(resp.Items))
	for i, v := range resp.Items {
		items[i] = v
	}
	return &ListResult{Items: items, TotalCount: resp.TotalCount}, nil
}

// ViewConnectionAddInput represents input for archi_view_connection_add tool.
type ViewConnectionAddInput struct {
	// View ID
	ViewID string `json:"view_id"`
	// Relationship ID to visualize
	RelationshipID string `json:"relationship_id"`
	// Source diagram object ID on the view
	SourceObjectID string `json:"source_object_id"`
	// Target diagram object ID on the view
	TargetObjectID string `json:"target_object_id"`
}

// Add creates a visual connection on a view for an existing relationship.
func (t *ViewConnectionTools) Add(ctx context.Context, input *ViewConnectionAddInput) (*client.ViewConnection, error) {
	if input.ViewID == "" || input.RelationshipID == "" || input.SourceObjectID == "" || input.TargetObjectID == "" {
		return nil, fmt.Errorf("%w: view_id, relationship_id, source_object_id, and target_object_id are required", ErrInvalidInput)
	}
	body := map[string]any{
		"relationship_id":  input.RelationshipID,
		"source_object_id": input.SourceObjectID,
		"target_object_id": input.TargetObjectID,
	}
	result, err := t.client.AddViewConnection(ctx, input.ViewID, body)
	if err != nil {
		return nil, WrapError("view_connection_add", err)
	}
	return result, nil
}

// ViewConnectionUpdateInput represents input for archi_view_connection_update tool.
type ViewConnectionUpdateInput struct {
	// View ID
	ViewID string `json:"view_id"`
	// Connection ID
	ConnectionID string `json:"connection_id"`
}

// Update updates a visual connection on a view.
func (t *ViewConnectionTools) Update(ctx context.Context, input *ViewConnectionUpdateInput) (*client.ViewConnection, error) {
	if input.ViewID == "" || input.ConnectionID == "" {
		return nil, fmt.Errorf("%w: view_id and connection_id are required", ErrInvalidInput)
	}
	body := map[string]any{}
	result, err := t.client.UpdateViewConnection(ctx, input.ViewID, input.ConnectionID, body)
	if err != nil {
		return nil, WrapError("view_connection_update", err)
	}
	return result, nil
}

// ViewConnectionRemoveInput represents input for archi_view_connection_remove tool.
type ViewConnectionRemoveInput struct {
	// View ID
	ViewID string `json:"view_id"`
	// Connection ID
	ConnectionID string `json:"connection_id"`
}

// Remove removes a visual connection from a view.
func (t *ViewConnectionTools) Remove(ctx context.Context, input *ViewConnectionRemoveInput) (*client.StatusResponse, error) {
	if input.ViewID == "" || input.ConnectionID == "" {
		return nil, fmt.Errorf("%w: view_id and connection_id are required", ErrInvalidInput)
	}
	result, err := t.client.DeleteViewConnection(ctx, input.ViewID, input.ConnectionID)
	if err != nil {
		return nil, WrapError("view_connection_remove", err)
	}
	return result, nil
}

// ViewConnectionCopyStyleInput represents input for archi_view_connection_copy_style tool.
type ViewConnectionCopyStyleInput struct {
	// View ID
	ViewID string `json:"view_id"`
	// Target connection ID to apply style to
	ConnectionID string `json:"connection_id"`
	// Source connection ID to copy style from
	SourceConnectionID string `json:"source_connection_id"`
}

// CopyStyle copies visual style from a source connection to a target connection.
func (t *ViewConnectionTools) CopyStyle(ctx context.Context, input *ViewConnectionCopyStyleInput) (*client.ViewConnection, error) {
	if input.ViewID == "" || input.ConnectionID == "" || input.SourceConnectionID == "" {
		return nil, fmt.Errorf("%w: view_id, connection_id, and source_connection_id are required", ErrInvalidInput)
	}
	body := map[string]any{"source_connection_id": input.SourceConnectionID}
	result, err := t.client.CopyViewConnectionStyle(ctx, input.ViewID, input.ConnectionID, body)
	if err != nil {
		return nil, WrapError("view_connection_copy_style", err)
	}
	return result, nil
}
