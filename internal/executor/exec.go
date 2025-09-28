package executor

import (
	"os/exec"
)

// 执行命令
func ExecCommand(command string) (string, error) {
	cmd := exec.Command("bash", "-c", command)
	out, err := cmd.CombinedOutput()
	return string(out), err
}
