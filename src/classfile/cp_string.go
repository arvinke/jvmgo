package classfile

type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
}

func (sl *ConstantStringInfo) readInfo(reader *ClassReader) {
	sl.stringIndex = reader.readUint16()
}

func (sl *ConstantStringInfo) String() string {
	return sl.cp.getUtf8(sl.stringIndex)
}
