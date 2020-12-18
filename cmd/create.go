/*
Copyright © 2020 Diego Magdaleno <diegomagdaleno@protonmail.com>
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice,
   this list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice,
   this list of conditions and the following disclaimer in the documentation
   and/or other materials provided with the distribution.

3. Neither the name of the copyright holder nor the names of its contributors
   may be used to endorse or promote products derived from this software
   without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
POSSIBILITY OF SUCH DAMAGE.
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"runtime"
	"time"

	"github.com/briandowns/spinner"
	"github.com/diegomagdaleno/EzChRoot/lib"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create [name] [path]",
	Short: "create a new chroot",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var coreApplications []string
		var coreLibs []string
		var (
			validateName, _ = regexp.Compile("^[a-zA-Z]+$")
			name            = args[0]
			path            = args[1]
			s               = spinner.New(spinner.CharSets[11], 100*time.Millisecond)
		)
		switch os := runtime.GOOS; os {
		case "darwin":
			coreApplications = []string{"/bin/sh", "/bin/bash", "/bin/cat", "/bin/ls", "/bin/mkdir", "/bin/mv", "/bin/rm", "/bin/rmdir", "/bin/sleep", "/sbin/ping", "/usr/bin/curl", "/usr/bin/dig", "/usr/bin/env", "/usr/bin/grep", "/usr/bin/host", "/usr/bin/id", "/usr/bin/less", "/usr/bin/ssh", "/usr/bin/ssh-add", "/usr/bin/uname", "/usr/bin/vi", "/usr/lib/dyld", "/usr/sbin/netstat", "/usr/bin/clear"}
		case "linux":
			coreApplications = []string{"/bin/bash", "/bin/ls", "/bin/cp", "/bin/rm", "/bin/cat", "/bin/mkdir", "/bin/ln", "/bin/grep", "/bin/cut", "/bin/sed", "/usr/bin/ssh", "/usr/bin/head", "/usr/bin/tail", "/usr/bin/which", "/usr/bin/id", "/usr/bin/find", "/usr/bin/xargs"}
		}
		if !validateName.Match([]byte(args[0])) {
			log.Fatal(args[0] + " is not a valid name!")
		}
		if _, err := os.Stat(args[1]); os.IsNotExist(err) {
			fmt.Println("Selected path doesn't exist, will attempt to create it...")
			makeDirErr := os.MkdirAll(args[1], 0777)
			if makeDirErr != nil {
				log.Fatal("The directory " + args[1] + " couldn't be created")
			}
			fmt.Println("\342\234\223 Path was created sucessfully")
		}
		s.Color("green")

		s.Suffix = " Calculting dependencies for core system... "
		s.FinalMSG = "\342\234\223 Done calculating dependencies... \n"
		s.Start()
		coreLibs = lib.GetLinkedLibs(coreApplications)

		s.Stop()

		s.Suffix = " Copying core system into selected directory... "
		s.FinalMSG = "\342\234\223 Done copying core system into selected directory... \n"

		s.Start()
		lib.CopyToDir(coreApplications, coreLibs, path+"/"+name)
		s.Stop()

		s.Suffix = " Copying additional files into selected directory... "
		s.FinalMSG = "\342\234\223 Done with copying additional files to selected directory... \n"

		s.Start()
		lib.CopyAdditionalSettings(path + "/" + name)
		s.Stop()

		lib.WriteNewPath(name, path, coreApplications, coreLibs)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

}
