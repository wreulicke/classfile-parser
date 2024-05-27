package code

import (
	"bytes"
	"errors"
	"fmt"
	"io"

	"github.com/wreulicke/classfile-parser/binary"
)

type codeParser struct {
	binary.Parser

	opcodeParseFns map[opcode]opcodeParseFn
}

type opcodeParseFn func(*Instruction) error

type CodeParser interface { //nolint:revive
	Parse() ([]*Instruction, error)
}

type Instruction struct {
	opcode   opcode
	operands []uint8
}

func NewCodeParser(code []byte) CodeParser {
	buf := bytes.NewBuffer(code)
	p := &codeParser{
		Parser:         binary.NewParser(buf),
		opcodeParseFns: map[opcode]opcodeParseFn{},
	}
	p.registerOpcodeParseFn(Aaload, p.nop)
	p.registerOpcodeParseFn(Aastore, p.nop)
	p.registerOpcodeParseFn(Aconst_null, p.nop)

	p.registerOpcodeParseFn(Aload, p.take1Operand)
	p.registerOpcodeParseFn(Aload_0, p.nop)
	p.registerOpcodeParseFn(Aload_1, p.nop)
	p.registerOpcodeParseFn(Aload_2, p.nop)
	p.registerOpcodeParseFn(Aload_3, p.nop)

	p.registerOpcodeParseFn(Anewarray, p.take2Operand)
	p.registerOpcodeParseFn(Areturn, p.nop)
	p.registerOpcodeParseFn(Arraylength, p.nop)

	p.registerOpcodeParseFn(Astore, p.take1Operand)
	p.registerOpcodeParseFn(Astore_0, p.nop)
	p.registerOpcodeParseFn(Astore_1, p.nop)
	p.registerOpcodeParseFn(Astore_2, p.nop)
	p.registerOpcodeParseFn(Astore_3, p.nop)

	p.registerOpcodeParseFn(Athrow, p.nop)
	p.registerOpcodeParseFn(Baload, p.nop)
	p.registerOpcodeParseFn(Bastore, p.nop)
	p.registerOpcodeParseFn(Bipush, p.take1Operand)

	p.registerOpcodeParseFn(Caload, p.nop)
	p.registerOpcodeParseFn(Castore, p.nop)
	p.registerOpcodeParseFn(Checkcast, p.take2Operand)

	p.registerOpcodeParseFn(D2f, p.nop)
	p.registerOpcodeParseFn(D2i, p.nop)
	p.registerOpcodeParseFn(D2l, p.nop)
	p.registerOpcodeParseFn(Dadd, p.nop)
	p.registerOpcodeParseFn(Daload, p.nop)
	p.registerOpcodeParseFn(Dastore, p.nop)
	p.registerOpcodeParseFn(Dcmpg, p.nop)
	p.registerOpcodeParseFn(Dcmpl, p.nop)
	p.registerOpcodeParseFn(Dconst_0, p.nop)
	p.registerOpcodeParseFn(Dconst_1, p.nop)
	p.registerOpcodeParseFn(Ddiv, p.nop)
	p.registerOpcodeParseFn(Dload, p.take1Operand)
	p.registerOpcodeParseFn(Dload_0, p.nop)
	p.registerOpcodeParseFn(Dload_1, p.nop)
	p.registerOpcodeParseFn(Dload_2, p.nop)
	p.registerOpcodeParseFn(Dload_3, p.nop)
	p.registerOpcodeParseFn(Dmul, p.nop)
	p.registerOpcodeParseFn(Dneg, p.nop)
	p.registerOpcodeParseFn(Drem, p.nop)
	p.registerOpcodeParseFn(Dreturn, p.nop)
	p.registerOpcodeParseFn(Dstore, p.take1Operand)
	p.registerOpcodeParseFn(Dstore_0, p.nop)
	p.registerOpcodeParseFn(Dstore_1, p.nop)
	p.registerOpcodeParseFn(Dstore_2, p.nop)
	p.registerOpcodeParseFn(Dstore_3, p.nop)
	p.registerOpcodeParseFn(Dsub, p.nop)

	p.registerOpcodeParseFn(Dup, p.nop)
	p.registerOpcodeParseFn(Dup_x1, p.nop)
	p.registerOpcodeParseFn(Dup_x2, p.nop)
	p.registerOpcodeParseFn(Dup2, p.nop)
	p.registerOpcodeParseFn(Dup2_x1, p.nop)
	p.registerOpcodeParseFn(Dup2_x2, p.nop)

	p.registerOpcodeParseFn(F2d, p.nop)
	p.registerOpcodeParseFn(F2i, p.nop)
	p.registerOpcodeParseFn(F2l, p.nop)
	p.registerOpcodeParseFn(Fadd, p.nop)
	p.registerOpcodeParseFn(Faload, p.nop)
	p.registerOpcodeParseFn(Fastore, p.nop)
	p.registerOpcodeParseFn(Fcmpg, p.nop)
	p.registerOpcodeParseFn(Fcmpl, p.nop)
	p.registerOpcodeParseFn(Fconst_0, p.nop)
	p.registerOpcodeParseFn(Fconst_1, p.nop)
	p.registerOpcodeParseFn(Fconst_2, p.nop)
	p.registerOpcodeParseFn(Fdiv, p.nop)
	p.registerOpcodeParseFn(Fload, p.take1Operand)
	p.registerOpcodeParseFn(Fload_0, p.nop)
	p.registerOpcodeParseFn(Fload_1, p.nop)
	p.registerOpcodeParseFn(Fload_2, p.nop)
	p.registerOpcodeParseFn(Fload_3, p.nop)
	p.registerOpcodeParseFn(Fmul, p.nop)
	p.registerOpcodeParseFn(Fneg, p.nop)
	p.registerOpcodeParseFn(Frem, p.nop)
	p.registerOpcodeParseFn(Freturn, p.nop)
	p.registerOpcodeParseFn(Fstore, p.take1Operand)
	p.registerOpcodeParseFn(Fstore_0, p.nop)
	p.registerOpcodeParseFn(Fstore_1, p.nop)
	p.registerOpcodeParseFn(Fstore_2, p.nop)
	p.registerOpcodeParseFn(Fstore_3, p.nop)
	p.registerOpcodeParseFn(Fsub, p.nop)

	p.registerOpcodeParseFn(Getfield, p.take2Operand)
	p.registerOpcodeParseFn(Getstatic, p.take2Operand)
	p.registerOpcodeParseFn(Goto, p.take2Operand)
	p.registerOpcodeParseFn(Goto_w, p.take4Operand)

	p.registerOpcodeParseFn(I2b, p.nop)
	p.registerOpcodeParseFn(I2c, p.nop)
	p.registerOpcodeParseFn(I2d, p.nop)
	p.registerOpcodeParseFn(I2f, p.nop)
	p.registerOpcodeParseFn(I2l, p.nop)
	p.registerOpcodeParseFn(I2s, p.nop)
	p.registerOpcodeParseFn(Iadd, p.nop)
	p.registerOpcodeParseFn(Iaload, p.nop)
	p.registerOpcodeParseFn(Iand, p.nop)
	p.registerOpcodeParseFn(Iastore, p.nop)
	p.registerOpcodeParseFn(Iconst_m1, p.nop)
	p.registerOpcodeParseFn(Iconst_0, p.nop)
	p.registerOpcodeParseFn(Iconst_1, p.nop)
	p.registerOpcodeParseFn(Iconst_2, p.nop)
	p.registerOpcodeParseFn(Iconst_3, p.nop)
	p.registerOpcodeParseFn(Iconst_4, p.nop)
	p.registerOpcodeParseFn(Iconst_5, p.nop)
	p.registerOpcodeParseFn(Idiv, p.nop)

	p.registerOpcodeParseFn(If_acmpeq, p.take2Operand)
	p.registerOpcodeParseFn(If_acmpne, p.take2Operand)
	p.registerOpcodeParseFn(If_icmpeq, p.take2Operand)
	p.registerOpcodeParseFn(If_icmpne, p.take2Operand)
	p.registerOpcodeParseFn(If_icmplt, p.take2Operand)
	p.registerOpcodeParseFn(If_icmpge, p.take2Operand)
	p.registerOpcodeParseFn(If_icmpgt, p.take2Operand)
	p.registerOpcodeParseFn(If_icmple, p.take2Operand)

	p.registerOpcodeParseFn(Ifeq, p.take2Operand)
	p.registerOpcodeParseFn(Ifne, p.take2Operand)
	p.registerOpcodeParseFn(Iflt, p.take2Operand)
	p.registerOpcodeParseFn(Ifge, p.take2Operand)
	p.registerOpcodeParseFn(Ifgt, p.take2Operand)
	p.registerOpcodeParseFn(Ifle, p.take2Operand)

	p.registerOpcodeParseFn(Ifnonnull, p.nop)
	p.registerOpcodeParseFn(Ifnull, p.nop)

	p.registerOpcodeParseFn(Iinc, p.take2Operand)
	p.registerOpcodeParseFn(Iload, p.take1Operand)
	p.registerOpcodeParseFn(Iload_0, p.nop)
	p.registerOpcodeParseFn(Iload_1, p.nop)
	p.registerOpcodeParseFn(Iload_2, p.nop)
	p.registerOpcodeParseFn(Iload_3, p.nop)

	p.registerOpcodeParseFn(Imul, p.nop)
	p.registerOpcodeParseFn(Ineg, p.nop)

	p.registerOpcodeParseFn(Instanceof, p.take2Operand)

	p.registerOpcodeParseFn(Invokedynamic, p.parseInvokedynamic)
	p.registerOpcodeParseFn(Invokeinterface, p.parseInvokeinterface)
	p.registerOpcodeParseFn(Invokespecial, p.take2Operand)
	p.registerOpcodeParseFn(Invokestatic, p.take2Operand)
	p.registerOpcodeParseFn(Invokevirtual, p.take2Operand)

	p.registerOpcodeParseFn(Ior, p.nop)
	p.registerOpcodeParseFn(Irem, p.nop)
	p.registerOpcodeParseFn(Ireturn, p.nop)
	p.registerOpcodeParseFn(Ishl, p.nop)
	p.registerOpcodeParseFn(Ishr, p.nop)
	p.registerOpcodeParseFn(Istore, p.take1Operand)
	p.registerOpcodeParseFn(Istore_0, p.nop)
	p.registerOpcodeParseFn(Istore_1, p.nop)
	p.registerOpcodeParseFn(Istore_2, p.nop)
	p.registerOpcodeParseFn(Istore_3, p.nop)
	p.registerOpcodeParseFn(Isub, p.nop)
	p.registerOpcodeParseFn(Iushr, p.nop)
	p.registerOpcodeParseFn(Ixor, p.nop)
	p.registerOpcodeParseFn(Jsr, p.take2Operand)
	p.registerOpcodeParseFn(Jsr_w, p.take4Operand)

	p.registerOpcodeParseFn(L2d, p.nop)
	p.registerOpcodeParseFn(L2f, p.nop)
	p.registerOpcodeParseFn(L2i, p.nop)

	p.registerOpcodeParseFn(Ladd, p.nop)
	p.registerOpcodeParseFn(Laload, p.nop)
	p.registerOpcodeParseFn(Land, p.nop)
	p.registerOpcodeParseFn(Lastore, p.nop)
	p.registerOpcodeParseFn(Lcmp, p.nop)
	p.registerOpcodeParseFn(Lconst_0, p.nop)
	p.registerOpcodeParseFn(Lconst_1, p.nop)
	p.registerOpcodeParseFn(Ldc, p.take1Operand)
	p.registerOpcodeParseFn(Ldc_w, p.take2Operand)
	p.registerOpcodeParseFn(Ldc2_w, p.take2Operand)
	p.registerOpcodeParseFn(Ldiv, p.nop)
	p.registerOpcodeParseFn(Lload, p.nop)
	p.registerOpcodeParseFn(Lload_0, p.nop)
	p.registerOpcodeParseFn(Lload_1, p.nop)
	p.registerOpcodeParseFn(Lload_2, p.nop)
	p.registerOpcodeParseFn(Lload_3, p.nop)
	p.registerOpcodeParseFn(Lmul, p.nop)
	p.registerOpcodeParseFn(Lneg, p.nop)
	p.registerOpcodeParseFn(Lookupswitch, p.parseLookupSwitch)
	p.registerOpcodeParseFn(Lload_2, p.nop)
	p.registerOpcodeParseFn(Lload_3, p.nop)
	p.registerOpcodeParseFn(Lmul, p.nop)
	p.registerOpcodeParseFn(Lneg, p.nop)

	p.registerOpcodeParseFn(Lor, p.nop)
	p.registerOpcodeParseFn(Lrem, p.nop)
	p.registerOpcodeParseFn(Lreturn, p.nop)
	p.registerOpcodeParseFn(Lshl, p.nop)
	p.registerOpcodeParseFn(Lshr, p.nop)
	p.registerOpcodeParseFn(Lstore, p.take1Operand)
	p.registerOpcodeParseFn(Lstore_0, p.nop)
	p.registerOpcodeParseFn(Lstore_1, p.nop)
	p.registerOpcodeParseFn(Lstore_2, p.nop)
	p.registerOpcodeParseFn(Lstore_3, p.nop)
	p.registerOpcodeParseFn(Lsub, p.nop)
	p.registerOpcodeParseFn(Lushr, p.nop)
	p.registerOpcodeParseFn(Lxor, p.nop)
	p.registerOpcodeParseFn(Monitorenter, p.nop)
	p.registerOpcodeParseFn(Monitorexit, p.nop)
	p.registerOpcodeParseFn(Multianewarray, p.parseMultianewarray)
	p.registerOpcodeParseFn(New, p.take2Operand)
	p.registerOpcodeParseFn(Newarray, p.take1Operand)
	p.registerOpcodeParseFn(Nop, p.nop)
	p.registerOpcodeParseFn(Pop, p.nop)
	p.registerOpcodeParseFn(Pop2, p.nop)
	p.registerOpcodeParseFn(Putfield, p.take2Operand)
	p.registerOpcodeParseFn(Putstatic, p.take2Operand)
	p.registerOpcodeParseFn(Ret, p.take1Operand)
	p.registerOpcodeParseFn(Return, p.nop)
	p.registerOpcodeParseFn(Saload, p.nop)
	p.registerOpcodeParseFn(Sastore, p.nop)
	p.registerOpcodeParseFn(Sipush, p.take2Operand)
	p.registerOpcodeParseFn(Swap, p.nop)
	p.registerOpcodeParseFn(Tableswitch, p.parseTableSwitch)
	p.registerOpcodeParseFn(Wide, p.parseWide)
	return p
}

