package main

import (
	"bytes"
	"errors"
	"fmt"
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

	tool = exec.Command("/opt/icewarp/tool.sh", append([]string{"get", object}, variables...)...)
	tool.Stdout = &toolOutBuf
	tool.Start()

	toolDone = make(chan error)
	go func() { toolDone <- tool.Wait() }()
	select {
	case err = <-toolDone:
	case <-time.After(5 * time.Second):
		tool.Process.Kill()
		err = errors.New("tool.sh - command timed out")
	}

	ret = make(map[string]string, 0)
	for _, line := range strings.Split(toolOutBuf.String(), "\n") {
		lineParsed := strings.SplitN(line, ":", 2)
		if len(lineParsed) == 2 {
			ret[lineParsed[0]] = lineParsed[1]
		} else {
			fmt.Printf("DEBUG: skip line %s", line)
		}
	}
	return
}
