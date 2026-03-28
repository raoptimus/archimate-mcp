package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// Client communicates with the jArchi HTTP server.
type Client struct {
	baseURL    string
	httpClient *http.Client
}

// NewClient creates a new jArchi HTTP client.
func NewClient(baseURL string, timeout time.Duration) *Client {
	return &Client{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}
}

func (c *Client) doRequest(ctx context.Context, method, path string, body any, result any) error {
	var bodyReader io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(data)
	}

	req, err := http.NewRequestWithContext(ctx, method, c.baseURL+path, bodyReader)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("http request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode >= 400 {
		var errResp ErrorResponse
		if json.Unmarshal(respBody, &errResp) == nil && errResp.Error != "" {
			return fmt.Errorf("%s (status %d)", errResp.Error, resp.StatusCode)
		}
		return fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(respBody))
	}

	if result != nil {
		if err := json.Unmarshal(respBody, result); err != nil {
			return fmt.Errorf("unmarshal response: %w", err)
		}
	}

	return nil
}

// --- Model ---

func (c *Client) GetModelInfo(ctx context.Context) (*ModelInfo, error) {
	var result ModelInfo
	err := c.doRequest(ctx, http.MethodGet, "/api/model", nil, &result)
	return &result, err
}

func (c *Client) SaveModel(ctx context.Context) (*StatusResponse, error) {
	var result StatusResponse
	err := c.doRequest(ctx, http.MethodPost, "/api/model/save", nil, &result)
	return &result, err
}

// --- Folders ---

func (c *Client) ListFolders(ctx context.Context) (*ListResponse[Folder], error) {
	var result ListResponse[Folder]
	err := c.doRequest(ctx, http.MethodGet, "/api/folders", nil, &result)
	return &result, err
}

func (c *Client) GetFolder(ctx context.Context, id string) (*Folder, error) {
	var result Folder
	err := c.doRequest(ctx, http.MethodGet, "/api/folders/"+id, nil, &result)
	return &result, err
}

func (c *Client) CreateFolder(ctx context.Context, body map[string]any) (*Folder, error) {
	var result Folder
	err := c.doRequest(ctx, http.MethodPost, "/api/folders", body, &result)
	return &result, err
}

// --- Model Operations ---

func (c *Client) RepairModel(ctx context.Context) (*StatusResponse, error) {
	var result StatusResponse
	err := c.doRequest(ctx, http.MethodPost, "/api/model/repair", nil, &result)
	return &result, err
}

// --- Elements ---

func (c *Client) ListElements(ctx context.Context, elemType, name string) (*ListResponse[Element], error) {
	params := url.Values{}
	if elemType != "" {
		params.Set("type", elemType)
	}
	if name != "" {
		params.Set("name", name)
	}
	path := "/api/elements"
	if len(params) > 0 {
		path += "?" + params.Encode()
	}

	var result ListResponse[Element]
	err := c.doRequest(ctx, http.MethodGet, path, nil, &result)
	return &result, err
}

func (c *Client) GetElement(ctx context.Context, id string) (*Element, error) {
	var result Element
	err := c.doRequest(ctx, http.MethodGet, "/api/elements/"+id, nil, &result)
	return &result, err
}

func (c *Client) CreateElement(ctx context.Context, body map[string]any) (*Element, error) {
	var result Element
	err := c.doRequest(ctx, http.MethodPost, "/api/elements", body, &result)
	return &result, err
}

func (c *Client) UpdateElement(ctx context.Context, id string, body map[string]any) (*Element, error) {
	var result Element
	err := c.doRequest(ctx, http.MethodPut, "/api/elements/"+id, body, &result)
	return &result, err
}

func (c *Client) DeleteElement(ctx context.Context, id string) (*StatusResponse, error) {
	var result StatusResponse
	err := c.doRequest(ctx, http.MethodDelete, "/api/elements/"+id, nil, &result)
	return &result, err
}

func (c *Client) MoveElement(ctx context.Context, id string, body map[string]any) (*Element, error) {
	var result Element
	err := c.doRequest(ctx, http.MethodPost, "/api/elements/"+id+"/move", body, &result)
	return &result, err
}

// --- Relationships ---