func (p *codeParser) Parse() ([]*Instruction, error) {
	var instructions []*Instruction
	for {
		inst, err := p.parseInstruction()
		if errors.Is(err, io.EOF) {
			instructions = append(instructions, inst)
			return instructions, nil
		} else if err != nil {
			return instructions, err
		}
		instructions = append(instructions, inst)
	}
}

func (p *codeParser) parseInstruction() (*Instruction, error) {
	b, err := p.ReadUint8()
	if err != nil {
		return nil, err
	}
	parse, ok := p.opcodeParseFns[opcode(b)]
	if !ok {
		return nil, fmt.Errorf("unknown opcode. %d", b)
	}
	inst := &Instruction{
		opcode: opcode(b),
	}
	return inst, parse(inst)
}

func (p *codeParser) registerOpcodeParseFn(code opcode, fn opcodeParseFn) {
	p.opcodeParseFns[code] = fn
}

func (p *codeParser) nop(_ *Instruction) error {
	return nil
}

func (p *codeParser) take1Operand(inst *Instruction) error {
	b, err := p.ReadUint8()
	if err != nil {
		return err
	}
	inst.operands = append(inst.operands, b)
	return nil
}

func (p *codeParser) take2Operand(inst *Instruction) error {
	err := p.take1Operand(inst)
	if err != nil {
		return err
	}
	return p.take1Operand(inst)
}

