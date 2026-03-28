package tools

import (
	"context"
	"encoding/json"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/raoptimus/archimate-mcp/client"
)

// Registry holds all tool handlers.
type Registry struct {
	Folder         *FolderTools
	Element        *ElementTools
	Relationship   *RelationshipTools
	View           *ViewTools
	ViewObject     *ViewObjectTools
	ViewConnection *ViewConnectionTools
	Model          *ModelTools
	Property       *PropertyTools
}

// NewRegistry creates a new Registry with all tools initialized.
func NewRegistry(c *client.Client) *Registry {
	return &Registry{
		Folder:         NewFolderTools(c),
		Element:        NewElementTools(c),
		Relationship:   NewRelationshipTools(c),
		View:           NewViewTools(c),
		ViewObject:     NewViewObjectTools(c),
		ViewConnection: NewViewConnectionTools(c),
		Model:          NewModelTools(c),
		Property:       NewPropertyTools(c),
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
	// Folder tools
	mcp.AddTool(server, &mcp.Tool{
		Name:        "archi_folder_list",
		Description: "List all folders in the ArchiMate model recursively. Returns folder id, name, and parent_id.",
	}, wrapHandler(r.Folder.List))

	mcp.AddTool(server, &mcp.Tool{
		Name:        "archi_folder_get",
		Description: "Get an ArchiMate folder by ID. Returns folder details including element_count and subfolder_count.",
	}, wrapHandler(r.Folder.Get))

	mcp.AddTool(server, &mcp.Tool{
		Name:        "archi_folder_create",
		Description: "Create a new folder inside a parent folder",
	}, wrapHandler(r.Folder.Create))

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

	mcp.AddTool(server, &mcp.Tool{
		Name:        "archi_element_move",
		Description: "Move an ArchiMate element to a different folder. Returns the element with updated folder_id, folder_name, folder_path.",
	}, wrapHandler(r.Element.Move))

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

	mcp.AddTool(server, &mcp.Tool{
		Name:        "archi_relationship_move",
		Description: "Move an ArchiMate relationship to a different folder",
	}, wrapHandler(r.Relationship.Move))

	// View tools
	mcp.AddTool(server, &mcp.Tool{
		Name:        "archi_view_list",
		Description: "List all ArchiMate views (diagrams) in the model",
	}, wrapHandler(r.View.List))

	mcp.AddTool(server, &mcp.Tool{
		Name:        "archi_view_get",
		Description: "Get an ArchiMate view by ID. Returns folder_id, folder_path, object_count, connection_count.",
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
		Name:        "archi_view_move",
		Description: "Move an ArchiMate view to a different folder",
	}, wrapHandler(r.View.Move))

	mcp.AddTool(server, &mcp.Tool{
		Name:        "archi_view_auto_connect",
		Description: "Automatically create visual connections on a view for all relationships between elements already placed on the view. Returns the number of connections created.",
	}, wrapHandler(r.View.AutoConnect))

	mcp.AddTool(server, &mcp.Tool{
		Name:        "archi_view_export_image",
		Description: "Export an ArchiMate view as a PNG image (base64-encoded). Returns id, name, format, width, height, mime_type, and data fields.",
	}, wrapHandler(r.View.ExportImage))

	mcp.AddTool(server, &mcp.Tool{
		Name:        "archi_view_export_svg",
		Description: "Export an ArchiMate view as SVG. Note: requires SVG export plugin in Archi.",
	}, wrapHandler(r.View.ExportSVG))

	// View Object tools
	mcp.AddTool(server, &mcp.Tool{
		Name:        "archi_view_object_list",
		Description: "List objects on a view. Set recursive=true to include nested objects with parent_object_id, depth, and children_count. Each object includes style information.",
	}, wrapHandler(r.ViewObject.List))

	mcp.AddTool(server, &mcp.Tool{
		Name:        "archi_view_object_add",
		Description: "Add an ArchiMate element to a view at specified position. Supports parent_object_id to nest inside a container (e.g., System Software inside Application Component).",
	}, wrapHandler(r.ViewObject.Add))

	mcp.AddTool(server, &mcp.Tool{
		Name:        "archi_view_object_update",
		Description: "Update position, size, parent, and/or style of an element on a view. Supports parent_object_id for reparenting (empty string = move to view root). Style fields: fill_color, line_color, font_color, line_width, opacity, text_alignment.",
	}, wrapHandler(r.ViewObject.Update))

	mcp.AddTool(server, &mcp.Tool{
		Name:        "archi_view_object_remove",
		Description: "Remove an element from a view (does not delete the element itself)",
	}, wrapHandler(r.ViewObject.Remove))

	mcp.AddTool(server, &mcp.Tool{
		Name:        "archi_view_object_copy_style",
		Description: "Copy visual style (fill_color, line_color, font_color, line_width, opacity, text_alignment) from a source object to a target object on the same view",
	}, wrapHandler(r.ViewObject.CopyStyle))

	mcp.AddTool(server, &mcp.Tool{
		Name:        "archi_view_object_clone_from_template",
		Description: "Create a new diagram object with the same style and size as a source object but for a different element. Useful for creating visually consistent diagrams.",
	}, wrapHandler(r.ViewObject.Clone))

	// View Connection tools
	mcp.AddTool(server, &mcp.Tool{
		Name:        "archi_view_connection_list",
		Description: "List all visual connections on a view. Returns relationship_id, relationship_type, source_object_id, target_object_id for each connection.",
	}, wrapHandler(r.ViewConnection.List))

	mcp.AddTool(server, &mcp.Tool{
		Name:        "archi_view_connection_add",
		Description: "Add a visual connection on a view for an existing relationship. Requires relationship_id, source_object_id (diagram object), and target_object_id (diagram object).",
	}, wrapHandler(r.ViewConnection.Add))

	mcp.AddTool(server, &mcp.Tool{
		Name:        "archi_view_connection_update",
		Description: "Update a visual connection on a view",
	}, wrapHandler(r.ViewConnection.Update))

	mcp.AddTool(server, &mcp.Tool{
		Name:        "archi_view_connection_remove",
		Description: "Remove a visual connection from a view (does not delete the underlying relationship)",
	}, wrapHandler(r.ViewConnection.Remove))

	mcp.AddTool(server, &mcp.Tool{
		Name:        "archi_view_connection_copy_style",
		Description: "Copy visual style (line_color, line_width) from a source connection to a target connection",
	}, wrapHandler(r.ViewConnection.CopyStyle))

	// Model tools
	mcp.AddTool(server, &mcp.Tool{
		Name:        "archi_model_info",
		Description: "Get ArchiMate model metadata (id, name, purpose)",
	}, wrapHandler(r.Model.Info))

	mcp.AddTool(server, &mcp.Tool{
		Name:        "archi_model_save",
		Description: "Save the current ArchiMate model",
	}, wrapHandler(r.Model.Save))

	mcp.AddTool(server, &mcp.Tool{
		Name:        "archi_model_repair",
		Description: "Remove orphaned diagram references — objects with null concepts and connections with null relationships. Returns count of removed orphans.",
	}, wrapHandler(r.Model.Repair))

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
