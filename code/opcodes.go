package code

type opcode uint8

// Instruction set is here.
// https://docs.oracle.com/javase/specs/jvms/se19/html/jvms-6.html#jvms-6.5
const (
	_               opcode = iota
	Aaload                 = 50  // (0x32)
	Aastore                = 83  // (0x53)
	Aconst_null            = 1   // (0x1)
	Aload                  = 25  // (0x19)
	Aload_0                = 42  // (0x2a)
	Aload_1                = 43  // (0x2b)
	Aload_2                = 44  // (0x2c)
	Aload_3                = 45  // (0x2d)
	Anewarray              = 189 // (0xbd)
	Areturn                = 176 // (0xb0)
	Arraylength            = 190 // (0xbe)
	Astore                 = 58  // (0x3a)
	Astore_0               = 75  // (0x4b)
	Astore_1               = 76  // (0x4c)
	Astore_2               = 77  // (0x4d)
	Astore_3               = 78  // (0x4e)
	Athrow                 = 191 // (0xbf)
	Baload                 = 51  // (0x33)
	Bastore                = 84  // (0x54)
	Bipush                 = 16  // (0x10)
	Caload                 = 52  // (0x34)
	Castore                = 85  // (0x55)
	Checkcast              = 192 // (0xc0)
	D2f                    = 144 // (0x90)
	D2i                    = 142 // (0x8e)
	D2l                    = 143 // (0x8f)
	Dadd                   = 99  // (0x63)
	Daload                 = 49  // (0x31)
	Dastore                = 82  // (0x52)
	Dcmpg                  = 152 // (0x98)
	Dcmpl                  = 151 // (0x97)
	Dconst_0               = 14  // (0xe)
	Dconst_1               = 15  // (0xf)
	Ddiv                   = 111 // (0x6f)
	Dload                  = 24  // (0x18)
	Dload_0                = 38  // (0x26)
	Dload_1                = 39  // (0x27)
	Dload_2                = 40  // (0x28)
	Dload_3                = 41  // (0x29)
	Dmul                   = 107 // (0x6b)
	Dneg                   = 119 // (0x77)
	Drem                   = 115 // (0x73)
	Dreturn                = 175 // (0xaf)
	Dstore                 = 57  // (0x39)
	Dstore_0               = 71  // (0x47)
	Dstore_1               = 72  // (0x48)
	Dstore_2               = 73  // (0x49)
	Dstore_3               = 74  // (0x4a)
	Dsub                   = 103 // (0x67)
	Dup                    = 89  // (0x59)
	Dup_x1                 = 90  // (0x5a)
	Dup_x2                 = 91  // (0x5b)
	Dup2                   = 92  // (0x5c)
	Dup2_x1                = 93  // (0x5d)
	Dup2_x2                = 94  // (0x5e)
	F2d                    = 141 // (0x8d)
	F2i                    = 139 // (0x8b)
	F2l                    = 140 // (0x8c)
	Fadd                   = 98  // (0x62)
	Faload                 = 48  // (0x30)
	Fastore                = 81  // (0x51)
	Fcmpg                  = 150 // (0x96)
	Fcmpl                  = 149 // (0x95)
	Fconst_0               = 11  // (0xb)
	Fconst_1               = 12  // (0xc)
	Fconst_2               = 13  // (0xd)
	Fdiv                   = 110 // (0x6e)
	Fload                  = 23  // (0x17)
	Fload_0                = 34  // (0x22)
	Fload_1                = 35  // (0x23)
	Fload_2                = 36  // (0x24)
	Fload_3                = 37  // (0x25)
	Fmul                   = 106 // (0x6a)
	Fneg                   = 118 // (0x76)
	Frem                   = 114 // (0x72)
	Freturn                = 174 // (0xae)
	Fstore                 = 56  // (0x38)
	Fstore_0               = 67  // (0x43)
	Fstore_1               = 68  // (0x44)
	Fstore_2               = 69  // (0x45)
	Fstore_3               = 70  // (0x46)
	Fsub                   = 102 // (0x66)
	Getfield               = 180 // (0xb4)
	Getstatic              = 178 // (0xb2)
	Goto                   = 167 // (0xa7)
	Goto_w                 = 200 // (0xc8)
	I2b                    = 145 // (0x91)
	I2c                    = 146 // (0x92)
	I2d                    = 135 // (0x87)
	I2f                    = 134 // (0x86)
	I2l                    = 133 // (0x85)
	I2s                    = 147 // (0x93)
	Iadd                   = 96  // (0x60)
	Iaload                 = 46  // (0x2e)
	Iand                   = 126 // (0x7e)
	Iastore                = 79  // (0x4f)
	Iconst_m1              = 2   // (0x2)
	Iconst_0               = 3   // (0x3)
	Iconst_1               = 4   // (0x4)
	Iconst_2               = 5   // (0x5)
	Iconst_3               = 6   // (0x6)
	Iconst_4               = 7   // (0x7)
	Iconst_5               = 8   // (0x8)
	Idiv                   = 108 // (0x6c)
	If_acmpeq              = 165 // (0xa5)
	If_acmpne              = 166 // (0xa6)
	If_icmpeq              = 159 // (0x9f)
	If_icmpne              = 160 // (0xa0)
	If_icmplt              = 161 // (0xa1)
	If_icmpge              = 162 // (0xa2)
	If_icmpgt              = 163 // (0xa3)
	If_icmple              = 164 // (0xa4)
	Ifeq                   = 153 // (0x99)
	Ifne                   = 154 // (0x9a)
	Iflt                   = 155 // (0x9b)
	Ifge                   = 156 // (0x9c)
	Ifgt                   = 157 // (0x9d)
	Ifle                   = 158 // (0x9e)
	Ifnonnull              = 199 // (0xc7)
	Ifnull                 = 198 // (0xc6)
	Iinc                   = 132 // (0x84)
	Iload                  = 21  // (0x15)
	Iload_0                = 26  // (0x1a)
	Iload_1                = 27  // (0x1b)
	Iload_2                = 28  // (0x1c)
	Iload_3                = 29  // (0x1d)
	Imul                   = 104 // (0x68)
	Ineg                   = 116 // (0x74)
	Instanceof             = 193 // (0xc1)
	Invokedynamic          = 186 // (0xba)
	Invokeinterface        = 185 // (0xb9)
	Invokespecial          = 183 // (0xb7)
	Invokestatic           = 184 // (0xb8)
	Invokevirtual          = 182 // (0xb6)
	Ior                    = 128 // (0x80)
	Irem                   = 112 // (0x70)
	Ireturn                = 172 // (0xac)
	Ishl                   = 120 // (0x78)
	Ishr                   = 122 // (0x7a)
	Istore                 = 54  // (0x36)
	Istore_0               = 59  // (0x3b)
	Istore_1               = 60  // (0x3c)
	Istore_2               = 61  // (0x3d)
	Istore_3               = 62  // (0x3e)
	Isub                   = 100 // (0x64)
	Iushr                  = 124 // (0x7c)
	Ixor                   = 130 // (0x82)
	Jsr                    = 168 // (0xa8)
	Jsr_w                  = 201 // (0xc9)
	L2d                    = 138 // (0x8a)
	L2f                    = 137 // (0x89)
	L2i                    = 136 // (0x88)
	Ladd                   = 97  // (0x61)
	Laload                 = 47  // (0x2f)
	Land                   = 127 // (0x7f)
	Lastore                = 80  // (0x50)
	Lcmp                   = 148 // (0x94)
	Lconst_0               = 9   // (0x9)
	Lconst_1               = 10  // (0xa)
	Ldc                    = 18  // (0x12)
	Ldc_w                  = 19  // (0x13)
	Ldc2_w                 = 20  // (0x14)
	Ldiv                   = 109 // (0x6d)
	Lload                  = 22  // (0x16)
	Lload_0                = 30  // (0x1e)
	Lload_1                = 31  // (0x1f)
	Lload_2                = 32  // (0x20)
	Lload_3                = 33  // (0x21)
	Lmul                   = 105 // (0x69)
	Lneg                   = 117 // (0x75)
	Lookupswitch           = 171 // (0xab)
	Lor                    = 129 // (0x81)
	Lrem                   = 113 // (0x71)
	Lreturn                = 173 // (0xad)
	Lshl                   = 121 // (0x79)
	Lshr                   = 123 // (0x7b)
	Lstore                 = 55  // (0x37)
	Lstore_0               = 63  // (0x3f)
	Lstore_1               = 64  // (0x40)
	Lstore_2               = 65  // (0x41)
	Lstore_3               = 66  // (0x42)
	Lsub                   = 101 // (0x65)
	Lushr                  = 125 // (0x7d)
	Lxor                   = 131 // (0x83)
	Monitorenter           = 194 // (0xc2)
	Monitorexit            = 195 // (0xc3)
	Multianewarray         = 197 // (0xc5)
	New                    = 187 // (0xbb)
	Newarray               = 188 // (0xbc)
	Nop                    = 0   // (0x0)
	Pop                    = 87  // (0x57)
	Pop2                   = 88  // (0x58)
	Putfield               = 181 // (0xb5)
	Putstatic              = 179 // (0xb3)
	Ret                    = 169 // (0xa9)
	Return                 = 177 // (0xb1)
	Saload                 = 53  // (0x35)
	Sastore                = 86  // (0x56)
	Sipush                 = 17  // (0x11)
	Swap                   = 95  // (0x5f)
	Tableswitch            = 170 // (0xaa)
)