package tools

import (
	"context"
	"fmt"

	"github.com/raoptimus/archimate-mcp/client"
)

// PropertyTools provides MCP tool handlers for property operations.
type PropertyTools struct {
	client *client.Client
}

// NewPropertyTools creates a new PropertyTools instance.
func NewPropertyTools(c *client.Client) *PropertyTools {
	return &PropertyTools{client: c}
}

// PropertyListInput represents input for archi_property_list tool.
type PropertyListInput struct {
	// Element or relationship ID
	ID string `json:"id"`
}

// List returns properties of an element or relationship.
func (t *PropertyTools) List(ctx context.Context, input *PropertyListInput) (*ListResult, error) {
	if input.ID == "" {
		return nil, fmt.Errorf("%w: id is required", ErrInvalidInput)
	}
	resp, err := t.client.ListProperties(ctx, input.ID)
	if err != nil {
		return nil, WrapError("property_list", err)
	}

	items := make([]any, len(resp.Items))
	for i, v := range resp.Items {
		items[i] = v
	}
	return &ListResult{Items: items, TotalCount: resp.TotalCount}, nil
}

// PropertySetInput represents input for archi_property_set tool.
type PropertySetInput struct {
	// Element or relationship ID
	ID string `json:"id"`
	// Property key
	Key string `json:"key"`
	// Property value
	Value string `json:"value"`
}

// Set sets a property on an element or relationship.
func (t *PropertyTools) Set(ctx context.Context, input *PropertySetInput) (*client.Property, error) {
	if input.ID == "" || input.Key == "" {
		return nil, fmt.Errorf("%w: id and key are required", ErrInvalidInput)
	}
	body := map[string]any{
		"key":   input.Key,
		"value": input.Value,
	}
	result, err := t.client.SetProperty(ctx, input.ID, body)
	if err != nil {
		return nil, WrapError("property_set", err)
	}
	return result, nil
}
