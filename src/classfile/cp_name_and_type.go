package classfile

type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (sl *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	sl.nameIndex = reader.readUint16()
	sl.descriptorIndex = reader.readUint16()
}
