package classfile

type ConstantMemberrefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (sl *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
	sl.classIndex = reader.readUint16()
	sl.nameAndTypeIndex = reader.readUint16()
}

func (sl *ConstantMemberrefInfo) ClassName() string {
	return sl.cp.getClassName(sl.classIndex)
}

func (sl *ConstantMemberrefInfo) NameAndType() (string, string) {
	return sl.cp.getNameAndType(sl.nameAndTypeIndex)
}

type ConstantFieldrefInfo struct {
	ConstantMemberrefInfo
}
type ConstantMethodrefInfo struct {
	ConstantMemberrefInfo
}
type ConstantInterfaceMethodrefInfo struct {
	ConstantMemberrefInfo
}
