package main

type APIResponse struct {
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Items  []map[string]string
	Paging struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
		Count  int `json:"count"`
		Pages  int `json:"pages"`
	} `json:"paging"`
}

type Domain struct {
	Name string `json:"name"`
	ID   string `json:"uuid"`
	Type string `json:"type"`
}

type DeviceList struct {
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Items []Device

	Paging struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
		Count  int `json:"count"`
		Pages  int `json:"pages"`
	} `json:"paging"`
}

type Device struct {
	DeploymentStatus string `json:"deploymentStatus"`
	ID               string `json:"id"`
	Type             string `json:"type"`
	Links            struct {
		Self string `json:"self"`
	} `json:"links"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Model         string `json:"model"`
	ModelID       string `json:"modelId"`
	ModelNumber   string `json:"modelNumber"`
	ModelType     string `json:"modelType"`
	HealthStatus  string `json:"healthStatus"`
	HealthMessage string `json:"healthMessage"`
	SwVersion     string `json:"sw_version"`
	HealthPolicy  struct {
		ID   string `json:"id"`
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"healthPolicy"`
	AccessPolicy struct {
		Name string `json:"name"`
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"accessPolicy"`
	Advanced struct {
		EnableOGS bool `json:"enableOGS"`
	} `json:"advanced"`
	HostName               string   `json:"hostName"`
	LicenseCaps            []string `json:"license_caps"`
	PerformanceTier        string   `json:"performanceTier"`
	KeepLocalEvents        bool     `json:"keepLocalEvents"`
	ProhibitPacketTransfer bool     `json:"prohibitPacketTransfer"`
	IsConnected            bool     `json:"isConnected"`
	FtdMode                string   `json:"ftdMode"`
	AnalyticsOnly          bool     `json:"analyticsOnly"`
	Metadata               struct {
		ReadOnly struct {
			State bool `json:"state"`
		} `json:"readOnly"`
		InventoryData struct {
			CPUCores   string `json:"cpuCores"`
			CPUType    string `json:"cpuType"`
			MemoryInMB string `json:"memoryInMB"`
		} `json:"inventoryData"`
		DeviceSerialNumber        string `json:"deviceSerialNumber"`
		Domain                    Domain `json:"domain"`
		IsMultiInstance           bool   `json:"isMultiInstance"`
		SnortVersion              string `json:"snortVersion"`
		VdbVersion                string `json:"vdbVersion"`
		LspVersion                string `json:"lspVersion"`
		ClusterBootstrapSupported bool   `json:"clusterBootstrapSupported"`
	} `json:"metadata"`
	SnortEngine string `json:"snortEngine"`
}
