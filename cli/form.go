package cli

import (
	"bufio"
	"fmt"
	"github.com/howeyc/gopass"
	"os"
	"strings"
)

var (
	reader = bufio.NewReader(os.Stdin)
)

func AskPassword(label string) string {
	fmt.Printf(label)
	password := gopass.GetPasswd()

	return string(password)
}

func AskOption(options []string) string {
	for {
		fmt.Printf("%v ? ", options)

		answer, _ := reader.ReadString('\n')
		answer = strings.Trim(answer, "\n")

		for _, option := range options {
			if answer == option {
				return answer
			}
		}

	}
}
