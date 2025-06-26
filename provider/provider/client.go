package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Client defines the interface for Envoy Gateway API operations.
type Client interface {
	CreateRoute(routeObj map[string]interface{}) error
	DeleteRoute(routeID string) error
	GetRoutes() ([]map[string]interface{}, error)
	CreateCluster(clusterObj map[string]interface{}) error
	DeleteCluster(clusterID string) error
	GetClusters() ([]map[string]interface{}, error)
}

// EnvoyClient implements the Client interface for the Envoy Gateway REST API.
type EnvoyClient struct {
	Endpoint   string
	AuthToken  string
	HTTPClient *http.Client
}

// NewEnvoyClient creates a new EnvoyClient instance.
func NewEnvoyClient(endpoint, authToken string) *EnvoyClient {
	return &EnvoyClient{
		Endpoint:   endpoint,
		AuthToken:  authToken,
		HTTPClient: &http.Client{},
	}
}

// Helper to set auth header if token is present
func (c *EnvoyClient) setAuth(req *http.Request) {
	if c.AuthToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.AuthToken)
	}
}

// CreateRoute creates a new route via Envoy Gateway REST API
func (c *EnvoyClient) CreateRoute(routeObj map[string]interface{}) error {
	url := c.Endpoint + "/v1/routes"
	body, _ := json.Marshal(routeObj)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	c.setAuth(req)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		respData, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to create route: %s", string(respData))
	}
	return nil
}

// DeleteRoute deletes a route via Envoy Gateway REST API
func (c *EnvoyClient) DeleteRoute(routeID string) error {
	url := fmt.Sprintf("%s/v1/routes/%s", c.Endpoint, routeID)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	c.setAuth(req)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		respData, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to delete route: %s", string(respData))
	}
	return nil
}

// GetRoutes fetches all routes via Envoy Gateway REST API
func (c *EnvoyClient) GetRoutes() ([]map[string]interface{}, error) {
	url := c.Endpoint + "/v1/routes"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	c.setAuth(req)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		respData, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to get routes: %s", string(respData))
	}
	var routes []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&routes); err != nil {
		return nil, err
	}
	return routes, nil
}

// CreateCluster creates a new cluster via Envoy Gateway REST API
func (c *EnvoyClient) CreateCluster(clusterObj map[string]interface{}) error {
	url := c.Endpoint + "/v1/clusters"
	body, _ := json.Marshal(clusterObj)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	c.setAuth(req)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		respData, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to create cluster: %s", string(respData))
	}
	return nil
}

// DeleteCluster deletes a cluster via Envoy Gateway REST API
func (c *EnvoyClient) DeleteCluster(clusterID string) error {
	url := fmt.Sprintf("%s/v1/clusters/%s", c.Endpoint, clusterID)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	c.setAuth(req)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		respData, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to delete cluster: %s", string(respData))
	}
	return nil
}

// GetClusters fetches all clusters via Envoy Gateway REST API
func (c *EnvoyClient) GetClusters() ([]map[string]interface{}, error) {
	url := c.Endpoint + "/v1/clusters"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	c.setAuth(req)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		respData, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to get clusters: %s", string(respData))
	}
	var clusters []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&clusters); err != nil {
		return nil, err
	}
	return clusters, nil
}

// RealClientFactory creates a real EnvoyClient from provider config.
func RealClientFactory(ctx context.Context, config Config) (Client, error) {
	return NewEnvoyClient(config.Endpoint, config.Token), nil
}
