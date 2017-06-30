package main

import (
	"os"
	"testing"
)

func TestNewMeta(t *testing.T) {
	meta := newMeta()
	if meta.Ui == nil {
		t.Errorf("NewMeta returned nil Ui")
	}

	if meta.Driver == nil {
		t.Errorf("NewMeta returned nil Driver")
	}
}

func TestGetEnvVars(t *testing.T) {
	tests := []struct {
		accessKey string
		secretKey string
		region    string
		cluster   string
	}{
		{"", "", "", ""},
		{"a", "b", "c", "xyz"},
	}

	for _, tt := range tests {
		os.Setenv("ALIYUN_ACCESS_KEY", tt.accessKey)
		os.Setenv("ALIYUN_SECRET_KEY", tt.secretKey)
		os.Setenv("ALIYUN_REGION", tt.region)
		os.Setenv("ALIYUN_CLUSTER", tt.cluster)

		cfg := getDriverConfig()
		if cfg.AccessKey != tt.accessKey {
			t.Errorf("accessKey error. got: %s want: %s", cfg.AccessKey, tt.accessKey)
		}

		if cfg.SecretKey != tt.secretKey {
			t.Errorf("secretKey error. got: %s want: %s", cfg.SecretKey, tt.secretKey)
		}

		if cfg.Region != tt.region {
			t.Errorf("region error. got: %s want: %s", cfg.Region, tt.region)
		}

		if cfg.Cluster != tt.cluster {
			t.Errorf("cluster error. got: %s want: %s", cfg.Cluster, tt.cluster)
		}

		os.Unsetenv("ALIYUN_ACCESS_KEY")
		os.Unsetenv("ALIYUN_SECRET_KEY")
		os.Unsetenv("ALIYUN_REGION")
		os.Unsetenv("ALIYUN_CLUSTER")
	}
}
