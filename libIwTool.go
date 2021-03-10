package main

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"
	"time"
)

func iwToolGet(object string, variables ...string) (ret map[string]string, err error) {
	var (
		tool       *exec.Cmd
		toolDone   chan error
		toolOutBuf bytes.Buffer
	)

	tool = exec.Command(conf.IceWarp.Tool.Path, append([]string{"get", object}, variables...)...)
	tool.Stdout = &toolOutBuf
	tool.Start()

	toolDone = make(chan error)
	go func() { toolDone <- tool.Wait() }()
	select {
	case err = <-toolDone:
	case <-time.After(conf.IceWarp.Tool.Timeout * time.Second):
		tool.Process.Kill()
		err = errors.New("tool.sh - command timed out")
	}

	ret = make(map[string]string, 0)
	for _, line := range strings.Split(toolOutBuf.String(), "\n") {
		lineParsed := strings.SplitN(strings.Trim(line, " \t\r\n"), ":", 2)
		if len(lineParsed) != 2 {
			continue
		}
		ret[lineParsed[0]] = strings.Trim(lineParsed[1], " \t")
	}
	return
}
