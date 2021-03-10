package main

type (
	tConfRuntime struct {
		StorageMailPath string
	}
)

var (
	confRuntime tConfRuntime
)

func confRuntimeLoad() {
	iwResponse, _ := iwToolGet("system", "c_system_storage_dir_mailpath")
	confRuntime.StorageMailPath = iwResponse["c_system_storage_dir_mailpath"]
}
