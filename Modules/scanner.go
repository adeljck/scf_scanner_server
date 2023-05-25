package Modules

import (
	"fmt"
	"os/exec"
)

type Scanner struct {
	Ip        string `form:"ip" json:"ip"`
	Ports     string `form:"port" json:"ports"`
	ExecParam string `form:"execParam" json:"execParam"`
	Results   string
}

func (S *Scanner) FScan() {
	cmd := exec.Command("./tools/f", "-h", S.Ip, "-p", S.Ports, "-t", "2048", "-time", "5", "-full", "-np", "-no")
	if S.Ports == "" {
		cmd = exec.Command("./tools/f", "-h", S.Ip, "-t", "2048", "-time", "5", "-full", "-np", "-no")
	}
	out, err := cmd.CombinedOutput()
	if err != nil {
		S.Results = fmt.Sprintf("cmd.Run() failed with %s\n", err)
	}
	S.Results = S.Results + string(out)
}

func (S *Scanner) KScan() {
	cmd := exec.Command("./tools/k", "-t", S.Ip, "-p", S.Ports, "--threads", "2048", "--timeout", "5")
	if S.Ports == "" {
		cmd = exec.Command("./tools/k", "-t", S.Ip, "--threads", "2048", "--timeout", "5")
	}
	out, err := cmd.CombinedOutput()
	if err != nil {
		S.Results = ""
		S.Results = fmt.Sprintf("cmd.Run() failed with %s\n", err)
	}
	S.Results = S.Results + string(out)
}
