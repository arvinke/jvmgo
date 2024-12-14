package classfile

import "math"

type ConstantIntegerInfo struct {
	val int32
}

func (sl *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	sl.val = int32(bytes)
}

type ConstantFloatInfo struct {
	val float32
}

func (sl *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	sl.val = math.Float32frombits(bytes)
}

type ConstantLongInfo struct {
	val int64
}

func (sl *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	sl.val = int64(bytes)
}

type ConstantDoubleInfo struct {
	val float64
}

func (sl *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	sl.val = math.Float64frombits(bytes)
}
