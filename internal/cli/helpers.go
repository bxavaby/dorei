// ./../cli/helpers.go

package cli

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const (
	green  = "\033[46m"
	red    = "\033[160m"
	blue   = "\033[27m"
	yellow = "\033[226m"
	reset  = "\033[0m"
)

func greatSuccess(msg string) {
	fmt.Printf("%s✓ %s%s\n", green, msg, reset)
}

func ohNoes(msg string, err error) {
	fmt.Fprintf(os.Stderr, "%s✗ %s %s%s\n", red, msg, err, reset)
}

func singleWell(msg string) {
	fmt.Printf("%s→ %s%s\n", blue, msg, reset)
}

func tripleWell(msg string) {
	fmt.Printf("%s⚠ %s%s\n", yellow, msg, reset)
}

func YesOrNo(question string) bool {
	fmt.Printf(question + " (y/N): ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		answer := strings.ToLower(strings.TrimSpace(scanner.Text()))
		return answer == "y" || answer == "yes"
	}
	return false
}

func Wiper() error {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
