package main

import (
	"io/ioutil"
	"time"

	"gopkg.in/yaml.v2"
)

type (
	tConf struct {
		IceWarp tConfIceWarp `yaml:"icewarp"`
		API     tConfAPI     `yaml:"api"`
	}
	tConfIceWarp struct {
		Tool    tConfIceWarpTool    `yaml:"tool"`
		SNMP    tConfIceWarpSNMP    `yaml:"snmp"`
		Refresh tConfIceWarpRefresh `yaml:"refresh"`
	}
	tConfIceWarpTool struct {
		Path        string        `yaml:"path"`
		Timeout     time.Duration `yaml:"timeout"`
		Concurrency int           `yaml:"concurrency"`
	}
	tConfIceWarpSNMP struct {
		Address string        `yaml:"path"`
		Timeout time.Duration `yaml:"timeout"`
	}
	tConfIceWarpRefresh struct {
		Version time.Duration `yaml:"version"`
		FsMail  time.Duration `yaml:"fs_mail"`
	}
	tConfAPI struct {
		Listen     string   `yaml:"listen"`
		ACL        []string `yaml:"acl"`
		Rest       bool     `yaml:"rest"`
		Prometheus bool     `yaml:"prometheus"`
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
				Refresh: tConfIceWarpRefresh{
					Version: 3600,
					FsMail:  60,
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
