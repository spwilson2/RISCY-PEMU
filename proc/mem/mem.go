package mem

import (
	//"fmt"
	"bytes"
	"io/ioutil"
	"encoding/binary"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func read_int32(data []byte) (ret uint32) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.LittleEndian, &ret)
	return
}

/* Fill in the given array with the input file's data.*/
func Fill(inputfile string, arr *[]uint32) {

	data, err := ioutil.ReadFile(inputfile)
	check(err)

	for val := 0 ; val < len(data) / 4; val += 4 {
		*arr = append(*arr, read_int32(data[val*4:val*4+4]))
	}
}
