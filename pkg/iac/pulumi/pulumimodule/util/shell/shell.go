package shell

import (
	"io"
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

func RunCmd(cmd *exec.Cmd) error {
	log.Debugf("running command %s", cmd.String())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}

func RunCmdWithStdOut(cmd *exec.Cmd, stdOut io.Writer) error {
	if stdOut != nil {
		log.Infof("running command %s", cmd.String())
	}
	cmd.Stdout = stdOut
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}
