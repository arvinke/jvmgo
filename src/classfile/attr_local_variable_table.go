package classfile

type LocalVariableTableAttribute struct {
	localVariableTable []*LocalVariableTableEntry
}

type LocalVariableTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

func (attr *LocalVariableTableAttribute) readInfo(reader *ClassReader) {
	localVariableTableLength := reader.readUint16()
	attr.localVariableTable = make([]*LocalVariableTableEntry, localVariableTableLength)
	for i := range attr.localVariableTable {
		attr.localVariableTable[i] = &LocalVariableTableEntry{
			startPc:    reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
}
