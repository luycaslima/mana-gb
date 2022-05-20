package cpu

import (
	"mana-gb/types"
)

type Register struct {
	REG types.Word
	LO  byte //4 bytes inferiores Exemplo: AF , -> A
	HI  byte // 4bytes superiores , logo  ^  F
}

//TODO FAZER PRIMEIRO A CPU E SUAS FUNÇÕES
//TODO COMO DECLARAR AS FLAGS?

//Quebrar em 8 registradores de uint8?
type Registers struct {
	AF Register
	BC Register
	DE Register
	HL Register
	//Flags uint8
}

type CPU struct { //frame atual da cpu
	R  Registers
	SP types.Word
	PC types.Word
}

func (cpu *CPU) ExecuteOPCode(opcode uint16) {

}
