# MOS 6502 Modular Emulator
Another MOS 6502 Emulator but this is a modular one.

# Overview

The goal of this emulator is to be modular. That is done by not having the instructions implemented at the emulator itself. Instead of that, the emulator provides a mechanism for "add" an instruction implementation. 

This way to implement it's like "event" programming. Each instruction is an event to be listen to. 

The emulator also provides interfaces for: `CPU`, `Memory`, `Registers` and `Clock`.

## The Cycles
The MOS 6502 take one cycle each time operates with the memory. So, the emulator already integrates a "clock" that accumulates the cycles taken by memory operations.

## Registers
The MOS 6502 has three registers: A, X and Y, but also has PC, SP and the flags of PS. Those can be operated like bytes/uint, in the case of the flags, the `Registers` interface provides many methods for operate each flag.

## The Memory
The memory has a max size of 64KB and the only methods that takes cycles are: Read and Write.

## CPU
The CPU executes the current instruction at PC. In case to execute more than one it must be called as many time as number of instruction the program has.

## Emulator
The emulator itself provides method for: add a new instruction, load a program and run the emulator.

When the emulator runs, it will execute the program by calling the CPU until reach the end of the program.

# Example

````go
// cpu reset mechanism (copied from C64)
func reset(cpu mos.CPU, mem mos.Memory, reg mos.Registers, clk mos.Clock) error {
	reg.Regs().PC = cpu.ProgramOffset()
	reg.Regs().SP = 0xFF
	reg.Regs().PS = 0
	reg.Regs().A = 0
	reg.Regs().X = 0
	reg.Regs().Y = 0
	return nil
}

func main() {
  // instance a new emulator
	emu := mos.NewMOS6502(mos.Config{
		ProgramOffset:  0xFFFC,
		ResetMechanism: reset,
	})

  // Add the LDA (A9) instruction
	emu.AddInstruction(0xA9, func(cpu mos.CPU, m mos.Memory, r mos.Registers, c mos.Clock) error {
    // Get the LDA value
		a, err := cpu.Fetch()
		if err != nil {
			return err
		}

    // Load the value to the A register.
		r.Regs().A = a

    // Set the SP flags

    // Zero flag, if and only if, the A register value is 0.
		r.Regs().SetPSZero(a == 0) 
    // Negative flag, if and only if, the 7 bit of the A register is set.
		r.Regs().SetPSNegative(a&0b10000000 != 0)

		return nil
	})

  // load a program: LDA #$10
	if err := emu.LoadProgram([]byte{0xA9, 0x10}); err != nil {
		panic(err)		
	}
	
  // run the emulator.
	if err := emu.Run(); err != nil {
		panic(err)
	}

  // use this for breakpoint.
	println("program executed")
}

````