package volume

type DriverConfig struct {
	Region    string `json:"region"`
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
	Cluster   string `json:"cluster"`
}

func (c *DriverConfig) HasAliyunCredentials() bool {
	return c.AccessKey != "" && c.SecretKey != "" && c.Region != ""
}
