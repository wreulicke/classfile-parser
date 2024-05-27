package parser

type AccessFlags uint16

// See https://docs.oracle.com/javase/specs/jvms/se22/html/jvms-4.html#jvms-4.1-200-E.1 for classfile
// See https://docs.oracle.com/javase/specs/jvms/se22/html/jvms-4.html#jvms-4.5-200-A.1 for field
// See https://docs.oracle.com/javase/specs/jvms/se22/html/jvms-4.html#jvms-4.6-200-A.1 for method
// See https://docs.oracle.com/javase/specs/jvms/se22/html/jvms-4.html#jvms-4.7.6-300-D.1-D.1 for inner class
const (
	ACC_PUBLIC       AccessFlags = 0x0001
	ACC_PRIVATE      AccessFlags = 0x0002
	ACC_PROTECTED    AccessFlags = 0x0004
	ACC_STATIC       AccessFlags = 0x0008
	ACC_FINAL        AccessFlags = 0x0010
	ACC_SUPER        AccessFlags = 0x0020
	ACC_SYNCHRONIZED AccessFlags = 0x0020
	ACC_BRIDGE       AccessFlags = 0x0040
	ACC_VOLATILE     AccessFlags = 0x0040
	ACC_VARARGS      AccessFlags = 0x0080
	ACC_TRANSIENT    AccessFlags = 0x0080
	ACC_NATIVE       AccessFlags = 0x0100
	ACC_ABSTRACT     AccessFlags = 0x0400
	ACC_STRICT       AccessFlags = 0x0800
	ACC_SYNTHETIC    AccessFlags = 0x1000
	ACC_ANNOTATION   AccessFlags = 0x2000
	ACC_ENUM         AccessFlags = 0x4000
	ACC_MANDATED     AccessFlags = 0x8000 // See https://docs.oracle.com/javase/specs/jvms/se22/html/jvms-4.html#jvms-4.7.24 in MethodParameters attribute
	ACC_MODULE       AccessFlags = 0x8000
)
