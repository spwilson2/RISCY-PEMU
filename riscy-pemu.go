package main

import (
	"github.com/spwilson2/riscy-pemu/proc/mem"
	"flag"
	"fmt"
)

func main() {
	inputfilename := flag.String("i", "", "The input file path") 
	dumpfilename := flag.String("m", "default", "File to dump memory to")
	dumpashex := flag.Bool("h", false, "Pipe memory dump into a hex formater")
	dumpregs := flag.Bool("r", false, "Dump the register file with stats.")
	printcounts := flag.Bool("c", false, "Print count of each executed instruction")
	quiet := flag.Bool("q", true, "Don't print stats")
	detailed := flag.Bool("detailed", false, "Stats are very detailed")

	flag.Parse()

	_ = dumpfilename
	_ = inputfilename
	_ = dumpashex
	_ = dumpregs
	_ = printcounts
	_ = quiet
	_ = detailed

	if *inputfilename == "" {
		fmt.Println("Must supply an input file path.")
		fmt.Println("")
		flag.PrintDefaults()
		return
	}

	var memory []uint32

	mem.Fill(*inputfilename, &memory)
	for _, val := range memory {
		fmt.Printf("0x%x\n", val)
	}
}

/*
Arguments are:

`riscy-pemu -i <binary> -q [-m file] [-h] [-r] [-c|-q] [--detailed]`

* -i binary: The binary _input_ file
* -m file: _Memory_ dump to the file.
* -h : Pipe the dump into a _hex_ examiner to make it more readable before writing to a file.
* -r : dump the _register_ file to the output after stats.
* -c : Print executed instruction counts.
* -q : Don't print stats.
* --detailed : Stats are very detailed.
*/
func cli() {
}
