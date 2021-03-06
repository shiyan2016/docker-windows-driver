package sdk

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
)

type Proto string

const (
	ProtoTCP Proto = "tcp"
	ProtoNamedPipe Proto = "npipe"
)

func writeSpec(name, address string, proto Proto) (string, error) {

	var pluginSpecDir string
	if runtime.GOOS == "windows" {
		pluginSpecDir = ([]string{filepath.Join(os.Getenv("programdata"), "docker", "plugins")})[0]
	} else {
		pluginSpecDir = "/etc/docker/plugins"
	}

	if err := os.MkdirAll(pluginSpecDir, 0755); err != nil {
		return "", err
	}
	spec := filepath.Join(pluginSpecDir, name+".spec")

	url := string(proto) + "://" + address
	if err := ioutil.WriteFile(spec, []byte(url), 0644); err != nil {
		return "", err
	}

	return spec, nil
}
