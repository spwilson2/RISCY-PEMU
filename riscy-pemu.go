package main

import (
	"flag"
)

func main() {
	cli()
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
	dumpfilename := flag.String("m", "", "File to dump memory to")
	inputfilename := flag.String("i", "", "The input file path")
}
