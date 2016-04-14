# RISCY-PEMU

## UNDER DEVELOPMENT

There are no guarantees on the state of the project, any information within this README.md is likely incorrect or subject to change. Development may (and likely will) occur on the master branch, there is no guarantee that previously pushed hashes won't change.


## The RISCV Processor EMUlator

riscy-pemu is a processor emulator written in go. Since it's in its early stages it emulates a single cycle processor with no regard for memory timings. It assumes that memory can be written to and read from in the same cycle.

See [RISCY-AS](https://github.com/spwilson2/RISCY-AS) for an assembler!


## Supported RISC-V Instructions
Only a subset of the RV32I ISA will be implemented. The following table lists supported instructions.

| Format |Args| Instruction |
| :------------- | ---- | ------------------------- |
|LUI  |%rd,$imm| Load Upper Immediate |
|JAL  |%rd,$imm |Jump And Link |
|BEQ  |%rs1,%rs2,$imm| Branch =|
|BNE  |%rs1,%rs2,$imm| Branch != |
|BLT  |%rs1,%rs2,$imm| Branch <|
|BGE  |%rs1,%rs2,$imm| Branch >=|
|BLTU |%rs1,%rs2,$imm|Branch <= Unsigned|
|BGEU |%rs1,%rs2,$imm| Breanch >= Unsigned|
|LW   |%rd,%rs1,$imm| Load Word |
|SW   |%rs1,%rs2,$imm|Store Word|
|ADDI |%rd,%rs1,$imm| ADD Immediate|
|XORI |%rd,%rs1,$imm|XOR Immediate |
|ORI  |%rd,%rs1,$imm| OR Immediate|
|ANDI |%rd,%rs1,$imm| AND Imemdiate|
|ADD  |%rd,%rs1,%rs2| ADD|
|SUB  |%rd,%rs1,%rs2| SUB |
|XOR  |%rd,%rs1,%rs2|XOR|
|OR   |%rd,%rs1,%rs2|OR|
|AND  |%rd,%rs1,%rs2|AND|
|SBREAK||SBREAK|
|NOP||NOP|
