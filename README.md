# Table of contents

- [EzChroot](#ezchroot)
- [Why it was deprecated?](#why-was-it-deprecated)
- [How could it be improved?](#how-could-it-be-improved)
- [Features](#features)
- [Installation](#installation)
- [License](#license)

## EzChroot

EzChroot was a proof of concept of making a way to easily create Chroot "jails" in modern versions of macOS, before writing this I tried scattering across the internet to find a solution that I liked, however, I ended up not liking any implementation, as they were bash scripts and so on, this is how I decided to write my own simple implementation.

Some days later I added support for Linux, to also be able to create some Chroot there.

## Why was it deprecated?

In retrospect, it isn't good or special at all, but it was deprecated due to changes in macOS, during sometime in 2020 during the release of Big Sur, Apple changed the way System Dynamic Libraries worked, as they weren't in the filesystem anymore, but rather in a custom cache designed by Apple, this represented a big problem for EzChroot, as it depended on getting the linked libraries of each binary and then copying them over to the folder where the Chroot env was going to be located.

While I tried some other techniques to fix it, such as dumping the cache by forking [dyld](https://github.com/DiegoMagdaleno/dyld) it was ultimately impossible to get the Chroot to work, because of some changes on how the libraries are loaded in macOS.

While Linux support was still an option to maintain this project, there are much better utilities for Linux that do satisfy my necessities, so there is no reason to develop Yet Another Tool.

## How could it be improved?

EzChroot was my first project on Golang after years of only writing Python, so it does have a lot of room for improvement if I were to come back and time and do it from scratch:

- Don't call Otool: At the time of writing this program, my knowledge on macOS was low to say at best, so I just researched how to dump the shared libraries, this ultimately ended up in me using the `os.exec()` to run a command that I found on StackOverflow, if I were to write this today, I would implement a Mach-O parser, and get the proper headers were the dynamic libraries are located.

- Use a database: At the time of writing this, my knowledge of databases was almost `nil` as you would say in Go, this had the implication that EzChroot tried to store stuff into a JSON file, which was extremely wrong and hard to use, the proper way if I were to write this today would be to use SQLite or another lightweight local database.

- Don't hardcode "essential" binaries: When I made EzChroot, I made it just thinking about my necessities, which ultimately resulted in my hardcoding paths that other people didn't have, if I were to write this today, I would have made a dotfile that is inside the directory where you want to create the chroot, containing the names of the binaries you might want to copy.


<p align="Center">
 <img width="460" height="300" src="https://media.giphy.com/media/LRlCIiIeUy3xkGsOSJ/giphy.gif">
</p>

## Installation

Here is some information on how to install EzChRoot in your Mac!

- **Option 1 - Precompiled bin**:
  - Go to GitHub releases in this repository
  - Download the latest release
  - Copy the bin inside `/usr/local` or whatever directory you prefer
  - Have fun!
  
- **Option 2 - From source**:
  - Download a tar or zip archive of this repository
  - Make sure you have [Go](https://golang.org) installed
  - Uncompress the tar or zip archive
  - `cd`into the directory where the source code is
  - Type `sudo make install` and wait for the installation!
  
  ## Features 
  
  Ezchroot is new and beta software, so don't expect a lot of features, as of now however, it can:
  
  - Create a chroot
  - Delete a chroot
  
## License

BSD-3-Clause
  
  
