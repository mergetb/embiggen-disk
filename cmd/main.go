/*
Copyright 2018 Google Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// The embiggen-disk command live resizes a filesystem and LVM objects
// and partition tables as needed. It's useful within a VM guest to make
// its filesystem bigger when the hypervisor live resizes the underlying
// block device.
package main

// TODO: test/fix on disks with non-512 byte sectors ( /sys/block/sda/queue/hw_sector_size)

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/google/embiggen-disk"
)

var (
	dry     = flag.Bool("dry-run", false, "don't make changes")
	verbose = flag.Bool("verbose", false, "verbose output")
)

func init() {
	flag.Usage = usage
	embiggen.Dry = *dry
	embiggen.Verbose = *verbose
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of embiggen-disk:\n\n")
	fmt.Fprintf(os.Stderr, "# embiggen-disk [flags] <mount-point-to-enlarge>\n\n")
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		usage()
	}
	if runtime.GOOS != "linux" {
		fatalf("embiggen-disk only runs on Linux.")
	}

	mnt := flag.Arg(0)
	err := embiggen.Embiggen(mnt)
	if err != nil {
		fatalf(err.Error())
	}

}

func fatalf(format string, args ...interface{}) {
	log.SetFlags(0)
	log.Fatalf(format, args...)
}
