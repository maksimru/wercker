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

package cmd;

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/wercker/wercker/util"
)

func GetApp() *cli.App {
	// logger.SetLevel(logger.DebugLevel)
	// util.RootLogger().SetLevel("debug")
	// util.RootLogger().Formatter = &logger.JSONFormatter{}

	app := cli.NewApp()
	setupUsageFormatter(app)
	app.Author = "Team wercker"
	app.Name = "wercker"
	app.Usage = "build and deploy from the command line"
	app.Email = "pleasemailus@wercker.com"
	app.Version = util.FullVersion()
	app.Flags = FlagsFor(GlobalFlagSet)
	app.Commands = []cli.Command{
		buildCommand,
		devCommand,
		checkConfigCommand,
		deployCommand,
		detectCommand,
		// inspectCommand,
		loginCommand,
		logoutCommand,
		pullCommand,
		versionCommand,
		documentCommand(app),
	}
	app.Before = func(ctx *cli.Context) error {
		if ctx.GlobalBool("debug") {
			util.RootLogger().Formatter = &util.VerboseFormatter{}
			util.RootLogger().SetLevel("debug")
		} else {
			util.RootLogger().Formatter = &util.TerseFormatter{}
			util.RootLogger().SetLevel("info")
		}
		// Register the global signal handler
		util.GlobalSigint().Register(os.Interrupt)
		return nil
	}
	return app
}

