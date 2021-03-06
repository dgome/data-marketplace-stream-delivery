//
// Copyright (c) 2019 LG Electronics Inc.
// SPDX-License-Identifier: Apache-2.0
//

package utils

import (
	"fmt"
	"io"
	"log"
	"os"
)

//CreateLogger creates a new Logger
func CreateLogger(name string) *log.Logger {

	return log.New(io.MultiWriter(os.Stdout), fmt.Sprintf("%s: ", name), log.Lshortfile|log.LstdFlags)

}
