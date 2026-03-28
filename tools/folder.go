package tools

import (
	"context"
	"fmt"

	"github.com/raoptimus/archimate-mcp/client"
)

// FolderTools provides MCP tool handlers for ArchiMate folder operations.
type FolderTools struct {
	client *client.Client
}

// NewFolderTools creates a new FolderTools instance.
func NewFolderTools(c *client.Client) *FolderTools {
	return &FolderTools{client: c}
}

// FolderListInput represents input for archi_folder_list tool.
type FolderListInput struct{}

// List returns all folders in the model.
func (t *FolderTools) List(ctx context.Context, input *FolderListInput) (*ListResult, error) {
	resp, err := t.client.ListFolders(ctx)
	if err != nil {
		return nil, WrapError("folder_list", err)
	}

	items := make([]any, len(resp.Items))
	for i, v := range resp.Items {
		items[i] = v
	}
	return &ListResult{Items: items, TotalCount: resp.TotalCount}, nil
}

// FolderGetInput represents input for archi_folder_get tool.
type FolderGetInput struct {
	// Folder ID
	ID string `json:"id"`
}

// Get returns a folder by ID with element and subfolder counts.
func (t *FolderTools) Get(ctx context.Context, input *FolderGetInput) (*client.Folder, error) {
	if input.ID == "" {
		return nil, fmt.Errorf("%w: id is required", ErrInvalidInput)
	}
	result, err := t.client.GetFolder(ctx, input.ID)
	if err != nil {
		return nil, WrapError("folder_get", err)
	}
	return result, nil
}

// FolderCreateInput represents input for archi_folder_create tool.
type FolderCreateInput struct {
	// Folder name
	Name string `json:"name"`
	// Parent folder ID
	ParentID string `json:"parent_id"`
}

// Create creates a new folder inside a parent folder.
func (t *FolderTools) Create(ctx context.Context, input *FolderCreateInput) (*client.Folder, error) {
	if input.Name == "" || input.ParentID == "" {
		return nil, fmt.Errorf("%w: name and parent_id are required", ErrInvalidInput)
	}
	body := map[string]any{
		"name":      input.Name,
		"parent_id": input.ParentID,
	}
	result, err := t.client.CreateFolder(ctx, body)
	if err != nil {
		return nil, WrapError("folder_create", err)
	}
	return result, nil
}
