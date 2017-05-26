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
	}{
		{"", "", ""},
		{"a", "b", "c"},
	}

	for _, tt := range tests {
		os.Setenv("ALIYUN_ACCESS_KEY", tt.accessKey)
		os.Setenv("ALIYUN_SECRET_KEY", tt.secretKey)
		os.Setenv("ALIYUN_REGION", tt.region)

		accessKey, secretKey, region := getEnvVars()
		if accessKey != tt.accessKey {
			t.Errorf("accessKey error. got: %s want: %s", accessKey, tt.accessKey)
		}

		if secretKey != tt.secretKey {
			t.Errorf("secretKey error. got: %s want: %s", secretKey, tt.secretKey)
		}

		if region != tt.region {
			t.Errorf("region error. got: %s want: %s", region, tt.region)
		}

		os.Unsetenv("ALIYUN_ACCESS_KEY")
		os.Unsetenv("ALIYUN_SECRET_KEY")
		os.Unsetenv("ALIYUN_REGION")
	}
}
