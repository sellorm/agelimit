package main

// TODO
// * Can we add windows support? (Or should windows users use WSL and Linux?)

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"time"
)

var version string

func myUsage() {
	fmt.Printf("Usage: %s [OPTIONS] maxage file\n", os.Args[0])
	fmt.Println("")
	fmt.Println("  maxage      The maximum age for your file and a unit specifier. For example, 100s")
	fmt.Println("              for 100 seconds, 20m for 20 minutes, 3h for 3 hours or 2d for 2 days")
	fmt.Println("  file        The path to the file to be tested")
	flag.PrintDefaults()
}

func ageToSeconds(maxage string) int64 {
	// Check that the input is in the numberunit format, eg 100s, 10h, 1d etc.
	match, _ := regexp.MatchString("^[0-9]*[smhd]$", maxage)
	if !match {
		fmt.Println("Error: The time unit specified must be in the format \"100d\"")
		os.Exit(99)
	}
	// Split our time specification into the numerical part and the s, m, h, d part
	maxage_str := string(maxage[0 : len(maxage)-1])
	units := string(maxage[len(maxage)-1:])

	// convert the input seconds to an int
	maxage_num, err := strconv.Atoi(maxage_str)
	if err != nil {
		// ... handle error
		panic(err)
	}

	// Convert the input age and units into seconds
	var age_seconds int64
	if units == "s" {
		age_seconds = int64(maxage_num)
	} else if units == "m" {
		age_seconds = int64(maxage_num) * 60
	} else if units == "h" {
		age_seconds = int64(maxage_num) * 3600
	} else if units == "d" {
		age_seconds = int64(maxage_num) * 86400
	} else {
		fmt.Println("Error: The time unit specified must be one of s, m, h, d")
		os.Exit(99)
	}
	return (age_seconds)
}

func main() {

	flag.Usage = myUsage

	// command line flags
	verbosePtr := flag.Bool("v", false, "Enable verbose output")
	silentPtr := flag.Bool("s", false, "Silence all output")
	versionPtr := flag.Bool("version", false,
		"Display the version number and quit")

	flag.Parse()

	// Print version and quit
	if *versionPtr {
		fmt.Println(version)
		os.Exit(0)
	}

	// Check both positional arguments have been supplied
	var maxage string
	var filename string

	if len(flag.Args()) == 2 {
		maxage = flag.Args()[0]
		filename = flag.Args()[1]
	} else {
		fmt.Println("Error: Please supply an age and a filepath, eg 1000s ./file.txt")
		os.Exit(99)
	}

	// Check the file actully exists
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		fmt.Println("Error: filename, ",
			filename,
			"not found. Please check the path and try again")
		os.Exit(99)
	}

	age_seconds := ageToSeconds(maxage)

	// get file information from the OS
	//fileinfo, _ := os.Stat(filename)
	//stat := fileinfo.Sys().(*syscall.Stat_t)

	// ------ old stuff -------
	// atime is the file access time.
	// fmt.Println(time.Unix(stat.Atimespec.Sec, stat.Atimespec.Nsec))

	// mtime is the file modify time.
	// fmt.Println(time.Unix(stat.Mtimespec.Sec, stat.Mtimespec.Nsec))
	// ------------------------

	file_ctime := GetCtime(filename)

	// Print our verbose outputs
	if *verbosePtr {
		fmt.Println("positional args:      ", flag.Args())
		fmt.Println("flags:")
		fmt.Println("  verbose:            ", *verbosePtr)
		fmt.Println("  version:            ", *versionPtr)
		fmt.Println("  silent:             ", *silentPtr)
		// fmt.Println("units:                ", units) // Removed when abstracting the ageToSeconds function out
		// fmt.Println("maxage input number:  ", maxage_num) // Removed when abstracting the ageToSeconds function out
		fmt.Println("Filename:             ", filename)
		fmt.Println("Max age (seconds):    ", age_seconds)
		// ctime is the inode or file change time.
		//fmt.Println("ctime:                ",
		//	time.Unix(stat.Ctimespec.Sec, stat.Ctimespec.Nsec))
		// fmt.Println("ctime:                ",
		//	time.Unix(file_ctime))
		// print the OS and architecture
		fmt.Println("OS:                   ", runtime.GOOS)
		fmt.Println("CPU Architecture:     ", runtime.GOARCH)
		// raw output
		fmt.Println("ctime timestamp:      ", file_ctime)
		// Current time
		fmt.Println("Current timstamp:     ", time.Now().Unix())
	}

	// Check if the input file is older than required
	if time.Now().Unix()-int64(age_seconds) > file_ctime {
		if !*silentPtr {
			fmt.Println("agelimit: Error, limit breached!",
				filename,
				"is over",
				maxage,
				"old.")
		}
		os.Exit(1)
	}
	// Print success message
	if !*silentPtr {
		fmt.Println("agelimit: Success!",
			filename,
			"is less that the",
			maxage,
			"limit.")
	}
	os.Exit(0)
}
