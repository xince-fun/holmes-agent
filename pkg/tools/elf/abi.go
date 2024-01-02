package elf

import (
	"fmt"
	"strings"
)

type ArgLocation struct {
	Type   LocationType
	Offset uint64
}

type ArgumentLocator interface {
	// GetLocation of the argument
	GetLocation(typeClass TypeClass, typeSize uint64, alignmentSize uint64, primitivesCount int, isRetArg bool) (*ArgLocation, error)
}

func NewArgumentLocator(r *DwarfReader) ArgumentLocator {
	if r.language == ReaderLanguageGolang {
		// is register-based ABI
		// example: "Go cmd/comile go1.18; regabi"
		if strings.Contains(r.producer, "regabi") {
			return newGolangRegisterLocator()
		}
		return &GolangStackLocator{}
	}
	return &UnknownLocator{}
}

type GolangRegisterLocator struct {
	curStackOffset       uint64
	curIntArgOffset      uint64
	curFloatArgOffset    uint64
	curIntRetArgOffset   uint64
	curFloatRetArgOffset uint64

	intArgRegisters      []RegisterName
	floatArgRegisters    []RegisterName
	intRetArgRegisters   []RegisterName
	floatRetArgRegisters []RegisterName
}

func newGolangRegisterLocator() *GolangRegisterLocator {
	intRegisters := []RegisterName{
		RegisterNameRAX, RegisterNameRBX, RegisterNameRCX,
		RegisterNameRDI, RegisterNameRSI, RegisterNameR8,
		RegisterNameR9, RegisterNameR10, RegisterNameR11,
	}
	floatRegisters := []RegisterName{
		RegisterNameXMM3, RegisterNameXMM4, RegisterNameXMM5,
		RegisterNameXMM6, RegisterNameXMM7, RegisterNameXMM8,
		RegisterNameXMM9, RegisterNameXMM10, RegisterNameXMM11,
		RegisterNameXMM12, RegisterNameXMM13, RegisterNameXMM14,
	}
	return &GolangRegisterLocator{
		intArgRegisters:      intRegisters,
		floatArgRegisters:    floatRegisters,
		intRetArgRegisters:   intRegisters,
		floatRetArgRegisters: floatRegisters,
	}
}

func (g *GolangRegisterLocator) GetLocation(typeClass TypeClass, typeSize uint64, alignmentSize uint64, primitivesCount int, isRetArg bool) (*ArgLocation, error) {
	var registers []RegisterName
	var offset *uint64
	if typeClass == TypeClassInteger {
		if isRetArg {
			registers = g.intRetArgRegisters
			offset = &g.curIntRetArgOffset
		} else {
			registers = g.intArgRegisters
			offset = &g.curIntArgOffset
		}
	} else if typeClass == TypeClassFloat {
		if isRetArg {
			registers = g.floatRetArgRegisters
			offset = &g.curFloatRetArgOffset
		} else {
			registers = g.floatArgRegisters
			offset = &g.curFloatArgOffset
		}
	} else {
		return nil, fmt.Errorf("unsupported type class for getting location, type class: %d", typeClass)
	}

	result := &ArgLocation{}
	if primitivesCount <= len(registers) {
		if typeClass == TypeClassInteger {
			result.Type = ArgLocationTypeRegister
		} else {
			result.Type = ArgLocationTypeRegisterFP
		}
		result.Offset = *offset
		*offset += uint64(primitivesCount * 8)
	} else {
		g.curStackOffset = snapUpToMultiple(g.curStackOffset, alignmentSize)
		result.Type = ArgLocationTypeStack
		result.Offset = g.curStackOffset

		g.curStackOffset += typeSize
	}
	return result, nil
}

type GolangStackLocator struct {
	curStackOffset uint64
}

func (g *GolangStackLocator) GetLocation(_ TypeClass, typeSize, alignmentSize uint64, _ int, _ bool) (*ArgLocation, error) {
	g.curStackOffset = snapUpToMultiple(g.curStackOffset, alignmentSize)
	result := &ArgLocation{}
	result.Type = ArgLocationTypeStack
	result.Offset = g.curStackOffset

	g.curStackOffset += typeSize
	return result, nil
}

type UnknownLocator struct {
	language int
}

func (u *UnknownLocator) GetLocation(typeClass TypeClass, typeSize, alignmentSize uint64, primitivesCount int, isRetArg bool) (*ArgLocation, error) {
	return nil, fmt.Errorf("unknown locator for language: %d", u.language)
}

func snapUpToMultiple(curSize, alignmentSize uint64) uint64 {
	return ((curSize + (alignmentSize - 1)) / alignmentSize) * alignmentSize
}
