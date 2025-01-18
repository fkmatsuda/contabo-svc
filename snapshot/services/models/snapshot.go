package models

type Snapshot struct {
	TenantID       string `json:"tenantId"`
	CustomerID     string `json:"customerId"`
	SnapshotID     string `json:"snapshotId"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	InstanceID     int    `json:"instanceId"`
	CreatedDate    string `json:"createdDate"`
	AutoDeleteDate string `json:"autoDeleteDate"`
	ImageID        string `json:"imageId"`
	ImageName      string `json:"imageName"`
}
