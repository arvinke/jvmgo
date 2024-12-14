package classfile

type ConstantMethodTypeInfo struct {
	descriptorIndex uint16
}

func (sl *ConstantMethodTypeInfo) readInfo(reader *ClassReader) {
	sl.descriptorIndex = reader.readUint16()
}

type ConstantMethodHandleInfo struct {
	referenceKind  uint8
	referenceIndex uint16
}

func (sl *ConstantMethodHandleInfo) readInfo(reader *ClassReader) {
	sl.referenceKind = reader.readUint8()
	sl.referenceIndex = reader.readUint16()
}

type ConstantInvokeDynamicInfo struct {
	bootstrapMethodAttrIndex uint16
	nameAndTypeIndex         uint16
}

func (sl *ConstantInvokeDynamicInfo) readInfo(reader *ClassReader) {
	sl.bootstrapMethodAttrIndex = reader.readUint16()
	sl.nameAndTypeIndex = reader.readUint16()
}
