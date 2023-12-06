package main

import "github.com/agustin-del-pino/mos-6502-modular-emulator/pkg/mos"

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
	emu := mos.NewMOS6502(mos.Config{
		ProgramOffset:  0xFFFC,
		ResetMechanism: reset,
	})

	emu.AddInstruction(0xA9, func(cpu mos.CPU, m mos.Memory, r mos.Registers, c mos.Clock) error {
		a, err := cpu.Fetch()
		if err != nil {
			return err
		}

		r.Regs().A = a

		r.Regs().SetPSZero(a == 0)
		r.Regs().SetPSNegative(a&0b10000000 != 0)

		return nil
	})

	if err := emu.LoadProgram([]byte{0xA9, 0x10}); err != nil {
		panic(err)		
	}
	
	if err := emu.Run(); err != nil {
		panic(err)
	}

	println("program executed")
}
