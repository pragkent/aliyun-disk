package command

import (
	"encoding/json"

	"github.com/pragkent/aliyun-disk/volume"
)

func jsonify(status *volume.DriverStatus) string {
	b, _ := json.Marshal(status)
	return string(b)
}
