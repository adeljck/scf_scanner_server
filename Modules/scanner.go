package Modules

import (
	"fmt"
	"os/exec"
	ip2 "scf_scanner_server/Modules/ip"
	"strings"
)

type Scanner struct {
	Type    int    `json:"type" binding:"required"`
	Args    string `json:"args"  binding:"required"`
	Results string
}

func (S *Scanner) Scan() {
	runArgs := strings.Split(S.Args, " ")
	cmd := new(exec.Cmd)
	switch S.Type {
	case FSCAN:
		runArgs = append(runArgs, "-no")
		cmd = exec.Command("./tools/f", runArgs...)
		break
	case KSCAN:
		cmd = exec.Command("./tools/k", runArgs...)
		break
	default:
		return
	}
	out, err := cmd.CombinedOutput()
	if S.Args == "-h" {
		S.Results = string(out)
		return
	}
	if err != nil {
		S.Results = fmt.Sprintf("Execute Comand is %s \ncmd.Run() failed with %s\n", cmd.String(), err)
		return
	}
	var ip ip2.IPInfo
	ip.GetCurrentRequestIPInfo()
	S.Results = S.Results + fmt.Sprintf("requests_ip:%s\nrun command:%s\n\n\n\n", ip.Query, cmd.String()) + string(out)
}
