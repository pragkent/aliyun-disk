package flexvolume

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
)

const (
	// StatusSuccess represents the successful completion of command.
	StatusSuccess = "Success"
	// StatusFailed represents that the command failed.
	StatusFailure = "Failed"
	// StatusNotSupported represents that the command is not supported.
	StatusNotSupported = "Not supported"
)

// DriverStatus represents the return value of the driver callout.
type DriverStatus struct {
	// Status of the callout. One of "Success", "Failure" or "Not supported".
	Status string `json:"status"`
	// Reason for success/failure.
	Message string `json:"message,omitempty"`
	// Path to the device attached. This field is valid only for attach calls.
	// ie: /dev/sdx
	DevicePath string `json:"device,omitempty"`
	// Cluster wide unique name of the volume.
	VolumeName string `json:"volumeName,omitempty"`
	// Represents volume is attached on the node
	Attached bool `json:"attached,omitempty"`
}
