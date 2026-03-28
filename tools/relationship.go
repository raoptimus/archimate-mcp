package tools

import (
	"context"
	"fmt"

	"github.com/raoptimus/archimate-mcp/client"
)

// RelationshipTools provides MCP tool handlers for ArchiMate relationship operations.
type RelationshipTools struct {
	client *client.Client
}

// NewRelationshipTools creates a new RelationshipTools instance.
func NewRelationshipTools(c *client.Client) *RelationshipTools {
	return &RelationshipTools{client: c}
}

// RelationshipListInput represents input for archi_relationship_list tool.
type RelationshipListInput struct {
	// Relationship type (e.g., "serving-relationship")
	Type string `json:"type,omitempty"`
	// Filter by source element ID
	SourceID string `json:"source_id,omitempty"`
	// Filter by target element ID
	TargetID string `json:"target_id,omitempty"`
}

// List returns a list of ArchiMate relationships.
func (t *RelationshipTools) List(ctx context.Context, input *RelationshipListInput) (*ListResult, error) {
	resp, err := t.client.ListRelationships(ctx, input.Type, input.SourceID, input.TargetID)
	if err != nil {
		return nil, WrapError("relationship_list", err)
	}

	items := make([]any, len(resp.Items))
	for i, v := range resp.Items {
		items[i] = v
	}
	return &ListResult{Items: items, TotalCount: resp.TotalCount}, nil
}

// RelationshipGetInput represents input for archi_relationship_get tool.
type RelationshipGetInput struct {
	ID string `json:"id"`
}

// Get returns an ArchiMate relationship by ID.
func (t *RelationshipTools) Get(ctx context.Context, input *RelationshipGetInput) (*client.Relationship, error) {
	if input.ID == "" {
		return nil, fmt.Errorf("%w: id is required", ErrInvalidInput)
	}
	result, err := t.client.GetRelationship(ctx, input.ID)
	if err != nil {
		return nil, WrapError("relationship_get", err)
	}
	return result, nil
}

// RelationshipCreateInput represents input for archi_relationship_create tool.
type RelationshipCreateInput struct {
	// Relationship type (e.g., "serving-relationship")
	Type string `json:"type"`
	// Source element ID
	SourceID string `json:"source_id"`
	// Target element ID
	TargetID string `json:"target_id"`
	// Optional name
	Name string `json:"name,omitempty"`
	// Optional documentation
	Documentation string `json:"documentation,omitempty"`
	// Optional folder ID to place the relationship in
	FolderID string `json:"folder_id,omitempty"`
}

// Create creates a new ArchiMate relationship.
func (t *RelationshipTools) Create(ctx context.Context, input *RelationshipCreateInput) (*client.Relationship, error) {
	if input.Type == "" || input.SourceID == "" || input.TargetID == "" {
		return nil, fmt.Errorf("%w: type, source_id, and target_id are required", ErrInvalidInput)
	}
	body := map[string]any{
		"type":      input.Type,
		"source_id": input.SourceID,
		"target_id": input.TargetID,
	}
	if input.Name != "" {
		body["name"] = input.Name
	}
	if input.Documentation != "" {
		body["documentation"] = input.Documentation
	}
	if input.FolderID != "" {
		body["folder_id"] = input.FolderID
	}
	result, err := t.client.CreateRelationship(ctx, body)
	if err != nil {
		return nil, WrapError("relationship_create", err)
	}
	return result, nil
}

// RelationshipUpdateInput represents input for archi_relationship_update tool.
type RelationshipUpdateInput struct {
	ID            string `json:"id"`
	Name          string `json:"name,omitempty"`
	Documentation string `json:"documentation,omitempty"`
}

// Update updates an ArchiMate relationship.
func (t *RelationshipTools) Update(ctx context.Context, input *RelationshipUpdateInput) (*client.Relationship, error) {
	if input.ID == "" {
		return nil, fmt.Errorf("%w: id is required", ErrInvalidInput)
	}
	body := map[string]any{}
	if input.Name != "" {
		body["name"] = input.Name
	}
	if input.Documentation != "" {
		body["documentation"] = input.Documentation
	}
	result, err := t.client.UpdateRelationship(ctx, input.ID, body)
	if err != nil {
		return nil, WrapError("relationship_update", err)
	}
	return result, nil
}

// RelationshipDeleteInput represents input for archi_relationship_delete tool.
type RelationshipDeleteInput struct {
	ID string `json:"id"`
}

// Delete deletes an ArchiMate relationship.
func (t *RelationshipTools) Delete(ctx context.Context, input *RelationshipDeleteInput) (*client.StatusResponse, error) {
	if input.ID == "" {
		return nil, fmt.Errorf("%w: id is required", ErrInvalidInput)
	}
	result, err := t.client.DeleteRelationship(ctx, input.ID)
	if err != nil {
		return nil, WrapError("relationship_delete", err)
	}
	return result, nil
}

// RelationshipMoveInput represents input for archi_relationship_move tool.
type RelationshipMoveInput struct {
	// Relationship ID
	ID string `json:"id"`
	// Target folder ID
	FolderID string `json:"folder_id"`
}

// Move moves a relationship to a different folder.
func (t *RelationshipTools) Move(ctx context.Context, input *RelationshipMoveInput) (*client.Relationship, error) {
	if input.ID == "" || input.FolderID == "" {
		return nil, fmt.Errorf("%w: id and folder_id are required", ErrInvalidInput)
	}
	body := map[string]any{"folder_id": input.FolderID}
	result, err := t.client.MoveRelationship(ctx, input.ID, body)
	if err != nil {
		return nil, WrapError("relationship_move", err)
	}
	return result, nil
}
