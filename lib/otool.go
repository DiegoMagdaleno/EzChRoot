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
	"io/ioutil"
	"os/exec"
	"strings"
)

func stringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

/*GetLinkedLibs exports unique list of libs required for the programs
to run under the chroot. Logic can be improved so
TODO: Improve piping method*/
func GetLinkedLibs(libFile string) ([]string, error) {
	cleaned := []string{}
	content, err := ioutil.ReadFile(libFile)
	var (
		firstCalc []string
		splitOut  []string
	)
	if err != nil {
		panic("What?")
	}
	lines := strings.Split(string(content), "\n")
	for i := range lines {
		cmd := "otool -L " + lines[i] + "| sed 1d | awk '{print $1}' | xargs echo -n"
		out, err := exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			panic("Huoh")
		}
		splitOut = strings.Split((string(out)), " ")
		for k := range splitOut {
			firstCalc = append(firstCalc, splitOut[k])
		}
	}
	for _, value := range firstCalc {

		if !stringInSlice(value, cleaned) {
			cleaned = append(cleaned, value)
		}
	}
	return cleaned, err
}
