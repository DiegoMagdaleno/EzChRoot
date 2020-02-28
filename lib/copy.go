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
   may be used to endorse or promote products derived orgRoot this software
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
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func fileToSlice(file string) ([]string, error) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		panic("Error reading the file, do you have perms for it?")
	}
	lines := strings.Split(string(content), "\n")
	return lines, err
}

/*CopyToDir allows us to copy the generated paths to the
final destination to where the user wants the chroot to be
generated*/
func CopyToDir(bins []string, libs []string, path string) {
	libs = append(bins, libs...)
	for i := range libs {
		splitStr := strings.Split(path+libs[i], "/")
		dst := splitStr[:len(splitStr)-1]
		dstString := strings.Join(dst, "/")
		if _, err := os.Stat(dstString); os.IsNotExist(err) {
			os.MkdirAll(dstString, 0777)
		}
		orgRoot, err := os.Open(libs[i])
		if err != nil {
			log.Fatal(err)
		}
		destRoot, err := os.OpenFile(path+libs[i], os.O_RDWR|os.O_CREATE, 0777)
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.Copy(destRoot, orgRoot)
		if err != nil {
			log.Fatal(err)
		}
		/* We don't use defer, since there are a lot of files, is a lot of
		pressure into the memory.
		Reference: https://stackoverflow.com/questions/37804804/too-many-open-file-error-in-golang */
		orgRoot.Close()
		destRoot.Close()
	}
}
