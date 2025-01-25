package core

// ValueType is the type of value
type ValueType int

const (
	TypeString          ValueType = 0
	TypeList            ValueType = 1
	TypeSet             ValueType = 2
	TypeZSet            ValueType = 3
	TypeHash            ValueType = 4
	TypeHashypeZSet2    ValueType = 5
	TypeModule          ValueType = 6
	TypeModule2         ValueType = 7
	TypeHashZipMap      ValueType = 9
	TypeListZipList     ValueType = 10
	TypeSetIntSet       ValueType = 11
	TypeZSetZipList     ValueType = 12
	TypeHashZipList     ValueType = 13
	TypeListQuickList   ValueType = 14
	TypeStreamListPacks ValueType = 15
)

// LengthSign is the type of length sign
type LengthSign uint8

const (
	Len6Bit  LengthSign = 0b00
	Len14Bit LengthSign = 0b01
	Len32Bit LengthSign = 0b10
	Encoded  LengthSign = 0b11
)

const unsignedMask = 0b00111111
const signMask = 0b11000000

// EncodedNumberType is the type of encoded number
type EncodedStringType int

const (
	EncInt8  EncodedStringType = 0
	EncInt16 EncodedStringType = 1
	EncInt32 EncodedStringType = 2
	EncLZF   EncodedStringType = 3
)
