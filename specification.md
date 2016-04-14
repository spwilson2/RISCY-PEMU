# Specification (API) v1 

## Memory 

Memory in the RISCY-PEMU should be 32-bit and byte 

Memory internally will read from a file filling itself in on startup. It will then extend the following API to the processor:

```go
/*
    Fetch fetches a word at address from processor's simulated memory.
*/
func Fetch(uint32) uint32

/*
    Writes the word to the given address in processor's simulated memory.
*/
func Write(src uint32, dest uint32) 
```

Additionally, Memory is expected to offer tools for collecting stats and debugging:
```go
func Writes () uint64
func Fetches () uint64
```

Future thinking: In this stage of the specification, memory can both be written to and read from in the same cycle. However, in the future, it is likely this won't be the case. In-fact it's likely that memory will not directly communicate with the processor and instead will interact with a cache hierarchy.
    

## Register File

The register file will follow a similar API to that of memory:

```go
/*
    Fetch fetches and returns a word from the given register number.
*/
func Fetch(uint32) uint8

/*
    Write writes a word to the given register number.
*/
func Write(src uint32, dest uint32) 
```

Additionally, the register file should export the following for debugging and stat collection: 

```go
/*
    Dump the register file as an array of Reg
*/
Dump () []reg.Reg

/*
    Pc is the current program counter.
*/
Pc chan uint32
```

## Processor

Implement a single cycle processor that fetches an instruction from program memory, using the `mem.Fetch()`, and writes to memory using the `mem.Write()`. 

The processor itself will need to decode and execute each instruction from memory . The program counter itself will be used to fetch from memory at the rising edge of the clock. 

Data will be stored using a set of 31 registers, the zero register, and the PC, both of which will be in the register file.

The processor itself will need to implement the following interface:

```go
/*
    Processor emualtes a single cycle processor of the RISCV instruction set.

    clk is a channel representing the clock signal to the processor, rising edge of the clock
    signals the next instruction should be read in.

    hlt is a channel signalling the processor has halted.
*/
func Processor(clk chan bool) hlt chan bool
```

Separately, the Processor must also export the following go channels for debugging:

```go
Instr chan  // The current instruction, decoded 
```

Future thinking: The processor could be used to simulate different ISA's in the distant future, it would be a good idea to implement instructions in a manner that is easily extensible (think of using lambda functions mapped on top of their decoding).

# CLI

`riscy-pemu -i <binary> -q [-m file] [-h] [-r] [-c|-q] [--detailed]`

Will output various statistics about the simulation:

* Mem writes
* Mem reads
* Total Instructions executed

* [Count of each instruction executed]

* [Register dump]


`riscy-pemu -i <binary> -q [-m file] [-h] [-r] [-c|-q] [--detailed]`

* -i binary: The binary _input_ file
* -m file: _Memory_ dump to the file.
* -h : Pipe the dump into a _hex_ examiner to make it more readable before writing to a file. 
* -r : dump the _register_ file to the output after stats.
* -c : Print executed instruction counts.
* -q : Don't print stats.
* --detailed : Stats are very detailed.
