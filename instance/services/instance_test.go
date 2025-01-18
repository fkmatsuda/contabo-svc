package services

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/fkmatsuda/contabo-svc/instance/services/models"
	"github.com/stretchr/testify/assert"
)

type MockHTTPClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}

func TestListInstances(t *testing.T) {
	tests := []struct {
		name           string
		mockResponse   *models.ListInstancesResponse
		expectedError  bool
		expectedStatus int
	}{
		{
			name: "successful_response",
			mockResponse: &models.ListInstancesResponse{
				Data: []models.Instance{
					{
						InstanceId:  123,
						Name:        "test-instance",
						DisplayName: "Test Instance",
					},
				},
			},
			expectedStatus: http.StatusOK,
			expectedError:  false,
		},
		{
			name:           "error_response",
			mockResponse:   nil,
			expectedStatus: http.StatusInternalServerError,
			expectedError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := &MockHTTPClient{
				DoFunc: func(req *http.Request) (*http.Response, error) {
					responseBody := new(bytes.Buffer)
					if tt.mockResponse != nil {
						json.NewEncoder(responseBody).Encode(tt.mockResponse)
					}
					return &http.Response{
						StatusCode: tt.expectedStatus,
						Body:       io.NopCloser(responseBody),
					}, nil
				},
			}

			service := &InstanceService{
				accessToken: "test-token",
				httpClient:  mockClient,
			}

			requestID, response, err := service.ListInstances("trace-123")

			if tt.expectedError {
				assert.Error(t, err)
				assert.Empty(t, response)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, requestID)
				assert.Equal(t, tt.mockResponse, response)
			}
		})
	}
}

func TestGetInstanceIDByName(t *testing.T) {
	tests := []struct {
		name          string
		instanceName  string
		mockInstances *models.ListInstancesResponse
		expectedID    int64
		expectedError bool
	}{
		{
			name:         "found_by_name",
			instanceName: "test-instance",
			mockInstances: &models.ListInstancesResponse{
				Data: []models.Instance{
					{
						InstanceId:  123,
						Name:        "test-instance",
						DisplayName: "Test Instance",
					},
				},
			},
			expectedID:    123,
			expectedError: false,
		},
		{
			name:         "not_found",
			instanceName: "non-existent",
			mockInstances: &models.ListInstancesResponse{
				Data: []models.Instance{
					{
						InstanceId:  123,
						Name:        "test-instance",
						DisplayName: "Test Instance",
					},
				},
			},
			expectedID:    0,
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := &MockHTTPClient{
				DoFunc: func(req *http.Request) (*http.Response, error) {
					responseBody := new(bytes.Buffer)
					json.NewEncoder(responseBody).Encode(tt.mockInstances)
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(responseBody),
					}, nil
				},
			}

			service := &InstanceService{
				BaseURL:     "https://api.contabo.com",
				accessToken: "test-token",
				httpClient:  mockClient,
			}

			requestID, id, err := service.GetInstanceIDByName("trace-123", tt.instanceName)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Equal(t, int64(0), id)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, requestID)
				assert.Equal(t, tt.expectedID, id)
			}
		})
	}
}
