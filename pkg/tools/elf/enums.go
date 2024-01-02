package elf

type TypeClass int

const (
	TypeClassNone    TypeClass = 0
	TypeClassInteger TypeClass = 1
	TypeClassFloat   TypeClass = 2
	TypeClassMixed   TypeClass = 3
)

func (c TypeClass) Combine(other TypeClass) TypeClass {
	if c == TypeClassMixed || other == TypeClassMixed {
		return TypeClassMixed
	}
	if c == TypeClassNone {
		return other
	}
	if other == TypeClassNone {
		return c
	}

	if c != other {
		return TypeClassMixed
	}
	return c
}

type LocationType uint64

const (
	ArgLocationTypeUnknown    LocationType = 0
	ArgLocationTypeStack      LocationType = 1 // frame stack pointer
	ArgLocationTypeStackBP    LocationType = 2 // frame base pointer
	ArgLocationTypeRegister   LocationType = 3 // integer register
	ArgLocationTypeRegisterFP LocationType = 4 // float-point register
)

type RegisterName int

const (
	// for int type class
	RegisterNameRAX RegisterName = 0
	RegisterNameRBX RegisterName = 1
	RegisterNameRCX RegisterName = 2
	RegisterNameRDX RegisterName = 3
	RegisterNameRDI RegisterName = 4
	RegisterNameRSI RegisterName = 5
	RegisterNameR8  RegisterName = 6
	RegisterNameR9  RegisterName = 7
	RegisterNameR10 RegisterName = 8
	RegisterNameR11 RegisterName = 9

	// for float type class
	RegisterNameXMM0  RegisterName = 100
	RegisterNameXMM1  RegisterName = 101
	RegisterNameXMM2  RegisterName = 102
	RegisterNameXMM3  RegisterName = 103
	RegisterNameXMM4  RegisterName = 104
	RegisterNameXMM5  RegisterName = 105
	RegisterNameXMM6  RegisterName = 106
	RegisterNameXMM7  RegisterName = 107
	RegisterNameXMM8  RegisterName = 108
	RegisterNameXMM9  RegisterName = 109
	RegisterNameXMM10 RegisterName = 110
	RegisterNameXMM11 RegisterName = 111
	RegisterNameXMM12 RegisterName = 112
	RegisterNameXMM13 RegisterName = 113
	RegisterNameXMM14 RegisterName = 114
)
