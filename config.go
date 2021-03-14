package main

import (
	"io/ioutil"
	"net"
	"os"
	"runtime"
	"strconv"
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
		Enabled bool   `yaml:"enabled"`
		Address string `yaml:"path"`
		parsed  struct {
			address string
			port    uint16
		}
		Timeout time.Duration `yaml:"timeout"`
	}
	tConfIceWarpRefresh struct {
		Version int `yaml:"version"`
		FsMail  int `yaml:"fs_mail"`
		SNMP    int `yaml:"snmp"`
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
				SNMP: tConfIceWarpSNMP{
					Enabled: true,
					Address: "127.0.0.1:161",
					Timeout: 5,
				},
				Refresh: tConfIceWarpRefresh{
					Version: 3600,
					FsMail:  60,
					SNMP:    60,
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
	if runtime.GOOS == "windows" {
		newConf.IceWarp.Tool.Path = "C:/IceWarp/tool.exe"
	}
	if buf, err = ioutil.ReadFile(configPath); err != nil {
		return
	}
	if err = yaml.Unmarshal(buf, newConf); err != nil {
		return
	}

	// Parse SNMP
	var (
		snmpPortString string
		snmpPort       int
	)
	newConf.IceWarp.SNMP.parsed.address, snmpPortString, err = net.SplitHostPort(newConf.IceWarp.SNMP.Address)
	if err != nil {
		logger.Errorf("Can not parse SNMP address: %s", err.Error())
		newConf.IceWarp.SNMP.parsed.address = "127.0.0.1"
		snmpPortString = "161"
	}
	snmpPort, err = strconv.Atoi(snmpPortString)
	if err != nil {
		logger.Errorf("Can not parse SNMP port: %s", err.Error())
		snmpPort = 161
	}
	newConf.IceWarp.SNMP.parsed.port = uint16(snmpPort)

	conf = newConf

	return
}
