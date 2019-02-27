package main

import "os/exec"

func main(){
	cmds := []*exec.Cmd{
		exec.Command("ps","aux"),
	}
}
