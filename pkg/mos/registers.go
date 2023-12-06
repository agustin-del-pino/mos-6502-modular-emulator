package mos

type CPURegisters struct {
	// A is accumulator register.
	A byte
	// X is X register.
	X byte
	// Y is Y register.
	Y byte

	// PC is the program counter.
	PC uint
	// SP is the stack pointer.
	SP byte
	// PS is the processor status.
	PS byte
}

// PSFlag returns an specified flag (zero-base index).
func (rgs *CPURegisters) PSFlag(f byte) byte {
	return (byte(rgs.PS) >> f) & 1
}

// IsPSFlagOn reports if an specified flag (zero-base index) is on (1).
func (rgs *CPURegisters) IsPSFlagOn(f byte) bool {
	return rgs.PSFlag(f) == 1
}

// IsPSNegative reports if the negative flag is on.
func (rgs *CPURegisters) IsPSNegative() bool {
	return rgs.IsPSFlagOn(7)
}

// IsPSOverflow reports if the overflow flag is on.
func (rgs *CPURegisters) IsPSOverflow() bool {
	return rgs.IsPSFlagOn(6)
}

// IsPSBreak reports if the break command flag is on.
func (rgs *CPURegisters) IsPSBreakCommand() bool {
	return rgs.IsPSFlagOn(4)
}

// IsPSDecimalMode reports if the decimal mode flag is on.
func (rgs *CPURegisters) IsPSDecimalMode() bool {
	return rgs.IsPSFlagOn(3)
}

// IsPSInterruptDisable reports if the interrupt disable flag is on.
func (rgs *CPURegisters) IsPSInterruptDisable() bool {
	return rgs.IsPSFlagOn(2)
}

// IsPSZero reports if the zero flag is on.
func (rgs *CPURegisters) IsPSZero() bool {
	return rgs.IsPSFlagOn(1)
}

// IsPSCarry reports if the carry flag is on.
func (rgs *CPURegisters) IsPSCarry() bool {
	return rgs.IsPSFlagOn(0)
}

func (rgs *CPURegisters) SetPSFlag(f byte, v bool) {
	if v {
		rgs.PS = rgs.PS | (1 << f)
	} else {
		rgs.PS = rgs.PS & (^(1 << f))
	}
}

// SetPSNegative reports if the negative flag is on.
func (rgs *CPURegisters) SetPSNegative(v bool) {
	rgs.SetPSFlag(7, v)
}

// SetPSOverflow reports if the overflow flag is on.
func (cpu *CPURegisters) SetPSOverflow(v bool) {
	cpu.SetPSFlag(6, v)
}

// SetPSBreak reports if the break command flag is on.
func (cpu *CPURegisters) SetPSBreakCommand(v bool) {
	cpu.SetPSFlag(4, v)
}

// SetPSDecimalMode reports if the decimal mode flag is on.
func (cpu *CPURegisters) SetPSDecimalMode(v bool) {
	cpu.SetPSFlag(3, v)
}

// SetPSInterruptDisable reports if the interrupt disable flag is on.
func (cpu *CPURegisters) SetPSInterruptDisable(v bool) {
	cpu.SetPSFlag(2, v)
}

// SetPSZero reports if the zero flag is on.
func (cpu *CPURegisters) SetPSZero(v bool) {
	cpu.SetPSFlag(1, v)
}

// SetPSCarry reports if the carry flag is on.
func (cpu *CPURegisters) SetPSCarry(v bool) {
	cpu.SetPSFlag(0, v)
}
