// Copyright (c) 2015 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package torchlog

import (
	"fmt"
	"log"
	"time"

	"github.com/fatih/color"
)

var toRedString = color.New(color.FgRed).SprintFunc()
var toBlueString = color.New(color.FgBlue).SprintFunc()

func init() {
	log.SetFlags(0) // disable default flags
}

// Fatalf wraps log.Fatalf and adds the current time and color.
func Fatalf(format string, v ...interface{}) {
	currentTime := time.Now().Format("15:04:05")
	prefix := toRedString(fmt.Sprintf("FATA[%s] ", currentTime))
	log.Fatalf(prefix+format, v...)
}

// Printf wraps log.Printf and adds the current time and color.
func Printf(format string, v ...interface{}) {
	currentTime := time.Now().Format("15:04:05")
	prefix := toBlueString(fmt.Sprintf("INFO[%s] ", currentTime))
	log.Printf(prefix+format, v...)
}

// Print wraps log.Print and adds the current time and color.
func Print(v ...interface{}) {
	currentTime := time.Now().Format("15:04:05")
	prefix := toBlueString(fmt.Sprintf("INFO[%s] ", currentTime))
	log.Print(prefix + fmt.Sprint(v...))
}