func (p *codeParser) take4Operand(inst *Instruction) error {
	err := p.take2Operand(inst)
	if err != nil {
		return err
	}
	return p.take2Operand(inst)
}

func (p *codeParser) parseInvokedynamic(inst *Instruction) error {
	err := p.take2Operand(inst)
	if err != nil {
		return err
	}
	// skip 2 bytes
	_, err = p.ReadUint16()
	return err
}

func (p *codeParser) parseInvokeinterface(inst *Instruction) error {
	err := p.take2Operand(inst)
	if err != nil {
		return err
	}
	if err := p.take1Operand(inst); err != nil {
		return err
	}
	// skip 1 bytes
	_, err = p.ReadUint8()
	return err
}

func (p *codeParser) parseMultianewarray(inst *Instruction) error {
	err := p.take2Operand(inst)
	if err != nil {
		return err
	}
	return p.take1Operand(inst)
}

func (p *codeParser) parseLookupSwitch(inst *Instruction) error {
	var u uint8
	var err error
	// skip pad
	for {
		u, err = p.ReadUint8()
		if err != nil {
			return err
		}
		if u > 0 {
			break
		}
	}

	// read default
	if err := p.take4Operand(inst); err != nil {
		return err
	}
	npair, err := p.ReadUint32()
	if err != nil {
		return err
	}
	inst.operands = append(inst.operands, uint8(npair>>24))
	inst.operands = append(inst.operands, uint8(npair>>16))
	inst.operands = append(inst.operands, uint8(npair>>8))
	inst.operands = append(inst.operands, uint8(npair))
	for i := uint32(0); i < npair; i++ {
		err := p.take4Operand(inst)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *codeParser) parseTableSwitch(inst *Instruction) error {
	var u uint8
	var err error
	// skip pad
	for {
		u, err = p.ReadUint8()
		if err != nil {
			return err
		}
		if u > 0 {
			break
		}
	}

	// read default
	if err := p.take4Operand(inst); err != nil {
		return err
	}
	low, err := p.ReadUint32()
	if err != nil {
		return err
	}
	inst.operands = append(inst.operands, uint8(low>>24))
	inst.operands = append(inst.operands, uint8(low>>16))
	inst.operands = append(inst.operands, uint8(low>>8))
	inst.operands = append(inst.operands, uint8(low))

	high, err := p.ReadUint32()
	if err != nil {
		return err
	}
	inst.operands = append(inst.operands, uint8(high>>24))
	inst.operands = append(inst.operands, uint8(high>>16))
	inst.operands = append(inst.operands, uint8(high>>8))
	inst.operands = append(inst.operands, uint8(high))

	count := high - low + 1
	for i := uint32(0); i < count; i++ {
		err := p.take4Operand(inst)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *codeParser) parseWide(inst *Instruction) error {
	b, err := p.ReadUint8()
	if err != nil {
		return err
	}
	// TODO: fix data structure for instruction
	// Next is not actual operand, but added to operands here.
	inst.operands = append(inst.operands, b)

	switch opcode(b) {
	case Iload, Fload, Aload, Lload, Dload, Istore, Fstore, Astore, Lstore, Dstore, Ret:
		return p.take2Operand(inst)
	case Iinc:
		err := p.take2Operand(inst)
		if err != nil {
			return err
		}
		return p.take2Operand(inst)
	default:
		return fmt.Errorf("unknown opcode. %d", b)
	}
}
