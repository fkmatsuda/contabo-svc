package models

type ListSnapshotPagination struct {
	Size          int `json:"size"`
	TotalElements int `json:"totalElements"`
	TotalPages    int `json:"totalPages"`
	Page          int `json:"page"`
}

type ListSnapshotLinks struct {
	First    string `json:"first"`
	Next     string `json:"next"`
	Self     string `json:"self"`
	Previous string `json:"previous"`
	Last     string `json:"last"`
}

type ListSnapshotResponse struct {
	Pagination ListSnapshotPagination `json:"_pagination"`
	Data       []Snapshot             `json:"data"`
	Links      ListSnapshotLinks      `json:"_links"`
}

func (l *ListSnapshotResponse) GetOldestSnapshot() *Snapshot {
	if len(l.Data) == 0 {
		return nil
	}
	oldest := l.Data[0]
	for _, snapshot := range l.Data {
		if snapshot.CreatedDate < oldest.CreatedDate {
			oldest = snapshot
		}
	}
	return &oldest
}
