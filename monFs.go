package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/spetr/go-zabbix-sender"
)

var (
	monFsMailMkdir = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "fs_mail_mkdir",
			Help: "Filesystem mail storage - mkdir (microseconds)",
		},
	)
	monFsMailList = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "fs_mail_list",
			Help: "Filesystem mail storage - list (microseconds)",
		},
	)
	monFsMailCreate = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "fs_mail_create",
			Help: "Filesystem mail storage - create (microseconds)",
		},
	)
	monFsMailOpen = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "fs_mail_open",
			Help: "Filesystem mail storage - open (microseconds)",
		},
	)
	monFsMailWrite = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "fs_mail_write",
			Help: "Filesystem mail storage - write (microseconds)",
		},
	)
	monFsMailSync = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "fs_mail_sync",
			Help: "Filesystem mail storage - sync (microseconds)",
		},
	)
	monFsMailRead = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "fs_mail_read",
			Help: "Filesystem mail storage - read (microseconds)",
		},
	)
	monFsMailClose = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "fs_mail_close",
			Help: "Filesystem mail storage - close (microseconds)",
		},
	)
	monFsMailStat = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "fs_mail_stat",
			Help: "Filesystem mail storage - stat (microseconds)",
		},
	)
	monFsMailStatNx = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "fs_mail_statnx",
			Help: "Filesystem mail storage - statnx (microseconds)",
		},
	)
	monFsMailDelete = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "fs_mail_delete",
			Help: "Filesystem mail storage - delete (microseconds)",
		},
	)
	monFsMailRmdir = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "fs_mail_rmdir",
			Help: "Filesystem mail storage - rmdir (microseconds)",
		},
	)
)

