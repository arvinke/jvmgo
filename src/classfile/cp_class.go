package classfile

type ConstantClassInfo struct {
	cp        ConstantPool
	nameIndex uint16
}

func (sl *ConstantClassInfo) readInfo(reader *ClassReader) {
	sl.nameIndex = reader.readUint16()
}
func (sl *ConstantClassInfo) Name() string {
	return sl.cp.getUtf8(sl.nameIndex)
}