func (c *Client) ListRelationships(ctx context.Context, relType, sourceID, targetID string) (*ListResponse[Relationship], error) {
	params := url.Values{}
	if relType != "" {
		params.Set("type", relType)
	}
	if sourceID != "" {
		params.Set("source_id", sourceID)
	}
	if targetID != "" {
		params.Set("target_id", targetID)
	}
	path := "/api/relationships"
	if len(params) > 0 {
		path += "?" + params.Encode()
	}

	var result ListResponse[Relationship]
	err := c.doRequest(ctx, http.MethodGet, path, nil, &result)
	return &result, err
}

func (c *Client) GetRelationship(ctx context.Context, id string) (*Relationship, error) {
	var result Relationship
	err := c.doRequest(ctx, http.MethodGet, "/api/relationships/"+id, nil, &result)
	return &result, err
}

func (c *Client) CreateRelationship(ctx context.Context, body map[string]any) (*Relationship, error) {
	var result Relationship
	err := c.doRequest(ctx, http.MethodPost, "/api/relationships", body, &result)
	return &result, err
}

func (c *Client) UpdateRelationship(ctx context.Context, id string, body map[string]any) (*Relationship, error) {
	var result Relationship
	err := c.doRequest(ctx, http.MethodPut, "/api/relationships/"+id, body, &result)
	return &result, err
}

func (c *Client) DeleteRelationship(ctx context.Context, id string) (*StatusResponse, error) {
	var result StatusResponse
	err := c.doRequest(ctx, http.MethodDelete, "/api/relationships/"+id, nil, &result)
	return &result, err
}

func (c *Client) MoveRelationship(ctx context.Context, id string, body map[string]any) (*Relationship, error) {
	var result Relationship
	err := c.doRequest(ctx, http.MethodPost, "/api/relationships/"+id+"/move", body, &result)
	return &result, err
}

// --- Views ---

func (c *Client) ListViews(ctx context.Context) (*ListResponse[View], error) {
	var result ListResponse[View]
	err := c.doRequest(ctx, http.MethodGet, "/api/views", nil, &result)
	return &result, err
}

func (c *Client) GetView(ctx context.Context, id string) (*View, error) {
	var result View
	err := c.doRequest(ctx, http.MethodGet, "/api/views/"+id, nil, &result)
	return &result, err
}

func (c *Client) CreateView(ctx context.Context, body map[string]any) (*View, error) {
	var result View
	err := c.doRequest(ctx, http.MethodPost, "/api/views", body, &result)
	return &result, err
}

func (c *Client) UpdateView(ctx context.Context, id string, body map[string]any) (*View, error) {
	var result View
	err := c.doRequest(ctx, http.MethodPut, "/api/views/"+id, body, &result)
	return &result, err
}

func (c *Client) DeleteView(ctx context.Context, id string) (*StatusResponse, error) {
	var result StatusResponse
	err := c.doRequest(ctx, http.MethodDelete, "/api/views/"+id, nil, &result)
	return &result, err
}

func (c *Client) MoveView(ctx context.Context, id string, body map[string]any) (*View, error) {
	var result View
	err := c.doRequest(ctx, http.MethodPost, "/api/views/"+id+"/move", body, &result)
	return &result, err
}

func (c *Client) AutoConnectView(ctx context.Context, viewID string) (*StatusResponse, error) {
	var result StatusResponse
	err := c.doRequest(ctx, http.MethodPost, "/api/views/"+viewID+"/auto-connect", nil, &result)
	return &result, err
}

