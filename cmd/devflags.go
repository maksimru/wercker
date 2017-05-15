//   Copyright 2016 Wercker Holding BV
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

// +build !windows

package cmd

import (
	"github.com/codegangsta/cli"
)

// These flags affect our local execution environment
var DevFlags = []cli.Flag{
	cli.StringFlag{Name: "environment", Value: "ENVIRONMENT", Usage: "Specify additional environment variables in a file."},
	cli.BoolFlag{Name: "verbose", Usage: "Print more information."},
	cli.BoolFlag{Name: "no-colors", Usage: "Wercker output will not use colors (does not apply to step output)."},
	cli.BoolFlag{Name: "debug", Usage: "Print additional debug information."},
	cli.BoolFlag{Name: "journal", Usage: "Send logs to systemd-journald. Suppresses stdout logging."},
}

