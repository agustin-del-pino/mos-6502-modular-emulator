package mos

const (
	maxMemo = 1024 * 64
)

type emulator struct {
	memo    [maxMemo]byte
	ins     map[byte]Instruction
	regs    *CPURegisters
	cycles  int
	pOffset uint
	pLen    uint
	reset   Instruction
}

func (e *emulator) Value() int {
	return e.cycles
}

func (e *emulator) Take(i int) {
	e.cycles += i
}

func (e *emulator) Read(a uint) (byte, error) {
	if a >= maxMemo {
		return 0, NewEmulationError("MemoryError", "memory exceed at reading: %d", a)
	}
	e.Take(1)
	return e.memo[a], nil
}

func (e *emulator) Write(a uint, v byte) error {
	if a >= maxMemo {
		return NewEmulationError("MemoryError", "memory exceed at writing: %d", a)
	}
	e.memo[a] = v
	e.Take(1)
	return nil
}

func (e *emulator) Load(a uint, v []byte) error {
	if len(v) >= maxMemo {
		return NewEmulationError("MemoryError", "memory exceed at loading: %d", a)
	}
	for i, b := range v {
		e.memo[a+uint(i)] = b
	}
	return nil
}

func (e *emulator) Dump() []byte {
	return e.memo[:]
}

func (e *emulator) Exec() error {
	b, err := e.Fetch()

	if err != nil {
		return err
	}

	ins, ok := e.ins[b]

	if !ok {
		return NewEmulationError("OpcodeError", "unhandled opcode: %x", b)
	}

	defer func() {
		e.cycles = 0
	}()

	return ins(CPU(e), Memory(e), Registers(e), Clock(e))
}

func (e *emulator) Fetch() (byte, error) {
	defer func() {
		e.regs.PC++
	}()
	return e.Read(e.regs.PC)
}

func (e *emulator) Reset() {
	e.reset(CPU(e), Memory(e), Registers(e), Clock(e))
	e.cycles = 0
}

func (e *emulator) ProgramOffset() uint {
	return e.pOffset
}

func (e *emulator) Regs() *CPURegisters {
	return e.regs
}

func (e *emulator) AddInstruction(op byte, i Instruction) {
	e.ins[op] = i
}

func (e *emulator) LoadProgram(p []byte) error {
	e.pLen = uint(len(p))
	return e.Load(e.pOffset, p)
}

func (e *emulator) Run() error {
	if e.pLen == 0 {
		return NewEmulationError("ProgramError", "there is no program loaded to run")
	}

	for (e.regs.PC - e.pOffset) < uint(e.pLen) {
		if err := e.Exec(); err != nil {
			return err
		}
	}

	return nil
}

// Config contains the emulator configuration.
type Config struct {
	// ProgramOffset is the memory address where the PC starts.
	ProgramOffset uint
	// ResetMechanism is the reset mechanism to be used by the CPU.
	ResetMechanism Instruction
}

// NewMOS6502 returns the default implementation of MOS6502Emulator.
func NewMOS6502(c Config) MOS6502Emulator {
	e := &emulator{
		ins:     make(map[byte]Instruction),
		regs:    &CPURegisters{},
		pOffset: c.ProgramOffset,
		reset:   c.ResetMechanism,
	}
	e.Reset()
	return e
}
