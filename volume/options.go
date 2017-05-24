package volume

import (
	"encoding/json"
	"fmt"
)

const (
	// Option keys
	optionFSType    = "kubernetes.io/fsType"
	optionReadWrite = "kubernetes.io/readwrite"
	optionKeySecret = "kubernetes.io/secret"
	optionFSGroup   = "kubernetes.io/fsGroup"
	optionMountsDir = "kubernetes.io/mountsDir"

	optionKeyPodName      = "kubernetes.io/pod.name"
	optionKeyPodNamespace = "kubernetes.io/pod.namespace"
	optionKeyPodUID       = "kubernetes.io/pod.uid"

	optionKeyServiceAccountName = "kubernetes.io/serviceAccount.name"

	// Custom keys
	optionDiskId = "diskId"
)

var requiredOptions = []string{"diskId"}

type Options map[string]interface{}

func ParseOptions(s string) (Options, error) {
	var o Options
	if err := json.Unmarshal([]byte(s), &o); err != nil {
		return nil, errInvalidVolumeOptions
	}

	if err := o.check(); err != nil {
		return nil, err
	}

	return o, nil
}

func (o Options) check() error {
	for _, k := range requiredOptions {
		if _, ok := o[k]; !ok {
			return fmt.Errorf("option %s is required", k)
		}
	}

	return nil
}

func (o Options) DiskId() string {
	diskId, _ := o["diskId"].(string)
	return diskId
}