func monFsMailUpdate(r *prometheus.Registry) {

	var (
		start               time.Time
		err                 error
		testPath            string
		fh                  *os.File
		buffer              []byte
		timeMonFsMailMkdir  float64
		timeMonFsMailList   float64
		timeMonFsMailCreate float64
		timeMonFsMailOpen   float64
		timeMonFsMailWrite  float64
		timeMonFsMailSync   float64
		timeMonFsMailRead   float64
		timeMonFsMailClose  float64
		timeMonFsMailStat   float64
		timeMonFsMailStatNx float64
		timeMonFsMailDelete float64
		timeMonFsMailRmdir  float64
	)

	if conf.API.Prometheus {
		r.MustRegister(monFsMailMkdir)
		r.MustRegister(monFsMailList)
		r.MustRegister(monFsMailCreate)
		r.MustRegister(monFsMailOpen)
		r.MustRegister(monFsMailWrite)
		r.MustRegister(monFsMailSync)
		r.MustRegister(monFsMailRead)
		r.MustRegister(monFsMailClose)
		r.MustRegister(monFsMailStat)
		r.MustRegister(monFsMailStatNx)
		r.MustRegister(monFsMailDelete)
		r.MustRegister(monFsMailRmdir)
	}

	for {
		func() {
			// Default values (no test / error in test)
			timeMonFsMailMkdir = -1
			timeMonFsMailList = -1
			timeMonFsMailCreate = -1
			timeMonFsMailOpen = -1
			timeMonFsMailWrite = -1
			timeMonFsMailSync = -1
			timeMonFsMailRead = -1
			timeMonFsMailClose = -1
			timeMonFsMailStat = -1
			timeMonFsMailStatNx = -1
			timeMonFsMailDelete = -1
			timeMonFsMailRmdir = -1

			if _, err = os.Stat(confRuntime.StorageMailPath); err != nil {
				logger.Errorf("Mail path error: %s", err.Error())
				time.Sleep(10 * time.Second)
				return
			}

			testPath = path.Join(confRuntime.StorageMailPath, "iwmon")
			// Create iwmon folder and prepare data
			if err = os.MkdirAll(testPath, os.ModePerm); err != nil {
				logger.Errorf("Can not create mail fs testing directort: %s", err.Error())
				time.Sleep(10 * time.Second)
				return
			}
			testFolder := getRandString(16)
			testFile := fmt.Sprintf("%s.dat", getRandString(16))
			buffer = []byte(getRandString(8192))

			defer func() {
				if fh != nil {
					fh.Close()
				}
			}()

			// mkdir
			start = time.Now()
			if err = os.Mkdir(path.Join(testPath, testFolder), os.ModePerm); err != nil {
				return
			}
			timeMonFsMailMkdir = float64(time.Since(start).Microseconds())

			// list
			start = time.Now()
			if _, err = ioutil.ReadDir(path.Join(testPath, testFolder)); err != nil {
				return
			}
			timeMonFsMailList = float64(time.Since(start).Microseconds())

			// create file
			start = time.Now()
			if fh, err = os.OpenFile(path.Join(testPath, testFolder, testFile), os.O_RDWR|os.O_CREATE, os.ModePerm); err != nil {
				return
			}
			timeMonFsMailCreate = float64(time.Since(start).Microseconds())
			fh.Close()

			// open file
			start = time.Now()
			if fh, err = os.OpenFile(path.Join(testPath, testFolder, testFile), os.O_RDWR, os.ModePerm); err != nil {
				return
			}
			timeMonFsMailOpen = float64(time.Since(start).Microseconds())

			// flock - TODO

			// write
			fh.SetWriteDeadline(time.Now().Add(2 * time.Second))
			start = time.Now()
			if _, err = fh.Write(buffer); err != nil {
				return
			}
			timeMonFsMailWrite = float64(time.Since(start).Microseconds())

			// sync
			fh.SetWriteDeadline(time.Now().Add(2 * time.Second))
			start = time.Now()
			if err = fh.Sync(); err != nil {
				return
			}
			timeMonFsMailSync = float64(time.Since(start).Microseconds())

			// read
			fh.SetReadDeadline(time.Now().Add(2 * time.Second))
			start = time.Now()
			if _, err = fh.ReadAt(buffer, 0); err != nil {
				return
			}
			timeMonFsMailRead = float64(time.Since(start).Microseconds())

			// close
			fh.SetWriteDeadline(time.Now().Add(2 * time.Second))
			start = time.Now()
			if err = fh.Close(); err != nil {
				return
			}
			timeMonFsMailClose = float64(time.Since(start).Microseconds())

			// stat
			start = time.Now()
			if _, err = os.Stat(path.Join(testPath, testFolder, testFile)); err != nil {
				return
			}
			timeMonFsMailStat = float64(time.Since(start).Microseconds())

			// statnx
			start = time.Now()
			_, _ = os.Stat(path.Join(testPath, testFolder, "non-existing.dat"))
			timeMonFsMailStatNx = float64(time.Since(start).Microseconds())

			// delete file
			start = time.Now()
			if err = os.Remove(path.Join(testPath, testFolder, testFile)); err != nil {
				return
			}
			timeMonFsMailDelete = float64(time.Since(start).Microseconds())

			// delete directory
			start = time.Now()
			if err = os.Remove(path.Join(testPath, testFolder)); err != nil {
				return
			}
			timeMonFsMailRmdir = float64(time.Since(start).Microseconds())

		}()

		// Prometheus Exporter
		if conf.API.Prometheus {
			monFsMailMkdir.Set(timeMonFsMailMkdir)
			monFsMailList.Set(timeMonFsMailList)
			monFsMailCreate.Set(timeMonFsMailCreate)
			monFsMailOpen.Set(timeMonFsMailOpen)
			monFsMailWrite.Set(timeMonFsMailWrite)
			monFsMailSync.Set(timeMonFsMailSync)
			monFsMailRead.Set(timeMonFsMailRead)
			monFsMailClose.Set(timeMonFsMailClose)
			monFsMailStat.Set(timeMonFsMailStat)
			monFsMailStatNx.Set(timeMonFsMailStatNx)
			monFsMailDelete.Set(timeMonFsMailDelete)
			monFsMailRmdir.Set(timeMonFsMailRmdir)
		}

		// Zabbix Sender
		if conf.ZabbixSender.Enabled {
			var (
				metrics []*zabbix.Metric
				t       = time.Now().Unix()
			)
			metrics = append(metrics, zabbix.NewMetric(conf.ZabbixSender.Hostname, "fs.mail_mkdir", fmt.Sprintf("%f", timeMonFsMailMkdir), true, t))
			metrics = append(metrics, zabbix.NewMetric(conf.ZabbixSender.Hostname, "fs.mail_list", fmt.Sprintf("%f", timeMonFsMailList), true, t))
			metrics = append(metrics, zabbix.NewMetric(conf.ZabbixSender.Hostname, "fs.mail_create", fmt.Sprintf("%f", timeMonFsMailCreate), true, t))
			metrics = append(metrics, zabbix.NewMetric(conf.ZabbixSender.Hostname, "fs.mail_open", fmt.Sprintf("%f", timeMonFsMailOpen), true, t))
			metrics = append(metrics, zabbix.NewMetric(conf.ZabbixSender.Hostname, "fs.mail_write", fmt.Sprintf("%f", timeMonFsMailWrite), true, t))
			metrics = append(metrics, zabbix.NewMetric(conf.ZabbixSender.Hostname, "fs.mail_sync", fmt.Sprintf("%f", timeMonFsMailSync), true, t))
			metrics = append(metrics, zabbix.NewMetric(conf.ZabbixSender.Hostname, "fs.mail_read", fmt.Sprintf("%f", timeMonFsMailRead), true, t))
			metrics = append(metrics, zabbix.NewMetric(conf.ZabbixSender.Hostname, "fs.mail_close", fmt.Sprintf("%f", timeMonFsMailClose), true, t))
			metrics = append(metrics, zabbix.NewMetric(conf.ZabbixSender.Hostname, "fs.mail_stat", fmt.Sprintf("%f", timeMonFsMailStat), true, t))
			metrics = append(metrics, zabbix.NewMetric(conf.ZabbixSender.Hostname, "fs.mail_statnx", fmt.Sprintf("%f", timeMonFsMailStatNx), true, t))
			metrics = append(metrics, zabbix.NewMetric(conf.ZabbixSender.Hostname, "fs.mail_delete", fmt.Sprintf("%f", timeMonFsMailDelete), true, t))
			metrics = append(metrics, zabbix.NewMetric(conf.ZabbixSender.Hostname, "fs.mail_rmdir", fmt.Sprintf("%f", timeMonFsMailRmdir), true, t))
			for i := range conf.ZabbixSender.Servers {
				z := zabbix.NewSender(conf.ZabbixSender.Servers[i])
				z.SendMetrics(metrics)
			}
		}

		time.Sleep(conf.IceWarp.Refresh.FsMail * time.Second)
	}
}
