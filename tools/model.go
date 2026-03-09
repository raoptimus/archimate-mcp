package tools

import (
	"context"

	"github.com/raoptimus/archimate-mcp/client"
)

// ModelTools provides MCP tool handlers for model operations.
type ModelTools struct {
	client *client.Client
}

// NewModelTools creates a new ModelTools instance.
func NewModelTools(c *client.Client) *ModelTools {
	return &ModelTools{client: c}
}

// ModelInfoInput represents input for archi_model_info tool.
type ModelInfoInput struct{}

// Info returns model metadata.
func (t *ModelTools) Info(ctx context.Context, input *ModelInfoInput) (*client.ModelInfo, error) {
	result, err := t.client.GetModelInfo(ctx)
	if err != nil {
		return nil, WrapError("model_info", err)
	}
	return result, nil
}

// ModelSaveInput represents input for archi_model_save tool.
type ModelSaveInput struct{}

// Save saves the current model.
func (t *ModelTools) Save(ctx context.Context, input *ModelSaveInput) (*client.StatusResponse, error) {
	result, err := t.client.SaveModel(ctx)
	if err != nil {
		return nil, WrapError("model_save", err)
	}
	return result, nil
}