func (c *Client) ExportViewAsImage(ctx context.Context, viewID string, scale float64, margin int) (*ViewImageExport, error) {
	params := url.Values{}
	if scale > 0 {
		params.Set("scale", strconv.FormatFloat(scale, 'f', -1, 64))
	}
	if margin > 0 {
		params.Set("margin", strconv.Itoa(margin))
	}
	path := "/api/views/" + viewID + "/export-image"
	if len(params) > 0 {
		path += "?" + params.Encode()
	}
	var result ViewImageExport
	err := c.doRequest(ctx, http.MethodGet, path, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// --- View Objects ---

func (c *Client) ListViewObjects(ctx context.Context, viewID string, recursive bool) (*ListResponse[ViewObject], error) {
	path := "/api/views/" + viewID + "/objects"
	if recursive {
		path += "?recursive=true"
	}
	var result ListResponse[ViewObject]
	err := c.doRequest(ctx, http.MethodGet, path, nil, &result)
	return &result, err
}

func (c *Client) AddViewObject(ctx context.Context, viewID string, body map[string]any) (*ViewObject, error) {
	var result ViewObject
	err := c.doRequest(ctx, http.MethodPost, "/api/views/"+viewID+"/objects", body, &result)
	return &result, err
}

func (c *Client) UpdateViewObject(ctx context.Context, viewID, objID string, body map[string]any) (*ViewObject, error) {
	var result ViewObject
	err := c.doRequest(ctx, http.MethodPut, "/api/views/"+viewID+"/objects/"+objID, body, &result)
	return &result, err
}

func (c *Client) DeleteViewObject(ctx context.Context, viewID, objID string) (*StatusResponse, error) {
	var result StatusResponse
	err := c.doRequest(ctx, http.MethodDelete, "/api/views/"+viewID+"/objects/"+objID, nil, &result)
	return &result, err
}

func (c *Client) CopyViewObjectStyle(ctx context.Context, viewID, objID string, body map[string]any) (*ViewObject, error) {
	var result ViewObject
	err := c.doRequest(ctx, http.MethodPost, "/api/views/"+viewID+"/objects/"+objID+"/copy-style", body, &result)
	return &result, err
}

func (c *Client) CloneViewObject(ctx context.Context, viewID string, body map[string]any) (*ViewObject, error) {
	var result ViewObject
	err := c.doRequest(ctx, http.MethodPost, "/api/views/"+viewID+"/objects/clone", body, &result)
	return &result, err
}

// --- View Connections ---

func (c *Client) ListViewConnections(ctx context.Context, viewID string) (*ListResponse[ViewConnection], error) {
	var result ListResponse[ViewConnection]
	err := c.doRequest(ctx, http.MethodGet, "/api/views/"+viewID+"/connections", nil, &result)
	return &result, err
}

func (c *Client) AddViewConnection(ctx context.Context, viewID string, body map[string]any) (*ViewConnection, error) {
	var result ViewConnection
	err := c.doRequest(ctx, http.MethodPost, "/api/views/"+viewID+"/connections", body, &result)
	return &result, err
}

func (c *Client) UpdateViewConnection(ctx context.Context, viewID, connID string, body map[string]any) (*ViewConnection, error) {
	var result ViewConnection
	err := c.doRequest(ctx, http.MethodPut, "/api/views/"+viewID+"/connections/"+connID, body, &result)
	return &result, err
}

func (c *Client) DeleteViewConnection(ctx context.Context, viewID, connID string) (*StatusResponse, error) {
	var result StatusResponse
	err := c.doRequest(ctx, http.MethodDelete, "/api/views/"+viewID+"/connections/"+connID, nil, &result)
	return &result, err
}

func (c *Client) CopyViewConnectionStyle(ctx context.Context, viewID, connID string, body map[string]any) (*ViewConnection, error) {
	var result ViewConnection
	err := c.doRequest(ctx, http.MethodPost, "/api/views/"+viewID+"/connections/"+connID+"/copy-style", body, &result)
	return &result, err
}

func (c *Client) ExportViewAsSVG(ctx context.Context, viewID string) (*ViewImageExport, error) {
	var result ViewImageExport
	err := c.doRequest(ctx, http.MethodGet, "/api/views/"+viewID+"/export-svg", nil, &result)
	return &result, err
}

// --- Properties ---

func (c *Client) ListProperties(ctx context.Context, id string) (*ListResponse[Property], error) {
	var result ListResponse[Property]
	err := c.doRequest(ctx, http.MethodGet, "/api/properties/"+id, nil, &result)
	return &result, err
}

func (c *Client) SetProperty(ctx context.Context, id string, body map[string]any) (*Property, error) {
	var result Property
	err := c.doRequest(ctx, http.MethodPost, "/api/properties/"+id, body, &result)
	return &result, err
}
