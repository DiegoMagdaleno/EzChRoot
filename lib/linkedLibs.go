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

package lib

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/codeskyblue/go-sh"
)

var (
	librariesCalc  []string
	librariesSplit []string
)

var libraries = []string{}

func stringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

/* Function that is specifically for Darwin, gets the linked
libs using otool */
func GetLinkedLibsDarwin(libraryFile []string) []string {
	for i := range libraryFile {
		test, err := sh.Command("otool", "-L", libraryFile[i]).Command("sed", "1d").Command("awk", "{print $1}").Command("xargs", "echo", "-n").Output()
		if err != nil {
			log.Panic(err)
		}
		librariesSplit = strings.Split(string((test)), " ")
		for k := range librariesSplit {
			librariesCalc = append(librariesCalc, librariesSplit[k])
		}

		for _, value := range librariesCalc {

			if !stringInSlice(value, libraries) && (value != "\n") && (len(value) != 0) {
				libraries = append(libraries, value)
			}
		}

		if reflect.DeepEqual(libraryFile, libraries) == true {
			return libraries
		}
	}
	return GetLinkedLibsDarwin(libraries)
}

/* Function that gets the linked libraries on Linux uses ldd and the
output is managed diferently) */
func GetLinkedLibsLinux(libraryFile []string) []string {
	for i := range libraryFile {
		// This was pain to write, it took me hours to figure out how to split the output
		// because I didnt know awk an sed.
		library, err := sh.Command("ldd", libraryFile[i]).Command("sed", "1d").Command("awk", "{print $3}").Output()
		if err != nil {
			fmt.Println(err)
		}
		librariesSplit = strings.Split(string(library), "\n")
		for k := range librariesSplit {
			librariesCalc = append(librariesCalc, librariesSplit[k])
		}
		for _, value := range librariesCalc {
			if !stringInSlice(value, libraries) && (value != "\n") && (len(value) != 0) {
				libraries = append(libraries, value)
			}
		}

		if reflect.DeepEqual(libraryFile, libraries) == true {
			return libraries
		}
	}
	return GetLinkedLibsLinux(libraries)
}
