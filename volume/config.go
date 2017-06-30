package volume

type DriverConfig struct {
	Region    string
	AccessKey string
	SecretKey string
	Cluster   string
}

func (c *DriverConfig) HasAliyunCredentials() bool {
	return c.AccessKey != "" && c.SecretKey != "" && c.Region != ""
}
