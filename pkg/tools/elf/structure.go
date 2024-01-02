package elf

type StructureInfo struct {
	name   string
	fields []*StructureFieldInfo
}

type StructureFieldInfo struct {
	Name   string
	Offset int64
}

func (s *StructureInfo) GetField(name string) *StructureFieldInfo {
	for _, f := range s.fields {
		if f.Name == name {
			return f
		}
	}
	return nil
}
