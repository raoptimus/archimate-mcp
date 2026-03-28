package tools

import (
	"context"
	"fmt"

	"github.com/raoptimus/archimate-mcp/client"
)

// ElementTools provides MCP tool handlers for ArchiMate element operations.
type ElementTools struct {
	client *client.Client
}

// NewElementTools creates a new ElementTools instance.
func NewElementTools(c *client.Client) *ElementTools {
	return &ElementTools{client: c}
}

// ElementListInput represents input for archi_element_list tool.
type ElementListInput struct {
	// ArchiMate element type in kebab-case (e.g., "business-actor", "application-component")
	Type string `json:"type,omitempty"`
	// Filter by name (case-insensitive substring match)
	Name string `json:"name,omitempty"`
}

// List returns a list of ArchiMate elements.
func (t *ElementTools) List(ctx context.Context, input *ElementListInput) (*ListResult, error) {
	resp, err := t.client.ListElements(ctx, input.Type, input.Name)
	if err != nil {
		return nil, WrapError("element_list", err)
	}

	items := make([]any, len(resp.Items))
	for i, v := range resp.Items {
		items[i] = v
	}
	return &ListResult{Items: items, TotalCount: resp.TotalCount}, nil
}

// ElementGetInput represents input for archi_element_get tool.
type ElementGetInput struct {
	// Element ID
	ID string `json:"id"`
}

// Get returns an ArchiMate element by ID.
func (t *ElementTools) Get(ctx context.Context, input *ElementGetInput) (*client.Element, error) {
	if input.ID == "" {
		return nil, fmt.Errorf("%w: id is required", ErrInvalidInput)
	}
	result, err := t.client.GetElement(ctx, input.ID)
	if err != nil {
		return nil, WrapError("element_get", err)
	}
	return result, nil
}

// ElementCreateInput represents input for archi_element_create tool.
type ElementCreateInput struct {
	// ArchiMate type in kebab-case (e.g., "business-actor")
	Type string `json:"type"`
	// Element name
	Name string `json:"name"`
	// Optional documentation
	Documentation string `json:"documentation,omitempty"`
	// Optional folder ID to place the element in
	FolderID string `json:"folder_id,omitempty"`
}

// Create creates a new ArchiMate element.
func (t *ElementTools) Create(ctx context.Context, input *ElementCreateInput) (*client.Element, error) {
	if input.Type == "" || input.Name == "" {
		return nil, fmt.Errorf("%w: type and name are required", ErrInvalidInput)
	}
	body := map[string]any{
		"type": input.Type,
		"name": input.Name,
	}
	if input.Documentation != "" {
		body["documentation"] = input.Documentation
	}
	if input.FolderID != "" {
		body["folder_id"] = input.FolderID
	}
	result, err := t.client.CreateElement(ctx, body)
	if err != nil {
		return nil, WrapError("element_create", err)
	}
	return result, nil
}

// ElementUpdateInput represents input for archi_element_update tool.
type ElementUpdateInput struct {
	// Element ID
	ID string `json:"id"`
	// New name (optional)
	Name string `json:"name,omitempty"`
	// New documentation (optional)
	Documentation string `json:"documentation,omitempty"`
}

// Update updates an ArchiMate element.
func (t *ElementTools) Update(ctx context.Context, input *ElementUpdateInput) (*client.Element, error) {
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
	result, err := t.client.UpdateElement(ctx, input.ID, body)
	if err != nil {
		return nil, WrapError("element_update", err)
	}
	return result, nil
}

// ElementDeleteInput represents input for archi_element_delete tool.
type ElementDeleteInput struct {
	// Element ID
	ID string `json:"id"`
}

// Delete deletes an ArchiMate element.
func (t *ElementTools) Delete(ctx context.Context, input *ElementDeleteInput) (*client.StatusResponse, error) {
	if input.ID == "" {
		return nil, fmt.Errorf("%w: id is required", ErrInvalidInput)
	}
	result, err := t.client.DeleteElement(ctx, input.ID)
	if err != nil {
		return nil, WrapError("element_delete", err)
	}
	return result, nil
}

// ElementMoveInput represents input for archi_element_move tool.
type ElementMoveInput struct {
	// Element ID
	ID string `json:"id"`
	// Target folder ID
	FolderID string `json:"folder_id"`
}

// Move moves an element to a different folder.
func (t *ElementTools) Move(ctx context.Context, input *ElementMoveInput) (*client.Element, error) {
	if input.ID == "" || input.FolderID == "" {
		return nil, fmt.Errorf("%w: id and folder_id are required", ErrInvalidInput)
	}
	body := map[string]any{"folder_id": input.FolderID}
	result, err := t.client.MoveElement(ctx, input.ID, body)
	if err != nil {
		return nil, WrapError("element_move", err)
	}
	return result, nil
}
