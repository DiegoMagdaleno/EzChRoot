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
package main

import (
	"fmt"

	"github.com/diegomagdaleno/EzChRoot/lib"
)

func main() {
	var coreApplications = []string{"/bin/sh", "/bin/bash", "/bin/cat", "/bin/ls", "/bin/mkdir", "/bin/mv", "/bin/rm", "/bin/rmdir", "/bin/sh", "/bin/sleep", "/sbin/ping", "/usr/bin/curl", "/usr/bin/dig", "/usr/bin/env", "/usr/bin/grep", "/usr/bin/host", "/usr/bin/id", "/usr/bin/less", "/usr/bin/ssh", "/usr/bin/ssh-add", "/usr/bin/uname", "/usr/bin/vi", "/usr/lib/dyld", "/usr/sbin/netstat", "/usr/bin/clear"}
	test := lib.GetLinkedLibsLinux(coreApplications)
	fmt.Println(test)
}
