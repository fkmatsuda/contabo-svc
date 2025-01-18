package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/fkmatsuda/contabo-svc/instance/services/models"
	"github.com/fkmatsuda/contabo-svc/lib/httplib"
	"github.com/google/uuid"
)

type InstanceService struct {
	BaseURL     string
	accessToken string
	httpClient  httplib.HTTPClient
}

func NewInstanceService(accessToken string) *InstanceService {
	return &InstanceService{
		BaseURL:     "https://api.contabo.com",
		accessToken: accessToken,
		httpClient:  &http.Client{},
	}
}
func (s *InstanceService) ListInstances(traceID string) (string, *models.ListInstancesResponse, error) {
	url := fmt.Sprintf("%s/v1/compute/instances", s.BaseURL)
	requestID := uuid.New().String()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return requestID, nil, fmt.Errorf("erro ao criar requisição: %v", err)
	}

	// Adiciona os headers necessários
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", s.accessToken))
	req.Header.Add("x-request-id", requestID)
	req.Header.Add("x-trace-id", traceID)

	// Realiza a requisição
	client := s.httpClient
	resp, err := client.Do(req)
	if err != nil {
		return requestID, nil, fmt.Errorf("erro ao fazer requisição: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return requestID, nil, fmt.Errorf("erro ao ler resposta: %v", err)
	}
	response := &models.ListInstancesResponse{}
	if err := json.Unmarshal(body, response); err != nil {
		return requestID, nil, fmt.Errorf("erro ao decodificar resposta: %v", err)
	}

	return requestID, response, nil
}

func (s *InstanceService) GetInstanceIDByName(traceID, instanceName string) (string, int64, error) {
	requestID, instances, err := s.ListInstances(traceID)
	if err != nil {
		return requestID, 0, fmt.Errorf("erro ao listar instâncias: %v", err)
	}
	for _, instance := range instances.Data {
		if instance.Name == instanceName || instance.DisplayName == instanceName {
			return requestID, instance.InstanceId, nil
		}
	}
	return requestID, 0, fmt.Errorf("instância com o nome '%s' não encontrada", instanceName)
}
