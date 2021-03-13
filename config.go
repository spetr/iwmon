package main

import (
	"io/ioutil"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type (
	tConf struct {
		IceWarp      tConfIceWarp      `yaml:"icewarp"`
		API          tConfAPI          `yaml:"api"`
		ZabbixSender tConfZabbixSender `yaml:"zabbix-sender"`
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
	tConfZabbixSender struct {
		Hostname string   `yaml:"hostname"`
		Enabled  bool     `yaml:"enabled"`
		Servers  []string `yaml:"servers"`
	}
)

var conf *tConf

func configLoad(configPath string) (err error) {
	hostname, _ := os.Hostname()
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
			ZabbixSender: tConfZabbixSender{
				Enabled:  false,
				Hostname: hostname,
				Servers:  []string{},
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
