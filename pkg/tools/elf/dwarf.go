package elf

import (
	"debug/dwarf"
	"encoding/binary"
	"fmt"
)

var ReaderLanguageGolang = 22

type DwarfReader struct {
	elfByteOrder binary.ByteOrder

	producer   string
	language   int
	functions  map[string]*FunctionInfo
	structures map[string]*StructureInfo
}

func (f *File) NewDwarfReader(attrNames ...string) (*DwarfReader, error) {
	data, err := f.realFile.DWARF()
	if err != nil {
		return nil, err
	}
	reader := &DwarfReader{}
	if err := reader.init(data, attrNames); err != nil {
		return nil, err
	}
	return reader, nil
}

func (r *DwarfReader) GetFunction(name string) *FunctionInfo {
	return r.functions[name]
}

func (r *DwarfReader) GetStructure(name string) *StructureInfo {
	return r.structures[name]
}

func (r *DwarfReader) GetStructureMemberOffset(structName, memberName string) (uint64, error) {
	structure := r.GetStructure(structName)
	if structure == nil {
		return 0, fmt.Errorf("the structure not found: %s", structName)
	}
	field := structure.GetField(memberName)
	if field == nil {
		return 0, fmt.Errorf("the field not found, struct name: %s, member name: %s", structName, memberName)
	}
	return uint64(field.Offset), nil
}

func (r *DwarfReader) init(data *dwarf.Data, readAttrNames []string) error {
	r.functions = make(map[string]*FunctionInfo)
	r.structures = make(map[string]*StructureInfo)

	reader := data.Reader()
	r.elfByteOrder = reader.ByteOrder()
	for {
		entry, err := reader.Next()
		if err != nil {
			return fmt.Errorf("read dwarf error: %v", err)
		}
		if entry == nil {
			break
		}

		if err := r.processProducer(data, entry); err != nil {
			return err
		}
		if err := r.processFunctions(readAttrNames, data, entry); err != nil {
			return err
		}
		if err := r.processStructures(readAttrNames, data, entry); err != nil {
			return err
		}
	}
	return nil
}
