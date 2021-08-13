package lib

import (
	"os"
	"strconv"
	"syscall"

	"golang.org/x/sys/unix"
)

func ptsname(file *os.File) (string, error) {
	result, err := unix.IoctlGetInt(int(file.Fd()), unix.TIOCGPTN)
	if err != nil {
		return "", err
	}
	return "/dev/pts" + strconv.Itoa(result), nil
}

func ptsopen() (controlPTY, processTTY *os.File, ttyName string, err error) {
	p, err := os.OpenFile("/dev/ptmx", os.O_RDWR, 0)
	if err != nil {
		return
	}

	ttyName, err = ptsname(p)
	if err != nil {
		return
	}

	file, err := os.OpenFile(ttyName, os.O_RDWR|syscall.O_NOCTTY, 0)
	if err != nil {
		return
	}

	return p, file, ttyName, nil

}
