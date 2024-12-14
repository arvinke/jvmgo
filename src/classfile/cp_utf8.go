package classfile

type ConstantUtf8Info struct {
	str string
}

func (sl *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := reader.readUint16()
	bytes := reader.readBytes(uint32(length))
	sl.str = decodeMUTF8(bytes)
}

func decodeMUTF8(bytes []byte) string {
	// 简化版，参考DataInputStream.readUTF
	return string(bytes)
}
