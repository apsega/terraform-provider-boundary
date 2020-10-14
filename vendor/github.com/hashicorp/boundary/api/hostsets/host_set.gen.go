// Code generated by "make api"; DO NOT EDIT.
package hostsets

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/hashicorp/boundary/api"
	"github.com/hashicorp/boundary/api/scopes"
)

type HostSet struct {
	Id            string                 `json:"id,omitempty"`
	HostCatalogId string                 `json:"host_catalog_id,omitempty"`
	Scope         *scopes.ScopeInfo      `json:"scope,omitempty"`
	Name          string                 `json:"name,omitempty"`
	Description   string                 `json:"description,omitempty"`
	CreatedTime   time.Time              `json:"created_time,omitempty"`
	UpdatedTime   time.Time              `json:"updated_time,omitempty"`
	Version       uint32                 `json:"version,omitempty"`
	Type          string                 `json:"type,omitempty"`
	HostIds       []string               `json:"host_ids,omitempty"`
	Attributes    map[string]interface{} `json:"attributes,omitempty"`

	responseBody *bytes.Buffer
	responseMap  map[string]interface{}
}

func (n HostSet) ResponseBody() *bytes.Buffer {
	return n.responseBody
}

func (n HostSet) ResponseMap() map[string]interface{} {
	return n.responseMap
}

type HostSetReadResult struct {
	Item         *HostSet
	responseBody *bytes.Buffer
	responseMap  map[string]interface{}
}

func (n HostSetReadResult) GetItem() interface{} {
	return n.Item
}

func (n HostSetReadResult) GetResponseBody() *bytes.Buffer {
	return n.responseBody
}

func (n HostSetReadResult) GetResponseMap() map[string]interface{} {
	return n.responseMap
}

type HostSetCreateResult = HostSetReadResult
type HostSetUpdateResult = HostSetReadResult

type HostSetDeleteResult struct {
	responseBody *bytes.Buffer
	responseMap  map[string]interface{}
}

func (n HostSetDeleteResult) GetResponseBody() *bytes.Buffer {
	return n.responseBody
}

func (n HostSetDeleteResult) GetResponseMap() map[string]interface{} {
	return n.responseMap
}

type HostSetListResult struct {
	Items        []*HostSet
	responseBody *bytes.Buffer
	responseMap  map[string]interface{}
}

func (n HostSetListResult) GetItems() interface{} {
	return n.Items
}

func (n HostSetListResult) GetResponseBody() *bytes.Buffer {
	return n.responseBody
}

func (n HostSetListResult) GetResponseMap() map[string]interface{} {
	return n.responseMap
}

// Client is a client for this collection
type Client struct {
	client *api.Client
}

// Creates a new client for this collection. The submitted API client is cloned;
// modifications to it after generating this client will not have effect. If you
// need to make changes to the underlying API client, use ApiClient() to access
// it.
func NewClient(c *api.Client) *Client {
	return &Client{client: c.Clone()}
}

// ApiClient returns the underlying API client
func (c *Client) ApiClient() *api.Client {
	return c.client
}

