// This code uses work derived from u-root:
// Copyright 2015-2017 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.uroot file.

package lib

import (
	"log"
	"os"
	"path/filepath"
	"syscall"
)

type mount struct {
	src, dst, mtype, opts string
	flags                 uintptr
	dir                   bool
}

func modedev(stat os.FileInfo) (uint32, int) {
	dev := int(stat.Sys().(*syscall.Stat_t).Dev)
	devShift := dev & 0xff
	dev >>= 8
	dev |= (devShift << 8)
	mode := uint32(stat.Sys().(*syscall.Stat_t).Mode)
	return mode, dev
}

func makeConsole(base string, console string) {

	err := os.Chmod(console, 0600)
	if err != nil {
		panic(err)
	}

	err = os.Chown(console, 0, 0)
	if err != nil {
		panic(err)
	}

	stat, err := os.Stat(console)
	if err != nil {
		panic(err)
	}

	target := filepath.Join(base, "/dev/console")
	mode, dev := modedev(stat)
	err = syscall.Mknod(target, mode, dev)
	if err != nil {
		panic(err)
	}

	err = syscall.Mount(console, target, "", syscall.MS_BIND, "")
	if err != nil {
		panic(err)
	}

}

func (m *mount) One(base string) {
	dst := filepath.Join(base, m.dst)
	if m.dir {
		if err := os.MkdirAll(dst, 0755); err != nil {
			log.Fatalf("One: mkdirall %v: %v", m.dst, err)
		}
	}
	if err := syscall.Mount(m.src, dst, m.mtype, m.flags, m.opts); err != nil {
		log.Fatalf("Mount :%s: on :%s: type :%s: flags %x: opts :%v: %v\n",
			m.src, m.dst, m.mtype, m.flags, m.opts, err)
	}
}

var (
	root   = &mount{"", "/", "", "", syscall.MS_SLAVE | syscall.MS_REC, true}
	mounts = []mount{
		{"proc", "/proc", "proc", "", syscall.MS_NOSUID | syscall.MS_NOEXEC | syscall.MS_NODEV, true},
		{"/proc/sys", "/proc/sys", "", "", syscall.MS_BIND, true},
		{"", "/proc/sys", "", "", syscall.MS_BIND | syscall.MS_RDONLY | syscall.MS_REMOUNT, true},
		{"sysfs", "/sys", "sysfs", "", syscall.MS_NOSUID | syscall.MS_NOEXEC | syscall.MS_NODEV | syscall.MS_RDONLY, true},
		{"tmpfs", "/dev", "tmpfs", "mode=755", syscall.MS_NOSUID | syscall.MS_STRICTATIME, true},
		{"devpts", "/dev/pts", "devpts", "newinstance,ptmxmode=0660,mode=0620", syscall.MS_NOSUID | syscall.MS_NOEXEC, true},
		{"tmpfs", "/dev/shm", "tmpfs", "mode=1777", syscall.MS_NOSUID | syscall.MS_STRICTATIME | syscall.MS_NODEV, true},
		{"tmpfs", "/run", "tmpfs", "mode=755", syscall.MS_NOSUID | syscall.MS_NODEV | syscall.MS_STRICTATIME, true},
	}
)

func MountAll(base string) {
	root.One("")
	for _, m := range mounts {
		m.One(base)
	}
}

func copyNodes(base string) {
	nodes := []string{"/dev/tty", "/dev/full", "/dev/zero", "/dev/null", "/dev/random", "/dev/urandom"}
	for _, node := range nodes {
		stat, err := os.Stat(node)
		if err != nil {
			panic(err)
		}
		target := filepath.Join(base, node)
		mode, dev := modedev(stat)
		err = syscall.Mknod(target, mode, dev)
		if err != nil {
			panic(err)
		}

	}
}
