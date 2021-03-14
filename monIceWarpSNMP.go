package main

import (
	"fmt"
	"time"

	"github.com/gosnmp/gosnmp"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spetr/go-zabbix-sender"
)

var (
	monIwSMTPCount = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "iw_smtp_count",
			Help: "",
		},
	)
	monIwPOP3Count = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "iw_pop3_count",
			Help: "",
		},
	)
	monIwIMAPCount = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "iw_imap_count",
			Help: "",
		},
	)
	monIwXMPPSCount = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "iw_xmpps_count",
			Help: "",
		},
	)
	monIwXMPPCCount = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "iw_xmppc_count",
			Help: "",
		},
	)
	monIwGwCount = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "iw_gw_count",
			Help: "",
		},
	)
	monIwWebCount = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "iw_web_count",
			Help: "",
		},
	)
	monIwMsgOutCount = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "iw_msgout_count",
			Help: "",
		},
	)
	monIwMsgInCount = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "iw_msgin_count",
			Help: "",
		},
	)
	monIwMsgFailCount = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "iw_msgfail_count",
			Help: "",
		},
	)
	monIwMsgFailDataCount = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "iw_msgfaildata_count",
			Help: "",
		},
	)
	monIwMsgFailVirusCount = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "iw_msgfailvirus_count",
			Help: "",
		},
	)
	monIwMsgFailCfCount = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "iw_msgfailcf_count",
			Help: "",
		},
	)
	monIwMsgFailCfExtCount = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "iw_msgfailcfext_count",
			Help: "",
		},
	)
	monIwMsgFailRuleCount = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "iw_msgfailrule_count",
			Help: "",
		},
	)
	monIwMsgFailDnsblCount = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "iw_msgfaildnsbl_count",
			Help: "",
		},
	)
	monIwMsgFailIpsCount = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "iw_msgfailips_count",
			Help: "",
		},
	)
	monIwMsgFailSpamCount = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "iw_msgfailspam_count",
			Help: "",
		},
	)
)

