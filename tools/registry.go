package tools

import (
	"context"
	"encoding/json"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/raoptimus/archimate-mcp/client"
)

// Registry holds all tool handlers.
type Registry struct {
	Element      *ElementTools
	Relationship *RelationshipTools
	View         *ViewTools
	ViewObject   *ViewObjectTools
	Model        *ModelTools
	Property     *PropertyTools
}

// NewRegistry creates a new Registry with all tools initialized.
func NewRegistry(c *client.Client) *Registry {
	return &Registry{
		Element:      NewElementTools(c),
		Relationship: NewRelationshipTools(c),
		View:         NewViewTools(c),
		ViewObject:   NewViewObjectTools(c),
		Model:        NewModelTools(c),
		Property:     NewPropertyTools(c),
	}
}

// wrapHandler wraps a typed handler function to work with MCP's generic interface.
func wrapHandler[In, Out any](handler func(context.Context, In) (Out, error)) func(context.Context, *mcp.CallToolRequest, In) (*mcp.CallToolResult, Out, error) {
	return func(ctx context.Context, req *mcp.CallToolRequest, args In) (*mcp.CallToolResult, Out, error) {
		result, err := handler(ctx, args)
		if err != nil {
			var zero Out
			return &mcp.CallToolResult{
				Content: []mcp.Content{
					&mcp.TextContent{Text: FormatToolError(err)},
				},
				IsError: true,
			}, zero, nil
		}

		jsonBytes, jsonErr := json.MarshalIndent(result, "", "  ")
		if jsonErr != nil {
			var zero Out
			return &mcp.CallToolResult{
				Content: []mcp.Content{
					&mcp.TextContent{Text: "Failed to serialize result: " + jsonErr.Error()},
				},
				IsError: true,
			}, zero, jsonErr
		}

		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: string(jsonBytes)},
			},
		}, result, nil
	}
}

// RegisterAll registers all tools with the MCP server.
func (r *Registry) RegisterAll(server *mcp.Server) {
	// Element tools
	mcp.AddTool(server, &mcp.Tool{
		Name:        "archi_element_list",
		Description: "List ArchiMate elements with optional filters by type and name. Types use kebab-case: business-actor, business-process, application-component, etc.",
	}, wrapHandler(r.Element.List))

	mcp.AddTool(server, &mcp.Tool{
		Name:        "archi_element_get",
		Description: "Get an ArchiMate element by ID",
	}, wrapHandler(r.Element.Get))

	mcp.AddTool(server, &mcp.Tool{
		Name:        "archi_element_create",
		Description: "Create a new ArchiMate element. Type uses kebab-case: business-actor, business-role, business-process, business-object, business-service, application-component, application-service, technology-node, etc.",
	}, wrapHandler(r.Element.Create))

	mcp.AddTool(server, &mcp.Tool{
		Name:        "archi_element_update",
		Description: "Update an ArchiMate element's name or documentation",
	}, wrapHandler(r.Element.Update))

	mcp.AddTool(server, &mcp.Tool{
		Name:        "archi_element_delete",
		Description: "Delete an ArchiMate element by ID",
	}, wrapHandler(r.Element.Delete))

	// Relationship tools
	mcp.AddTool(server, &mcp.Tool{
		Name:        "archi_relationship_list",
		Description: "List ArchiMate relationships with optional filters by type, source_id, target_id. Types: serving-relationship, composition-relationship, aggregation-relationship, assignment-relationship, realization-relationship, flow-relationship, triggering-relationship, association-relationship, access-relationship",
	}, wrapHandler(r.Relationship.List))

	mcp.AddTool(server, &mcp.Tool{
		Name:        "archi_relationship_get",
		Description: "Get an ArchiMate relationship by ID",
	}, wrapHandler(r.Relationship.Get))

	mcp.AddTool(server, &mcp.Tool{
		Name:        "archi_relationship_create",
		Description: "Create a new ArchiMate relationship between two elements",
	}, wrapHandler(r.Relationship.Create))

	mcp.AddTool(server, &mcp.Tool{
		Name:        "archi_relationship_update",
		Description: "Update an ArchiMate relationship's name or documentation",
	}, wrapHandler(r.Relationship.Update))

	mcp.AddTool(server, &mcp.Tool{
		Name:        "archi_relationship_delete",
		Description: "Delete an ArchiMate relationship by ID",
	}, wrapHandler(r.Relationship.Delete))

	// View tools
	mcp.AddTool(server, &mcp.Tool{
		Name:        "archi_view_list",
		Description: "List all ArchiMate views (diagrams) in the model",
	}, wrapHandler(r.View.List))

	mcp.AddTool(server, &mcp.Tool{
		Name:        "archi_view_get",
		Description: "Get an ArchiMate view by ID",
	}, wrapHandler(r.View.Get))

	mcp.AddTool(server, &mcp.Tool{
		Name:        "archi_view_create",
		Description: "Create a new ArchiMate view (diagram)",
	}, wrapHandler(r.View.Create))

	mcp.AddTool(server, &mcp.Tool{
		Name:        "archi_view_update",
		Description: "Update an ArchiMate view's name or documentation",
	}, wrapHandler(r.View.Update))

	mcp.AddTool(server, &mcp.Tool{
		Name:        "archi_view_delete",
		Description: "Delete an ArchiMate view by ID",
	}, wrapHandler(r.View.Delete))

	mcp.AddTool(server, &mcp.Tool{
		Name:        "archi_view_export_image",
		Description: "Export an ArchiMate view as a PNG image (base64-encoded). Returns JSON with id, name, format, and data fields.",
	}, wrapHandler(r.View.ExportImage))

	// View Object tools
	mcp.AddTool(server, &mcp.Tool{
		Name:        "archi_view_object_list",
		Description: "List all objects (elements) placed on an ArchiMate view",
	}, wrapHandler(r.ViewObject.List))

	mcp.AddTool(server, &mcp.Tool{
		Name:        "archi_view_object_add",
		Description: "Add an ArchiMate element to a view at specified position (x, y, width, height)",
	}, wrapHandler(r.ViewObject.Add))

	mcp.AddTool(server, &mcp.Tool{
		Name:        "archi_view_object_update",
		Description: "Update position and size of an element on a view",
	}, wrapHandler(r.ViewObject.Update))

	mcp.AddTool(server, &mcp.Tool{
		Name:        "archi_view_object_remove",
		Description: "Remove an element from a view (does not delete the element itself)",
	}, wrapHandler(r.ViewObject.Remove))

	// Model tools
	mcp.AddTool(server, &mcp.Tool{
		Name:        "archi_model_info",
		Description: "Get ArchiMate model metadata (id, name, purpose)",
	}, wrapHandler(r.Model.Info))

	mcp.AddTool(server, &mcp.Tool{
		Name:        "archi_model_save",
		Description: "Save the current ArchiMate model",
	}, wrapHandler(r.Model.Save))

	// Property tools
	mcp.AddTool(server, &mcp.Tool{
		Name:        "archi_property_list",
		Description: "List properties of an ArchiMate element or relationship",
	}, wrapHandler(r.Property.List))

	mcp.AddTool(server, &mcp.Tool{
		Name:        "archi_property_set",
		Description: "Set a property on an ArchiMate element or relationship",
	}, wrapHandler(r.Property.Set))
}
