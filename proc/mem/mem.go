package mem

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
)

type Memory struct {
	memory []uint32
}

func (memory Memory) String() (out string) {

	for _, val := range memory.memory {
		out += fmt.Sprintf("0x%x\n", val)
	}

	return
}

/* Fill in the given Memory with the input file's data.*/
func (memory *Memory) Fill(inputfile string) {

	data, err := ioutil.ReadFile(inputfile)
	check(err)

	for val := 0; val < len(data)/4; val += 4 {
		memory.memory = append(memory.memory, read_int32(data[val*4:val*4+4]))
	}
}

/* Fetch fetches a word at address from processor's simulated memory.
 */
func Fetch(uint32) uint32 {
	return 0
}

/* Writes the word to the given address in processor's simulated memory.
 */
func Write(src uint32, dest uint32) {
}

//TODO: Implement a string function on the mem struct

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
