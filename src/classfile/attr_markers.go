package classfile

type DeprecatedAttribute struct {MarkerAttribute}
type SyntheticAttribute struct {MarkerAttribute}

type MarkerAttribute struct {}

func (ma *MarkerAttribute) readInfo(reader *ClassReader) {
	// read nothing
}
