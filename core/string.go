package core

import (
	"bytes"
	"fmt"
	"strconv"
)

func parseEncodedString(buffer *bytes.Buffer, encType EncodedStringType) (string, error) {
	switch encType {
	case EncInt8:
		return parseInt8(buffer)
	case EncInt16:
		return parseInt16(buffer)
	case EncInt32:
		return parseInt32(buffer)
	case EncLZF:
		return parseLZF(buffer)
	default:
		return ``, fmt.Errorf("invalid encoded string type %d", encType)
	}
}

func parseInt8(buffer *bytes.Buffer) (string, error) {
	b, err := buffer.ReadByte()
	if err != nil {
		return ``, fmt.Errorf("failed to read byte: %v", err)
	}
	return strconv.Itoa(int(b)), nil
}

func parseInt16(buffer *bytes.Buffer) (string, error) {
	return "", nil
}

func parseInt32(buffer *bytes.Buffer) (string, error) {
	return "", nil
}

func parseLZF(buffer *bytes.Buffer) (string, error) {
	return "", nil
}

func readString(buffer *bytes.Buffer) (string, error) {
	b, err := buffer.ReadByte()
	if err != nil {
		return ``, fmt.Errorf("failed to read byte: %v", err)
	}
	sign := LengthSign((b & signMask) >> 6)
	lenBytes := make([]byte, 0, 4)
	switch sign {
	case Len6Bit:
		nextByte, err := buffer.ReadByte()
		if err != nil {
			return ``, fmt.Errorf("failed to read byte: %v", err)
		}
		lenBytes = append(lenBytes, nextByte&unsignedMask)
	case Len14Bit:
		nextByte, err := buffer.ReadByte()
		if err != nil {
			return ``, fmt.Errorf("failed to read byte: %v", err)
		}
		lenBytes = append(lenBytes, b&unsignedMask)
		lenBytes = append(lenBytes, nextByte)
	case Len32Bit:
		nextBytes := make([]byte, 4)
		_, err = buffer.Read(nextBytes)
		if err != nil {
			return ``, fmt.Errorf("failed to read bytes: %v", err)
		}
		lenBytes = append(lenBytes, nextBytes...)
	case Encoded:
		return parseEncodedString(buffer, EncodedStringType(b&unsignedMask))
	default:
		return ``, fmt.Errorf("invalid lengthsign %d", sign)

	}
	length := 0
	for _, l := range lenBytes {
		length <<= 8
		length += int(l)
	}
	return length, nil
}
