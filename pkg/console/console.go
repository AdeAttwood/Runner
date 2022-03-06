// Copyright 2022 Practically.io All rights reserved
//
// Use of this source is governed by a BSD-style
// licence that can be found in the LICENCE file or at
// https://www.practically.io/copyright
package console

import (
	"fmt"
)

// Writes formatted output
func Write(format string, arguments ...interface{}) {
	fmt.Printf(format, arguments...)
}

// Writes formatted output with a new line at the end
func WriteLine(format string, arguments ...interface{}) {
	Write(format, arguments...)
	fmt.Println()
}

// Replaces the current line on the terminal with the new line
func ReplaceLine(format string, arguments ...interface{}) {
	Write("\033[0G"+format+"\033[0K", arguments...)
}
