package tools

import (
	"context"
	"fmt"

	"github.com/raoptimus/archimate-mcp/client"
)

// ViewObjectTools provides MCP tool handlers for view object operations.
type ViewObjectTools struct {
	client *client.Client
}

// NewViewObjectTools creates a new ViewObjectTools instance.
func NewViewObjectTools(c *client.Client) *ViewObjectTools {
	return &ViewObjectTools{client: c}
}

// ViewObjectListInput represents input for archi_view_object_list tool.
type ViewObjectListInput struct {
	// View ID
	ViewID string `json:"view_id"`
}

// List returns all objects on a view.
func (t *ViewObjectTools) List(ctx context.Context, input *ViewObjectListInput) (*ListResult, error) {
	if input.ViewID == "" {
		return nil, fmt.Errorf("%w: view_id is required", ErrInvalidInput)
	}
	resp, err := t.client.ListViewObjects(ctx, input.ViewID)
	if err != nil {
		return nil, WrapError("view_object_list", err)
	}

	items := make([]any, len(resp.Items))
	for i, v := range resp.Items {
		items[i] = v
	}
	return &ListResult{Items: items, TotalCount: resp.TotalCount}, nil
}

// ViewObjectAddInput represents input for archi_view_object_add tool.
type ViewObjectAddInput struct {
	// View ID
	ViewID string `json:"view_id"`
	// Element ID to add to the view
	ElementID string `json:"element_id"`
	// X position (default: 10)
	X int `json:"x,omitempty"`
	// Y position (default: 10)
	Y int `json:"y,omitempty"`
	// Width (default: 120)
	Width int `json:"width,omitempty"`
	// Height (default: 55)
	Height int `json:"height,omitempty"`
}

// Add adds an element to a view.
func (t *ViewObjectTools) Add(ctx context.Context, input *ViewObjectAddInput) (*client.ViewObject, error) {
	if input.ViewID == "" || input.ElementID == "" {
		return nil, fmt.Errorf("%w: view_id and element_id are required", ErrInvalidInput)
	}
	body := map[string]any{
		"element_id": input.ElementID,
	}
	if input.X != 0 {
		body["x"] = input.X
	}
	if input.Y != 0 {
		body["y"] = input.Y
	}
	if input.Width != 0 {
		body["width"] = input.Width
	}
	if input.Height != 0 {
		body["height"] = input.Height
	}
	result, err := t.client.AddViewObject(ctx, input.ViewID, body)
	if err != nil {
		return nil, WrapError("view_object_add", err)
	}
	return result, nil
}

// ViewObjectUpdateInput represents input for archi_view_object_update tool.
type ViewObjectUpdateInput struct {
	// View ID
	ViewID string `json:"view_id"`
	// View object ID
	ObjectID string `json:"object_id"`
	// New X position
	X *int `json:"x,omitempty"`
	// New Y position
	Y *int `json:"y,omitempty"`
	// New width
	Width *int `json:"width,omitempty"`
	// New height
	Height *int `json:"height,omitempty"`
}

// Update updates position/size of an object on a view.
func (t *ViewObjectTools) Update(ctx context.Context, input *ViewObjectUpdateInput) (*client.ViewObject, error) {
	if input.ViewID == "" || input.ObjectID == "" {
		return nil, fmt.Errorf("%w: view_id and object_id are required", ErrInvalidInput)
	}
	body := map[string]any{}
	if input.X != nil {
		body["x"] = *input.X
	}
	if input.Y != nil {
		body["y"] = *input.Y
	}
	if input.Width != nil {
		body["width"] = *input.Width
	}
	if input.Height != nil {
		body["height"] = *input.Height
	}
	result, err := t.client.UpdateViewObject(ctx, input.ViewID, input.ObjectID, body)
	if err != nil {
		return nil, WrapError("view_object_update", err)
	}
	return result, nil
}

// ViewObjectRemoveInput represents input for archi_view_object_remove tool.
type ViewObjectRemoveInput struct {
	// View ID
	ViewID string `json:"view_id"`
	// View object ID
	ObjectID string `json:"object_id"`
}

// Remove removes an object from a view.
func (t *ViewObjectTools) Remove(ctx context.Context, input *ViewObjectRemoveInput) (*client.StatusResponse, error) {
	if input.ViewID == "" || input.ObjectID == "" {
		return nil, fmt.Errorf("%w: view_id and object_id are required", ErrInvalidInput)
	}
	result, err := t.client.DeleteViewObject(ctx, input.ViewID, input.ObjectID)
	if err != nil {
		return nil, WrapError("view_object_remove", err)
	}
	return result, nil
}
