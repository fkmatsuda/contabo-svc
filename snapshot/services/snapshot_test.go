package services

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/fkmatsuda/contabo-svc/snapshot/services/models"
	"github.com/stretchr/testify/assert"
)

type MockHTTPClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}

func TestGetInstanceSnapshots(t *testing.T) {
	tests := []struct {
		name           string
		instanceID     int64
		mockResponse   *models.ListSnapshotResponse
		expectedError  bool
		expectedStatus int
	}{
		{
			name:       "successful_response",
			instanceID: 123,
			mockResponse: &models.ListSnapshotResponse{
				Data: []models.Snapshot{
					{
						SnapshotID: "snap-1",
						Name:       "test-snapshot",
					},
				},
			},
			expectedStatus: http.StatusOK,
			expectedError:  false,
		},
		{
			name:           "error_response",
			instanceID:     123,
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

			service := &SnapshotService{
				AccessToken: "test-token",
				httpClient:  mockClient,
			}

			requestID, response, err := service.GetInstanceSnapshots("trace-123", tt.instanceID)

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

func TestDeleteSnapshot(t *testing.T) {
	tests := []struct {
		name           string
		instanceID     int64
		snapshotID     string
		expectedStatus int
		expectedError  bool
	}{
		{
			name:           "successful_deletion",
			instanceID:     123,
			snapshotID:     "snap-1",
			expectedStatus: http.StatusNoContent,
			expectedError:  false,
		},
		{
			name:           "error_deletion",
			instanceID:     123,
			snapshotID:     "snap-1",
			expectedStatus: http.StatusInternalServerError,
			expectedError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := &MockHTTPClient{
				DoFunc: func(req *http.Request) (*http.Response, error) {
					return &http.Response{
						StatusCode: tt.expectedStatus,
						Body:       io.NopCloser(bytes.NewBuffer(nil)),
					}, nil
				},
			}

			service := &SnapshotService{
				AccessToken: "test-token",
				httpClient:  mockClient,
			}

			requestID, err := service.DeleteSnapshot("trace-123", tt.instanceID, tt.snapshotID)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, requestID)
			}
		})
	}
}

func TestCreateSnapshot(t *testing.T) {
	tests := []struct {
		name           string
		instanceID     int64
		snapshotName   string
		description    string
		mockResponse   *models.CreateSnapshotResponse
		expectedStatus int
		expectedError  bool
	}{
		{
			name:         "successful_creation",
			instanceID:   123,
			snapshotName: "test-snapshot",
			description:  "test description",
			mockResponse: &models.CreateSnapshotResponse{
				Data: []models.Snapshot{
					{
						SnapshotID: "snap-1",
						Name:       "test-snapshot",
					},
				},
			},
			expectedStatus: http.StatusCreated,
			expectedError:  false,
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

			service := &SnapshotService{
				AccessToken: "test-token",
				httpClient:  mockClient,
			}

			requestID, response, err := service.CreateSnapshot("trace-123", tt.instanceID, tt.snapshotName, tt.description)

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
