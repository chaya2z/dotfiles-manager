package main

import (
	"fmt"
	"os/exec"
)

func link(target, link string) error {
	fmt.Println("link", target, link)

	cmd := exec.Command("ln", "-sr", target, link)
	return cmd.Run()
}
