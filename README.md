# ArchiMate MCP Tool

An MCP server in Go for managing ArchiMate models in [Archi](https://www.archimatetool.com/) via the [jArchi](https://www.archimatetool.com/plugins/) plugin.

Provides 23 MCP tools for creating and editing ArchiMate elements, relationships, views, and properties.

## Architecture

```
Claude ←(stdio/MCP)→ archimate-mcp (Go) ←(HTTP)→ ArchiMCPServer.ajs (jArchi) ←→ Archi
```

Archi has no HTTP API. The jArchi script runs an HTTP server (`com.sun.net.httpserver.HttpServer`) on port 9898 inside Archi, and the Go MCP server communicates with it over HTTP.

## Requirements

- [Archi](https://www.archimatetool.com/) with the [jArchi](https://www.archimatetool.com/plugins/) plugin installed
- Go 1.25+

## Install

```bash
make install
```

The binary `archimate-mcp` will be installed.

## Setup

### 1. Start jArchi HTTP Server

1. Open Archi, create or open a model
2. Go to **Scripts → Scripts Manager** (or **Window → Scripts Manager**)
3. Copy `jarchi/ArchiMCPServer.ajs` to the jArchi scripts folder
4. Select the script and click **Run**
5. The console should show:
   ```
   ArchiMate MCP HTTP Server started on port 9898
   Model: <model name>
   ```

Verify:

```bash
curl http://localhost:9898/api/model
```

### 2. Install MCP Server

#### Claude Code (CLI)

```bash
claude mcp add archimate -- archimate-mcp --archi-url http://localhost:9898
```

#### Claude Desktop

Add to `claude_desktop_config.json`:

```json
{
  "mcpServers": {
    "archimate": {
      "command": "$GOPATH/bin/archimate-mcp",
      "args": ["--archi-url", "http://localhost:9898"]
    }
  }
}
```

#### OpenAI Codex CLI

Add to `~/.codex/config.json`:

```json
{
  "mcpServers": {
    "archimate": {
      "command": "$GOPATH/bin/archimate-mcp",
      "args": ["--archi-url", "http://localhost:9898"]
    }
  }
}
```

### CLI Flags

| Flag | Env | Default | Description |
|------|-----|---------|-------------|
| `--archi-url`, `-u` | `ARCHI_URL` | `http://localhost:9898` | jArchi HTTP server URL |
| `--debug`, `-d` | `ARCHI_DEBUG` | `false` | Enable debug logging |
| `--timeout` | `ARCHI_TIMEOUT` | `30s` | HTTP request timeout |

## Available MCP Tools

### Elements

| Tool | Description |
|------|-------------|
| `archi_element_list` | List elements (filter by type, name) |
| `archi_element_get` | Get element by ID |
| `archi_element_create` | Create element (type, name, documentation) |
| `archi_element_update` | Update name/documentation |
| `archi_element_delete` | Delete element |

Element types (kebab-case): `business-actor`, `business-role`, `business-process`, `business-object`, `business-service`, `application-component`, `application-service`, `technology-node`, `technology-service`, etc.

### Relationships

| Tool | Description |
|------|-------------|
| `archi_relationship_list` | List relationships (filter by type, source_id, target_id) |
| `archi_relationship_get` | Get relationship by ID |
| `archi_relationship_create` | Create relationship (type, source_id, target_id) |
| `archi_relationship_update` | Update name/documentation |
| `archi_relationship_delete` | Delete relationship |

Relationship types: `serving-relationship`, `composition-relationship`, `aggregation-relationship`, `assignment-relationship`, `realization-relationship`, `flow-relationship`, `triggering-relationship`, `association-relationship`, `access-relationship`.

### Views (Diagrams)

| Tool | Description |
|------|-------------|
| `archi_view_list` | List all views |
| `archi_view_get` | Get view by ID |
| `archi_view_create` | Create view |
| `archi_view_update` | Update view |
| `archi_view_delete` | Delete view |

### View Objects

| Tool | Description |
|------|-------------|
| `archi_view_object_list` | List objects on a view |
| `archi_view_object_add` | Add element to view (x, y, width, height) |
| `archi_view_object_update` | Update position/size |
| `archi_view_object_remove` | Remove element from view |

### Model

| Tool | Description |
|------|-------------|
| `archi_model_info` | Model metadata (id, name, purpose) |
| `archi_model_save` | Save model |

### Properties

| Tool | Description |
|------|-------------|
| `archi_property_list` | List properties of element/relationship |
| `archi_property_set` | Set property |

## Verify jArchi

Before starting the MCP server, make sure jArchi is working:

1. Create a test script in Archi:
   ```javascript
   console.log("jArchi works! Model: " + model.name);
   ```
2. Run it — you should see the message in the console.
3. Test element creation:
   ```javascript
   var e = model.createElement("business-actor", "Test Actor");
   console.log("Created: " + e.id);
   ```

## Project Structure

```
├── main.go                    # Entry point, CLI
├── client/
│   ├── client.go              # HTTP client for jArchi
│   └── types.go               # Data types
├── tools/
│   ├── registry.go            # Tool registry
│   ├── common.go              # Common types
│   ├── errors.go              # Error handling
│   ├── element.go             # Element tools
│   ├── relationship.go        # Relationship tools
│   ├── view.go                # View tools
│   ├── view_object.go         # View object tools
│   ├── model.go               # Model tools
│   └── property.go            # Property tools
└── jarchi/
    └── ArchiMCPServer.ajs     # jArchi HTTP server
```
