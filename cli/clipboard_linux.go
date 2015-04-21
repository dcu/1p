// +build linux

package cli

import (
	"io"
	"os/exec"
)

func CopyToClipboard(text string) {
	command := exec.Command("xclip", "-selection", "clipboard")
	stdin, err := command.StdinPipe()
	if err != nil {
		panic(err)
	}
	defer stdin.Close()

	if err = command.Start(); err != nil {
		panic(err)
	}

	io.WriteString(stdin, text)
}
