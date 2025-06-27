package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/fkmatsuda/contabo-svc/lib/httplib"
	"github.com/fkmatsuda/contabo-svc/snapshot/services/models"
	"github.com/google/uuid"
)

type SnapshotService struct {
	BaseURL     string
	AccessToken string
	httpClient  httplib.HTTPClient
}

func NewSnapshotService(accessToken string) *SnapshotService {
	return &SnapshotService{
		BaseURL:     "https://api.contabo.com",
		AccessToken: accessToken,
		httpClient:  &http.Client{},
	}
}

func (s *SnapshotService) GetInstanceSnapshots(traceID string, instanceID int64) (string, *models.ListSnapshotResponse, error) {
	// Criar a URL completa
	url := fmt.Sprintf("%s/v1/compute/instances/%d/snapshots", s.BaseURL, instanceID)
	requestID := uuid.New().String()

	// Criar nova requisição
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return requestID, nil, fmt.Errorf("erro ao criar requisição: %v", err)
	}

	// Adicionar headers
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", s.AccessToken))
	req.Header.Add("x-request-id", requestID) // Idealmente, gerar um UUID único
	req.Header.Add("x-trace-id", traceID)

	// Fazer a requisição
	client := s.httpClient
	resp, err := client.Do(req)
	if err != nil {
		return requestID, nil, fmt.Errorf("erro ao fazer requisição: %v", err)
	}
	defer resp.Body.Close()

	// Ler o corpo da resposta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return requestID, nil, fmt.Errorf("erro ao ler resposta: %v", err)
	}
	response := &models.ListSnapshotResponse{}
	if err := json.Unmarshal(body, response); err != nil {
		return requestID, nil, fmt.Errorf("erro ao decodificar resposta: %v", err)
	}

	return requestID, response, nil
}

func (s *SnapshotService) DeleteSnapshot(traceID string, instanceID int64, snapshotID string) (string, error) {
	url := fmt.Sprintf("%s/v1/compute/instances/%d/snapshots/%s", s.BaseURL, instanceID, snapshotID)
	requestID := uuid.New().String()

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return requestID, fmt.Errorf("erro ao criar requisição: %v", err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", s.AccessToken))
	req.Header.Add("x-request-id", requestID)
	req.Header.Add("x-trace-id", traceID)

	client := s.httpClient
	resp, err := client.Do(req)
	if err != nil {
		return requestID, fmt.Errorf("erro ao fazer requisição: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		body, _ := io.ReadAll(resp.Body)
		return requestID, fmt.Errorf("erro inesperado: status code %d, body: %s", resp.StatusCode, string(body))
	}

	return requestID, nil
}

func (s *SnapshotService) DeleteOldSnapshots(traceID string, instanceID int64, snapshotsToKeep int) (string, error) {
	requestID, snapshots, err := s.GetInstanceSnapshots(traceID, instanceID)
	if err != nil {
		fmt.Printf("Error getting snapshots: \n\tRequest ID: %s\n", requestID)
		return requestID, fmt.Errorf("erro ao obter snapshots: %v", err)
	}
	if len(snapshots.Data) <= snapshotsToKeep {
		return requestID, nil
	}
	oldSnapshot := snapshots.GetOldestSnapshot()
	if oldSnapshot == nil {
		return requestID, fmt.Errorf("não foi possível encontrar snapshots antigos")
	}
	requestID, err = s.DeleteSnapshot(traceID, instanceID, oldSnapshot.SnapshotID)
	if err != nil {
		fmt.Printf("Error deleting snapshot: \n\tRequest ID: %s\n\tSnapshot ID: %s\n\tTrace ID: %s\n\tError: %v\n", requestID, oldSnapshot.SnapshotID, traceID, err)
		return requestID, fmt.Errorf("erro ao deletar snapshot: %v", err)
	}
	return s.DeleteOldSnapshots(traceID, instanceID, snapshotsToKeep)
}

func (s *SnapshotService) CreateSnapshot(traceID string, instanceID int64, name, description string) (string, *models.CreateSnapshotResponse, error) {
	url := fmt.Sprintf("%s/v1/compute/instances/%d/snapshots", s.BaseURL, instanceID)
	requestID := uuid.New().String()

	payload := models.CreateSnapshotRequest{
		Name:        name,
		Description: description,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return requestID, nil, fmt.Errorf("error marshaling request: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return requestID, nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", s.AccessToken))
	req.Header.Add("x-request-id", requestID)
	req.Header.Add("x-trace-id", traceID)

	client := s.httpClient
	resp, err := client.Do(req)
	if err != nil {
		return requestID, nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return requestID, nil, fmt.Errorf("error reading response body: %v", err)
	}

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		return requestID, nil, fmt.Errorf("unexpected status code %d: %s", resp.StatusCode, string(body))
	}
	response := &models.CreateSnapshotResponse{}
	if err := json.Unmarshal(body, response); err != nil {
		return requestID, nil, fmt.Errorf("error decoding response: %v", err)
	}

	return requestID, response, nil
}
