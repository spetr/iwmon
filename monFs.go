package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	monFsMailMkdir = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "fs_mail_mkdir",
			Help: "Filesystem mail sotrage - mkdir (microseconds)",
		},
	)
	monFsMailList = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "fs_mail_list",
			Help: "Filesystem mail sotrage - list (microseconds)",
		},
	)
	monFsMailCreate = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "fs_mail_create",
			Help: "Filesystem mail sotrage - create (microseconds)",
		},
	)
	monFsMailOpen = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "fs_mail_open",
			Help: "Filesystem mail sotrage - open (microseconds)",
		},
	)
	monFsMailLock = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "fs_mail_lock",
			Help: "Filesystem mail sotrage - lock (microseconds)",
		},
	)
	monFsMailWrite = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "fs_mail_write",
			Help: "Filesystem mail sotrage - write (microseconds)",
		},
	)
	monFsMailSync = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "fs_mail_sync",
			Help: "Filesystem mail sotrage - sync (microseconds)",
		},
	)
	monFsMailRead = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "fs_mail_read",
			Help: "Filesystem mail sotrage - read (microseconds)",
		},
	)
	monFsMailClose = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "fs_mail_close",
			Help: "Filesystem mail sotrage - close (microseconds)",
		},
	)
	monFsMailUnlock = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "fs_mail_unlock",
			Help: "Filesystem mail sotrage - unlock (microseconds)",
		},
	)
	monFsMailStat = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "fs_mail_stat",
			Help: "Filesystem mail sotrage - stat (microseconds)",
		},
	)
	monFsMailStatNx = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "fs_mail_statnx",
			Help: "Filesystem mail sotrage - statnx (microseconds)",
		},
	)
	monFsMailDelete = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "fs_mail_delete",
			Help: "Filesystem mail sotrage - delete (microseconds)",
		},
	)
	monFsMailRmdir = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "fs_mail_rmdir",
			Help: "Filesystem mail sotrage - rmdir (microseconds)",
		},
	)
)

func monFsMailUpdate(r *prometheus.Registry) {
	go func(r *prometheus.Registry) {
		var (
			start    time.Time
			err      error
			testPath string
			fh       *os.File
		)

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

		for {
			testPath = path.Join(confRuntime.StorageMailPath, "iwmon")
			// Create iwmon folder
			if err = os.MkdirAll(testPath, os.ModePerm); err != nil {
				fmt.Printf("TODO - ERROR: %s", err.Error())
				time.Sleep(10)
			}

			// Prepare random folder and file
			testFolder := getRandString(16)
			testFile := fmt.Sprintf("%s.dat", getRandString(16))
			fmt.Println("Test folder:", testFolder)
			fmt.Println("Test file:", testFile)

			// mkdir
			start = time.Now()
			if err = os.Mkdir(path.Join(testPath, testFolder), os.ModePerm); err != nil {
				monFsMailMkdir.Set(-1)
				continue
			}
			monFsMailMkdir.Set(float64(time.Since(start).Microseconds()))

			// list
			start = time.Now()
			if _, err = ioutil.ReadDir(path.Join(testPath, testFolder)); err != nil {
				monFsMailList.Set(-1)
				continue
			}
			monFsMailList.Set(float64(time.Since(start).Microseconds()))

			// create file
			start = time.Now()
			if fh, err = os.OpenFile(path.Join(testPath, testFolder, testFile), os.O_RDWR|os.O_CREATE, os.ModePerm); err != nil {
				monFsMailCreate.Set(-1)
				continue
			}
			monFsMailCreate.Set(float64(time.Since(start).Microseconds()))
			fh.Close()

			// open file
			start = time.Now()
			if fh, err = os.OpenFile(path.Join(testPath, testFolder, testFile), os.O_RDWR, os.ModePerm); err != nil {
				monFsMailOpen.Set(-1)
				continue
			}
			monFsMailOpen.Set(float64(time.Since(start).Microseconds()))

			// flock - TODO

			// write
			start = time.Now()
			monFsMailWrite.Set(float64(time.Since(start).Microseconds()))

			// sync
			start = time.Now()
			monFsMailSync.Set(float64(time.Since(start).Microseconds()))

			time.Sleep(conf.IceWarp.Refresh.FsMail * time.Second)

		}
	}(r)
}