// IceWarp version (from tool.sh)
func monIceWarpSNMPUpdate(r *prometheus.Registry) {
	const (
		oidSMTPCount         = ".1.3.6.1.4.1.23736.1.2.1.1.2.8.1"
		oidPOP3Count         = ".1.3.6.1.4.1.23736.1.2.1.1.2.8.2"
		oidIMAPCount         = ".1.3.6.1.4.1.23736.1.2.1.1.2.8.3"
		oidXMPPSCount        = ".1.3.6.1.4.1.23736.1.2.1.1.2.8.4"
		oidXMPPCCount        = ".1.3.6.1.4.1.23736.1.2.1.1.2.10.4"
		oidGwCount           = ".1.3.6.1.4.1.23736.1.2.1.1.2.8.5"
		oidWebCount          = ".1.3.6.1.4.1.23736.1.2.1.1.2.8.7"
		oidMsgOutCount       = ".1.3.6.1.4.1.23736.1.2.1.1.2.16.1"
		oidMsgInCount        = ".1.3.6.1.4.1.23736.1.2.1.1.2.17.1"
		oidMsgFailCount      = ".1.3.6.1.4.1.23736.1.2.1.1.2.18.1"
		oidMsgFailDataCount  = ".1.3.6.1.4.1.23736.1.2.1.1.2.19.1"
		oidMsgFailVirusCount = ".1.3.6.1.4.1.23736.1.2.1.1.2.20.1"
		oidMsgFailCfCount    = ".1.3.6.1.4.1.23736.1.2.1.1.2.21.1"
		oidMsgFailCfExtCount = ".1.3.6.1.4.1.23736.1.2.1.1.2.22.1"
		oidMsgFailRuleCount  = ".1.3.6.1.4.1.23736.1.2.1.1.2.23.1"
		oidMsgFailDnsblCount = ".1.3.6.1.4.1.23736.1.2.1.1.2.24.1"
		oidMsgFailIpsCount   = ".1.3.6.1.4.1.23736.1.2.1.1.2.25.1"
		oidMsgFailSpamCount  = ".1.3.6.1.4.1.23736.1.2.1.1.2.26.1"
	)

	var (
		err                    error
		snmpResponse           *gosnmp.SnmpPacket
		valueSMTPCount         float64
		valuePOP3Count         float64
		valueIMAPCount         float64
		valueXMPPSCount        float64
		valueXMPPCCount        float64
		valueGwCount           float64
		valueWebCount          float64
		valueMsgOutCount       float64
		valueMsgInCount        float64
		valueMsgFailCount      float64
		valueMsgFailDataCount  float64
		valueMsgFailVirusCount float64
		valueMsgFailCfCount    float64
		valueMsgFailCfExtCount float64
		valueMsgFailRuleCount  float64
		valueMsgFailDnsblCount float64
		valueMsgFailIpsCount   float64
		valueMsgFailSpamCount  float64
	)

	if conf.API.Prometheus {
		r.MustRegister(monIwSMTPCount)
		r.MustRegister(monIwPOP3Count)
		r.MustRegister(monIwIMAPCount)
		r.MustRegister(monIwXMPPSCount)
		r.MustRegister(monIwXMPPCCount)
		r.MustRegister(monIwGwCount)
		r.MustRegister(monIwWebCount)
		r.MustRegister(monIwMsgOutCount)
		r.MustRegister(monIwMsgInCount)
		r.MustRegister(monIwMsgFailCount)
		r.MustRegister(monIwMsgFailDataCount)
		r.MustRegister(monIwMsgFailVirusCount)
		r.MustRegister(monIwMsgFailCfCount)
		r.MustRegister(monIwMsgFailCfExtCount)
		r.MustRegister(monIwMsgFailRuleCount)
		r.MustRegister(monIwMsgFailIpsCount)
		r.MustRegister(monIwMsgFailSpamCount)
	}

	for {
		func() {
			// Default values (no test / error in test)
			valueSMTPCount = -1
			valuePOP3Count = -1
			valueIMAPCount = -1
			valueXMPPSCount = -1
			valueXMPPCCount = -1
			valueGwCount = -1
			valueWebCount = -1
			valueMsgOutCount = -1
			valueMsgInCount = -1
			valueMsgFailCount = -1
			valueMsgFailDataCount = -1
			valueMsgFailVirusCount = -1
			valueMsgFailCfCount = -1
			valueMsgFailCfExtCount = -1
			valueMsgFailRuleCount = -1
			valueMsgFailDnsblCount = -1
			valueMsgFailIpsCount = -1
			valueMsgFailSpamCount = -1

			gosnmp.Default.Target = conf.IceWarp.SNMP.Address
			if err = gosnmp.Default.Connect(); err != nil {
				logger.Errorf("Can not connect to IceWarp SNMP: %s", err.Error())
				time.Sleep(10 * time.Second)
				return
			}
			defer gosnmp.Default.Conn.Close()

			oids := []string{
				oidSMTPCount,
				oidPOP3Count,
				oidIMAPCount,
				oidXMPPSCount,
				oidXMPPCCount,
				oidGwCount,
				oidWebCount,
				oidMsgOutCount,
				oidMsgInCount,
				oidMsgFailCount,
				oidMsgFailDataCount,
				oidMsgFailVirusCount,
				oidMsgFailCfCount,
				oidMsgFailCfExtCount,
				oidMsgFailRuleCount,
				oidMsgFailDnsblCount,
				oidMsgFailIpsCount,
				oidMsgFailSpamCount,
			}
			if snmpResponse, err = gosnmp.Default.Get(oids); err != nil { // Get() accepts up to g.MAX_OIDS
				logger.Errorf("IceWarp SNMP Get() error: %s", err.Error())
				time.Sleep(10 * time.Second)
				return
			}
			for i := range snmpResponse.Variables {
				switch snmpResponse.Variables[i].Name {
				case oidSMTPCount:
					valueSMTPCount = 11
				case oidPOP3Count:
					valuePOP3Count = 11
				case oidIMAPCount:
					valueIMAPCount = 11
				case oidXMPPSCount:
					valueXMPPSCount = 11
				case oidXMPPCCount:
					valueXMPPSCount = 11
				case oidGwCount:
					valueGwCount = 11
				case oidWebCount:
					valueWebCount = 11
				case oidMsgOutCount:
					valueMsgOutCount = 11
				case oidMsgInCount:
					valueMsgInCount = 11
				case oidMsgFailCount:
					valueMsgFailCount = 11
				case oidMsgFailDataCount:
					valueMsgFailDataCount = 11
				case oidMsgFailVirusCount:
					valueMsgFailVirusCount = 11
				case oidMsgFailCfCount:
					valueMsgFailCfCount = 11
				case oidMsgFailCfExtCount:
					valueMsgFailCfExtCount = 11
				case oidMsgFailRuleCount:
					valueMsgFailRuleCount = 11
				case oidMsgFailDnsblCount:
					valueMsgFailDnsblCount = 11
				case oidMsgFailIpsCount:
					valueMsgFailIpsCount = 11
				case oidMsgFailSpamCount:
					valueMsgFailSpamCount = 11
				default:
					logger.Warningf("Unknown SNMP OID in response: %s:", snmpResponse.Variables[i].Name)
				}
			}

		}()

		// Prometheus Exporter
		if conf.API.Prometheus {
			monIwSMTPCount.Set(valueSMTPCount)
			monIwPOP3Count.Set(valuePOP3Count)
			monIwIMAPCount.Set(valueIMAPCount)
			monIwXMPPCCount.Set(valueXMPPCCount)
			monIwXMPPSCount.Set(valueXMPPSCount)
			monIwGwCount.Set(valueGwCount)
			monIwWebCount.Set(valueWebCount)
			monIwMsgOutCount.Set(valueMsgOutCount)
			monIwMsgInCount.Set(valueMsgInCount)
			monIwMsgFailCount.Set(valueMsgFailCount)
			monIwMsgFailDataCount.Set(valueMsgFailDataCount)
			monIwMsgFailVirusCount.Set(valueMsgFailVirusCount)
			monIwMsgFailCfCount.Set(valueMsgFailCfCount)
			monIwMsgFailCfExtCount.Set(valueMsgFailCfExtCount)
			monIwMsgFailRuleCount.Set(valueMsgFailRuleCount)
			monIwMsgFailDnsblCount.Set(valueMsgFailDnsblCount)
			monIwMsgFailIpsCount.Set(valueMsgFailIpsCount)
			monIwMsgFailSpamCount.Set(valueMsgFailSpamCount)
		}

		// Zabbix Sender
		if conf.ZabbixSender.Enabled {
			var (
				metrics []*zabbix.Metric
				t       = time.Now().Unix()
			)
			metrics = append(metrics, zabbix.NewMetric(conf.ZabbixSender.Hostname, "iw_smtp_count", fmt.Sprintf("%f", valueSMTPCount), true, t))
			for i := range conf.ZabbixSender.Servers {
				z := zabbix.NewSender(conf.ZabbixSender.Servers[i])
				z.SendMetrics(metrics)
			}
		}

		time.Sleep(conf.IceWarp.Refresh.SNMP * time.Second)

	}
}
