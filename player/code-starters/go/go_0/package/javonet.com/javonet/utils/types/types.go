package types

const (
	JavonetCommand byte = iota
	JavonetString
	JavonetInteger
	JavonetBoolean
	JavonetFloat
	JavonetByte
	JavonetChar
	JavonetLongLong
	JavonetDouble
	JavonetUnsignedLongLong
	JavonetUnsignedInteger
)

const (
	JavonetBooleanSize          = 1
	JavonetByteSize             = 1
	JavonetCharSize             = 1
	JavonetIntegerSize          = 4
	JavonetFloatSize            = 4
	JavonetLongLongSize         = 8
	JavonetDoubleSize           = 8
	JavonetUnsignedLongLongSize = 8
	JavonetUnsignedIntegerSize  = 4
)
