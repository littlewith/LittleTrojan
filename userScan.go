//Get the string type in by the keyboard.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ShellScanIn() string {
	fmt.Print(">>>")
	var res string
	reader := bufio.NewReader(os.Stdin)
	msg, err := reader.ReadString('\n')
	res = strings.TrimSpace(msg)
	chkError(err)
	return res
}