func (c *Client) Create(ctx context.Context, hostCatalogId string, opt ...Option) (*HostSetCreateResult, error) {
	if hostCatalogId == "" {
		return nil, fmt.Errorf("empty hostCatalogId value passed into Create request")
	}

	opts, apiOpts := getOpts(opt...)

	if c.client == nil {
		return nil, fmt.Errorf("nil client")
	}

	opts.postMap["host_catalog_id"] = hostCatalogId

	req, err := c.client.NewRequest(ctx, "POST", "host-sets", opts.postMap, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating Create request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during Create call: %w", err)
	}

	target := new(HostSetCreateResult)
	target.Item = new(HostSet)
	apiErr, err := resp.Decode(target.Item)
	if err != nil {
		return nil, fmt.Errorf("error decoding Create response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}
	target.responseBody = resp.Body
	target.responseMap = resp.Map
	return target, nil
}

func (c *Client) Read(ctx context.Context, hostSetId string, opt ...Option) (*HostSetReadResult, error) {
	if hostSetId == "" {
		return nil, fmt.Errorf("empty hostSetId value passed into Read request")
	}
	if c.client == nil {
		return nil, fmt.Errorf("nil client")
	}

	opts, apiOpts := getOpts(opt...)

	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("host-sets/%s", hostSetId), nil, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating Read request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during Read call: %w", err)
	}

	target := new(HostSetReadResult)
	target.Item = new(HostSet)
	apiErr, err := resp.Decode(target.Item)
	if err != nil {
		return nil, fmt.Errorf("error decoding Read response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}
	target.responseBody = resp.Body
	target.responseMap = resp.Map
	return target, nil
}

func (c *Client) Update(ctx context.Context, hostSetId string, version uint32, opt ...Option) (*HostSetUpdateResult, error) {
	if hostSetId == "" {
		return nil, fmt.Errorf("empty hostSetId value passed into Update request")
	}
	if c.client == nil {
		return nil, fmt.Errorf("nil client")
	}

	opts, apiOpts := getOpts(opt...)

	if version == 0 {
		if !opts.withAutomaticVersioning {
			return nil, errors.New("zero version number passed into Update request and automatic versioning not specified")
		}
		existingTarget, existingErr := c.Read(ctx, hostSetId, opt...)
		if existingErr != nil {
			if api.AsServerError(existingErr) != nil {
				return nil, fmt.Errorf("error from controller when performing initial check-and-set read: %w", existingErr)
			}
			return nil, fmt.Errorf("error performing initial check-and-set read: %w", existingErr)
		}
		if existingTarget == nil {
			return nil, errors.New("nil resource response found when performing initial check-and-set read")
		}
		if existingTarget.Item == nil {
			return nil, errors.New("nil resource found when performing initial check-and-set read")
		}
		version = existingTarget.Item.Version
	}

	opts.postMap["version"] = version

	req, err := c.client.NewRequest(ctx, "PATCH", fmt.Sprintf("host-sets/%s", hostSetId), opts.postMap, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating Update request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during Update call: %w", err)
	}

	target := new(HostSetUpdateResult)
	target.Item = new(HostSet)
	apiErr, err := resp.Decode(target.Item)
	if err != nil {
		return nil, fmt.Errorf("error decoding Update response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}
	target.responseBody = resp.Body
	target.responseMap = resp.Map
	return target, nil
}

func (c *Client) Delete(ctx context.Context, hostSetId string, opt ...Option) (*HostSetDeleteResult, error) {
	if hostSetId == "" {
		return nil, fmt.Errorf("empty hostSetId value passed into Delete request")
	}
	if c.client == nil {
		return nil, fmt.Errorf("nil client")
	}

	opts, apiOpts := getOpts(opt...)

	req, err := c.client.NewRequest(ctx, "DELETE", fmt.Sprintf("host-sets/%s", hostSetId), nil, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating Delete request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during Delete call: %w", err)
	}

	apiErr, err := resp.Decode(nil)
	if err != nil {
		return nil, fmt.Errorf("error decoding Delete response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}

	target := &HostSetDeleteResult{
		responseBody: resp.Body,
		responseMap:  resp.Map,
	}
	return target, nil
}

func (c *Client) List(ctx context.Context, hostCatalogId string, opt ...Option) (*HostSetListResult, error) {
	if hostCatalogId == "" {
		return nil, fmt.Errorf("empty hostCatalogId value passed into List request")
	}
	if c.client == nil {
		return nil, fmt.Errorf("nil client")
	}

	opts, apiOpts := getOpts(opt...)
	opts.queryMap["host_catalog_id"] = hostCatalogId

	req, err := c.client.NewRequest(ctx, "GET", "host-sets", nil, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating List request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during List call: %w", err)
	}

	target := new(HostSetListResult)
	apiErr, err := resp.Decode(target)
	if err != nil {
		return nil, fmt.Errorf("error decoding List response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}
	target.responseBody = resp.Body
	target.responseMap = resp.Map
	return target, nil
}

func (c *Client) AddHosts(ctx context.Context, hostSetId string, version uint32, hostIds []string, opt ...Option) (*HostSetUpdateResult, error) {
	if hostSetId == "" {
		return nil, fmt.Errorf("empty hostSetId value passed into AddHosts request")
	}
	if len(hostIds) == 0 {
		return nil, errors.New("empty hostIds passed into AddHosts request")
	}
	if c.client == nil {
		return nil, errors.New("nil client")
	}

	opts, apiOpts := getOpts(opt...)

	if version == 0 {
		if !opts.withAutomaticVersioning {
			return nil, errors.New("zero version number passed into AddHosts request")
		}
		existingTarget, existingErr := c.Read(ctx, hostSetId, opt...)
		if existingErr != nil {
			if api.AsServerError(existingErr) != nil {
				return nil, fmt.Errorf("error from controller when performing initial check-and-set read: %w", existingErr)
			}
			return nil, fmt.Errorf("error performing initial check-and-set read: %w", existingErr)
		}
		if existingTarget == nil {
			return nil, errors.New("nil resource response found when performing initial check-and-set read")
		}
		if existingTarget.Item == nil {
			return nil, errors.New("nil resource found when performing initial check-and-set read")
		}
		version = existingTarget.Item.Version
	}

	opts.postMap["version"] = version
	opts.postMap["host_ids"] = hostIds

	req, err := c.client.NewRequest(ctx, "POST", fmt.Sprintf("host-sets/%s:add-hosts", hostSetId), opts.postMap, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating AddHosts request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during AddHosts call: %w", err)
	}

	target := new(HostSetUpdateResult)
	target.Item = new(HostSet)
	apiErr, err := resp.Decode(target.Item)
	if err != nil {
		return nil, fmt.Errorf("error decoding AddHosts response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}
	target.responseBody = resp.Body
	target.responseMap = resp.Map
	return target, nil
}

func (c *Client) SetHosts(ctx context.Context, hostSetId string, version uint32, hostIds []string, opt ...Option) (*HostSetUpdateResult, error) {
	if hostSetId == "" {
		return nil, fmt.Errorf("empty hostSetId value passed into SetHosts request")
	}

	if c.client == nil {
		return nil, errors.New("nil client")
	}

	opts, apiOpts := getOpts(opt...)

	if version == 0 {
		if !opts.withAutomaticVersioning {
			return nil, errors.New("zero version number passed into SetHosts request")
		}
		existingTarget, existingErr := c.Read(ctx, hostSetId, opt...)
		if existingErr != nil {
			if api.AsServerError(existingErr) != nil {
				return nil, fmt.Errorf("error from controller when performing initial check-and-set read: %w", existingErr)
			}
			return nil, fmt.Errorf("error performing initial check-and-set read: %w", existingErr)
		}
		if existingTarget == nil {
			return nil, errors.New("nil resource response found when performing initial check-and-set read")
		}
		if existingTarget.Item == nil {
			return nil, errors.New("nil resource found when performing initial check-and-set read")
		}
		version = existingTarget.Item.Version
	}

	opts.postMap["version"] = version
	opts.postMap["host_ids"] = hostIds

	req, err := c.client.NewRequest(ctx, "POST", fmt.Sprintf("host-sets/%s:set-hosts", hostSetId), opts.postMap, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating SetHosts request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during SetHosts call: %w", err)
	}

	target := new(HostSetUpdateResult)
	target.Item = new(HostSet)
	apiErr, err := resp.Decode(target.Item)
	if err != nil {
		return nil, fmt.Errorf("error decoding SetHosts response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}
	target.responseBody = resp.Body
	target.responseMap = resp.Map
	return target, nil
}

func (c *Client) RemoveHosts(ctx context.Context, hostSetId string, version uint32, hostIds []string, opt ...Option) (*HostSetUpdateResult, error) {
	if hostSetId == "" {
		return nil, fmt.Errorf("empty hostSetId value passed into RemoveHosts request")
	}
	if len(hostIds) == 0 {
		return nil, errors.New("empty hostIds passed into RemoveHosts request")
	}
	if c.client == nil {
		return nil, errors.New("nil client")
	}

	opts, apiOpts := getOpts(opt...)

	if version == 0 {
		if !opts.withAutomaticVersioning {
			return nil, errors.New("zero version number passed into RemoveHosts request")
		}
		existingTarget, existingErr := c.Read(ctx, hostSetId, opt...)
		if existingErr != nil {
			if api.AsServerError(existingErr) != nil {
				return nil, fmt.Errorf("error from controller when performing initial check-and-set read: %w", existingErr)
			}
			return nil, fmt.Errorf("error performing initial check-and-set read: %w", existingErr)
		}
		if existingTarget == nil {
			return nil, errors.New("nil resource response found when performing initial check-and-set read")
		}
		if existingTarget.Item == nil {
			return nil, errors.New("nil resource found when performing initial check-and-set read")
		}
		version = existingTarget.Item.Version
	}

	opts.postMap["version"] = version
	opts.postMap["host_ids"] = hostIds

	req, err := c.client.NewRequest(ctx, "POST", fmt.Sprintf("host-sets/%s:remove-hosts", hostSetId), opts.postMap, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating RemoveHosts request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during RemoveHosts call: %w", err)
	}

	target := new(HostSetUpdateResult)
	target.Item = new(HostSet)
	apiErr, err := resp.Decode(target.Item)
	if err != nil {
		return nil, fmt.Errorf("error decoding RemoveHosts response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}
	target.responseBody = resp.Body
	target.responseMap = resp.Map
	return target, nil
}