package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type (
	tConf struct {
		IceWarp tConfIceWarp `yaml:"icewarp"`
		API     tConfAPI     `yaml:"api"`
	}
	tConfIceWarp struct {
		Tool tConfIceWarpTool `yaml:"tool"`
	}
	tConfIceWarpTool struct {
		Path        string `yaml:"path"`
		Timeout     uint   `yaml:"timeout"`
		Concurrency uint   `yaml:"concurrency"`
	}
	tConfAPI struct {
		Listen     string `yaml:"listen"`
		Rest       bool   `yaml:"rest"`
		Prometheus bool   `yaml:"prometheus"`
	}
)

var conf *tConf

func configLoad(configPath string) (err error) {
	var (
		buf     []byte
		newConf = &tConf{
			IceWarp: tConfIceWarp{
				Tool: tConfIceWarpTool{
					Path:        "/opt/icewarp/tool.sh",
					Timeout:     3,
					Concurrency: 2,
				},
			},
			API: tConfAPI{
				Listen:     "0.0.0.0:9090",
				Rest:       false,
				Prometheus: false,
			},
		}
	)
	if buf, err = ioutil.ReadFile(configPath); err != nil {
		return
	}
	if err = yaml.Unmarshal(buf, newConf); err != nil {
		return
	}
	conf = newConf
	return
}
