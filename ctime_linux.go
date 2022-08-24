//go:build linux || !darwin
// +build linux !darwin

package main

import (
	"os"
	"syscall"
)

func GetCtime(filename string) int64 {
	fileinfo, _ := os.Stat(filename)
	stat := fileinfo.Sys().(*syscall.Stat_t)
	return stat.Ctim.Sec
}
