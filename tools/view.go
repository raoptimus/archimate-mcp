package tools

import (
	"context"
	"fmt"

	"github.com/raoptimus/archimate-mcp/client"
)

// ViewTools provides MCP tool handlers for ArchiMate view operations.
type ViewTools struct {
	client *client.Client
}

// NewViewTools creates a new ViewTools instance.
func NewViewTools(c *client.Client) *ViewTools {
	return &ViewTools{client: c}
}

// ViewListInput represents input for archi_view_list tool.
type ViewListInput struct{}

// List returns a list of all ArchiMate views.
func (t *ViewTools) List(ctx context.Context, input *ViewListInput) (*ListResult, error) {
	resp, err := t.client.ListViews(ctx)
	if err != nil {
		return nil, WrapError("view_list", err)
	}

	items := make([]any, len(resp.Items))
	for i, v := range resp.Items {
		items[i] = v
	}
	return &ListResult{Items: items, TotalCount: resp.TotalCount}, nil
}

// ViewGetInput represents input for archi_view_get tool.
type ViewGetInput struct {
	ID string `json:"id"`
}

// Get returns an ArchiMate view by ID.
func (t *ViewTools) Get(ctx context.Context, input *ViewGetInput) (*client.View, error) {
	if input.ID == "" {
		return nil, fmt.Errorf("%w: id is required", ErrInvalidInput)
	}
	result, err := t.client.GetView(ctx, input.ID)
	if err != nil {
		return nil, WrapError("view_get", err)
	}
	return result, nil
}

// ViewCreateInput represents input for archi_view_create tool.
type ViewCreateInput struct {
	Name          string `json:"name"`
	Documentation string `json:"documentation,omitempty"`
	// Optional folder ID to place the view in
	FolderID string `json:"folder_id,omitempty"`
}

// Create creates a new ArchiMate view.
func (t *ViewTools) Create(ctx context.Context, input *ViewCreateInput) (*client.View, error) {
	if input.Name == "" {
		return nil, fmt.Errorf("%w: name is required", ErrInvalidInput)
	}
	body := map[string]any{"name": input.Name}
	if input.Documentation != "" {
		body["documentation"] = input.Documentation
	}
	if input.FolderID != "" {
		body["folder_id"] = input.FolderID
	}
	result, err := t.client.CreateView(ctx, body)
	if err != nil {
		return nil, WrapError("view_create", err)
	}
	return result, nil
}

// ViewUpdateInput represents input for archi_view_update tool.
type ViewUpdateInput struct {
	ID            string `json:"id"`
	Name          string `json:"name,omitempty"`
	Documentation string `json:"documentation,omitempty"`
}

// Update updates an ArchiMate view.
func (t *ViewTools) Update(ctx context.Context, input *ViewUpdateInput) (*client.View, error) {
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
	result, err := t.client.UpdateView(ctx, input.ID, body)
	if err != nil {
		return nil, WrapError("view_update", err)
	}
	return result, nil
}

// ViewDeleteInput represents input for archi_view_delete tool.
type ViewDeleteInput struct {
	ID string `json:"id"`
}

// Delete deletes an ArchiMate view.
func (t *ViewTools) Delete(ctx context.Context, input *ViewDeleteInput) (*client.StatusResponse, error) {
	if input.ID == "" {
		return nil, fmt.Errorf("%w: id is required", ErrInvalidInput)
	}
	result, err := t.client.DeleteView(ctx, input.ID)
	if err != nil {
		return nil, WrapError("view_delete", err)
	}
	return result, nil
}

// ViewMoveInput represents input for archi_view_move tool.
type ViewMoveInput struct {
	// View ID
	ID string `json:"id"`
	// Target folder ID
	FolderID string `json:"folder_id"`
}

// Move moves a view to a different folder.
func (t *ViewTools) Move(ctx context.Context, input *ViewMoveInput) (*client.View, error) {
	if input.ID == "" || input.FolderID == "" {
		return nil, fmt.Errorf("%w: id and folder_id are required", ErrInvalidInput)
	}
	body := map[string]any{"folder_id": input.FolderID}
	result, err := t.client.MoveView(ctx, input.ID, body)
	if err != nil {
		return nil, WrapError("view_move", err)
	}
	return result, nil
}

// ViewAutoConnectInput represents input for archi_view_auto_connect tool.
type ViewAutoConnectInput struct {
	// View ID
	ID string `json:"id"`
}

// AutoConnect automatically creates visual connections for all relationships between elements on the view.
func (t *ViewTools) AutoConnect(ctx context.Context, input *ViewAutoConnectInput) (*client.StatusResponse, error) {
	if input.ID == "" {
		return nil, fmt.Errorf("%w: id is required", ErrInvalidInput)
	}
	result, err := t.client.AutoConnectView(ctx, input.ID)
	if err != nil {
		return nil, WrapError("view_auto_connect", err)
	}
	return result, nil
}

// ViewExportImageInput represents input for archi_view_export_image tool.
type ViewExportImageInput struct {
	ID     string  `json:"id"`
	Scale  float64 `json:"scale,omitempty"`  // 1.0-4.0, default 1
	Margin int     `json:"margin,omitempty"` // pixels, default 10
}

// ExportImage exports an ArchiMate view as a PNG image (base64-encoded).
func (t *ViewTools) ExportImage(ctx context.Context, input *ViewExportImageInput) (*client.ViewImageExport, error) {
	if input.ID == "" {
		return nil, fmt.Errorf("%w: id is required", ErrInvalidInput)
	}
	scale := input.Scale
	if scale <= 0 {
		scale = 1.0
	}
	margin := input.Margin
	if margin <= 0 {
		margin = 10
	}
	result, err := t.client.ExportViewAsImage(ctx, input.ID, scale, margin)
	if err != nil {
		return nil, WrapError("view_export_image", err)
	}
	return result, nil
}

// ViewExportSVGInput represents input for archi_view_export_svg tool.
type ViewExportSVGInput struct {
	ID string `json:"id"`
}

// ExportSVG exports an ArchiMate view as SVG.
func (t *ViewTools) ExportSVG(ctx context.Context, input *ViewExportSVGInput) (*client.ViewImageExport, error) {
	if input.ID == "" {
		return nil, fmt.Errorf("%w: id is required", ErrInvalidInput)
	}
	result, err := t.client.ExportViewAsSVG(ctx, input.ID)
	if err != nil {
		return nil, WrapError("view_export_svg", err)
	}
	return result, nil
}
