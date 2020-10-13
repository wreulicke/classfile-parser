package parser

type AttributeStackMapTable struct {
	Entries []StackMapFrame
}

func (a *AttributeStackMapTable) Name() string {
	return "StackMapTable"
}

type StackMapFrame interface{}

type StackMapFrameSameFrame struct {
	FrameType uint8
}

type StackMapFrameSameLocals1StackItemFrame struct {
	FrameType uint8
	stack     VerificationTypeInfo
}

type StackMapFrameSameLocals1StackItemFrameExtended struct {
	FrameType   uint8
	OffsetDelta uint16
	stack       VerificationTypeInfo
}

type StackMapFrameChopFrame struct {
	FrameType   uint8
	OffsetDelta uint16
}

type StackMapFrameSameFrameExtended struct {
	FrameType   uint8
	OffsetDelta uint16
}

type StackMapFrameAppendFrame struct {
	FrameType   uint8
	OffsetDelta uint16
	Locals      []VerificationTypeInfo
}

type StackMapFrameFullFrame struct {
	FrameType   uint8
	OffsetDelta uint16
	Locals      []VerificationTypeInfo
	Stacks      []VerificationTypeInfo
}

type VerificationTypeInfo interface{}

var (
	_verificationTypeInfoTopVaribleInfo               = &VerificationTypeInfoDoubleVaribleInfo{}
	_verificationTypeInfoIntegerVaribleInfo           = &VerificationTypeInfoIntegerVaribleInfo{}
	_verificationTypeInfoFloatVaribleInfo             = &VerificationTypeInfoIntegerVaribleInfo{}
	_verificationTypeInfoNullVaribleInfo              = &VerificationTypeInfoNullVaribleInfo{}
	_verificationTypeInfoUninitializedThisVaribleInfo = &VerificationTypeInfoUninitializedThisVaribleInfo{}
	_verificationTypeInfoLongVaribleInfo              = &VerificationTypeInfoLongVaribleInfo{}
	_verificationTypeInfoDoubleVaribleInfo            = &VerificationTypeInfoDoubleVaribleInfo{}
)

type VerificationTypeInfoTopVaribleInfo struct{}
type VerificationTypeInfoIntegerVaribleInfo struct{}
type VerificationTypeInfoFloatVaribleInfo struct{}
type VerificationTypeInfoNullVaribleInfo struct{}
type VerificationTypeInfoUninitializedThisVaribleInfo struct{}
type VerificationTypeInfoObjectVaribleInfo struct {
	CpoolIndex uint16
}
type VerificationTypeInfoUninitializedVaribleInfo struct {
	Offset uint16
}
type VerificationTypeInfoLongVaribleInfo struct{}
type VerificationTypeInfoDoubleVaribleInfo struct{}
