package commandtype

const (
	Value byte = iota
	LoadLibrary
	InvokeStaticMethod
	GetStaticField
	SetStaticField
	CreateClassInstance
	GetType
	Reference
	GetModule
	InvokeInstanceMethod
	Exception
	HeartBeat
	Cast
	GetInstanceField
	Optimize
	GenerateLib
	InvokeGlobalMethod
	DestructReference
	ArrayReference
	ArrayGetItem
	ArrayGetSize
	ArrayGetRank
	ArraySetItem
	Array
	RetrieveArray
	SetInstanceField
	InvokeGenericStaticMethod
	InvokeGenericMethod
	GetEnumItem
	GetEnumName
	GetEnumValue
	AsRef
	AsOut
	GetRefValue
)
