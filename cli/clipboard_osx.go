// +build darwin

package cli

import (
	"io"
	"os/exec"
)

func CopyToClipboard(text string) {
	command := exec.Command("pbcopy")
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
