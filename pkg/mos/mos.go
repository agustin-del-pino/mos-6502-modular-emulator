package mos

// Registers provides the CPU registers management.
type Registers interface {
	// Regs returns the CPU registers.
	Regs() *CPURegisters
}

// Instruction represent the logic of a certain CPU instruction.
type Instruction func(CPU, Memory, Registers, Clock) error

// Memory provides the memory management of the emulator.
type Memory interface {
	// Read returns the value of an address memory.
	Read(a uint) (byte, error)
	// Write writes a value in an address memory.
	Write(a uint, v byte) error
	// Load sequentially loads a values starting at an address memory.
	Load(a uint, v []byte) error
	// Dump returns a copy of the whole memory.
	Dump() []byte
}

// CPU provides the emulated CPU for execute programs.
type CPU interface {
	// Exec executes the current loaded program.
	Exec() error
	// Fetch returns the current instruction at PC.
	Fetch() (byte, error)
	// Reset resets the CPU registers.
	Reset()
	// ProgramOffset returns the offset of memory where the PC starts.
	ProgramOffset() uint
}

// Clock provides the CPU clock cycles that takes an instruction.
type Clock interface {
	// Value returns the current cycles.
	Value() int
	// Take increases the current cycles.
	Take(int)
}

// MOS6502Emulator provides a modular emulator for MOS 6502.
type MOS6502Emulator interface {
	// AddInstruction adds a cpu instruction.
	AddInstruction(byte, Instruction)
	// LoadProgram loads a program into memory.
	LoadProgram([]byte) error
	// Run executes a loaded program.
	Run() error
}
