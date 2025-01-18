package models

type ipConfigDetails struct {
	Ip          string `json:"ip"`
	NetmaskCidr int    `json:"netmaskCidr"`
	Gateway     string `json:"gateway"`
}

type ipConfig struct {
	V4 ipConfigDetails `json:"v4"`
	V6 ipConfigDetails `json:"v6"`
}

type Pagination struct {
	Size          int `json:"size"`
	TotalElements int `json:"totalElements"`
	TotalPages    int `json:"totalPages"`
	Page          int `json:"page"`
}

type Links struct {
	First    string `json:"first"`
	Previous string `json:"previous"`
	Self     string `json:"self"`
	Next     string `json:"next"`
	Last     string `json:"last"`
}

type Addon struct {
	ID       int64 `json:"id"`
	Quantity int64 `json:"quantity"`
}

type Instance struct {
	TenantId      string   `json:"tenantId"`
	CustomerId    string   `json:"customerId"`
	AdditionalIps []string `json:"additionalIps"`
	Name          string   `json:"name"`
	DisplayName   string   `json:"displayName"`
	InstanceId    int64    `json:"instanceId"`
	DataCenter    string   `json:"dataCenter"`
	Region        string   `json:"region"`
	RegionName    string   `json:"regionName"`
	ProductId     string   `json:"productId"`
	ImageId       string   `json:"imageId"`
	IpConfig      ipConfig `json:"ipConfig"`
	MacAddress    string   `json:"macAddress"`
	RamMb         int32    `json:"ramMb"`
	CpuCores      int16    `json:"cpuCores"`
	OsType        string   `json:"osType"`
	DiskMb        int64    `json:"diskMb"`
	SshKeys       []string `json:"sshKeys"`
	CreatedDate   string   `json:"createdDate"`
	CancelDate    string   `json:"cancelDate"`
	Status        string   `json:"status"`
	VHostId       int64    `json:"vHostId"`
	VHostNumber   int32    `json:"vHostNumber"`
	VHostName     string   `json:"vHostName"`
	AddOns        []Addon  `json:"addOns"`
	ErrorMessage  string   `json:"errorMessage"`
	ProductType   string   `json:"productType"`
	ProductName   string   `json:"productName"`
	DefaultUser   string   `json:"defaultUser"`
}

type ListInstancesResponse struct {
	Pagination Pagination `json:"_pagination"`
	Data       []Instance `json:"data"`
	Links      Links      `json:"_links"`
}
