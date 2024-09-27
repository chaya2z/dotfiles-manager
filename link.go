package main

import (
	"fmt"
	"os/exec"
)

func link(source, target string) error {
	fmt.Println("link", source, target)

	cmd := exec.Command("ln", "-sr", source, target)
	return cmd.Run()
}
