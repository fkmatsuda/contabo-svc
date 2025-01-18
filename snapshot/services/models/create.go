package models

type CreateSnapshotRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateSnapshotResponse struct {
	Data  []Snapshot        `json:"data"`
	Links ListSnapshotLinks `json:"_links"`
}

type CreateSnapshotLinks struct {
	Self string `json:"self"`
}
