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
	_verificationTypeInfoTopVaribleInfo               = &VerificationTypeInfoDoubleVaribleInfo{}            //nolint:gochecknoglobals
	_verificationTypeInfoIntegerVaribleInfo           = &VerificationTypeInfoIntegerVaribleInfo{}           //nolint:gochecknoglobals
	_verificationTypeInfoFloatVaribleInfo             = &VerificationTypeInfoIntegerVaribleInfo{}           //nolint:gochecknoglobals
	_verificationTypeInfoNullVaribleInfo              = &VerificationTypeInfoNullVaribleInfo{}              //nolint:gochecknoglobals
	_verificationTypeInfoUninitializedThisVaribleInfo = &VerificationTypeInfoUninitializedThisVaribleInfo{} //nolint:gochecknoglobals
	_verificationTypeInfoLongVaribleInfo              = &VerificationTypeInfoLongVaribleInfo{}              //nolint:gochecknoglobals
	_verificationTypeInfoDoubleVaribleInfo            = &VerificationTypeInfoDoubleVaribleInfo{}            //nolint:gochecknoglobals
)

type (
	VerificationTypeInfoTopVaribleInfo               struct{}
	VerificationTypeInfoIntegerVaribleInfo           struct{}
	VerificationTypeInfoFloatVaribleInfo             struct{}
	VerificationTypeInfoNullVaribleInfo              struct{}
	VerificationTypeInfoUninitializedThisVaribleInfo struct{}
	VerificationTypeInfoObjectVaribleInfo            struct {
		CpoolIndex uint16
	}
)

type VerificationTypeInfoUninitializedVaribleInfo struct {
	Offset uint16
}
type (
	VerificationTypeInfoLongVaribleInfo   struct{}
	VerificationTypeInfoDoubleVaribleInfo struct{}
)
