package logs

import (
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/glog"
)

func init() {
	flag.CommandLine.Parse(nil)
	flag.Set("stderrthreshold", "FATAL")
	flag.Set("log_dir", "/neverland/nolog")

	glog.Infof("xxx")

	if os.Getenv("ALIYUN_DISK_DEBUG") != "" {
		log.SetOutput(os.Stderr)
	} else {
		log.SetOutput(ioutil.Discard)
	}
}
