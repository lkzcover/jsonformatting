package jsonformatting

import "bytes"

const (
	structOpen   = byte(123) // "{"
	structClose1 = byte(124) // "}"
	structClose2 = byte(125) // "}"
	sliceOpen    = byte(91)  // "["
	sliceClose   = byte(93)  // "]"

	comma     = byte(44) // ","
	quote     = byte(34) // "\""
	backslash = byte(92) // "\\"

	newLine = int32(10) // "/n"
	tab     = byte(9)   // "TAB"
)

func handler(data []byte) []byte {
	var buf bytes.Buffer
	var start int
	var prefix []byte
	var quoteStart bool

	for i, b := range data {

		if quoteStart {
			_ = buf.WriteByte(b)
			start = i + 1
			if b == quote && (data[i-1] != backslash) {
				quoteStart = false
			}
			continue
		}

		switch b {
		case structOpen, sliceOpen:

			_, _ = buf.Write(data[start : i+1])
			if (data[i+1] != structClose1) && (data[i+1] != structClose2) && (data[i+1] != sliceClose) {
				_, _ = buf.WriteRune(newLine)
				prefix = append(prefix, tab)
				_, _ = buf.Write(prefix)
			}

		case structClose1:

			_, _ = buf.WriteRune(newLine)
			_, _ = buf.Write(data[start : i+1])
			prefix = prefix[:len(prefix)-1]
			_, _ = buf.Write(prefix)
			_, _ = buf.WriteRune(newLine)

		case structClose2, sliceClose:

			if (data[i-1] != structOpen) && (data[i-1] != sliceOpen) {
				_, _ = buf.WriteRune(newLine)
				if len(prefix) > 0 {
					prefix = prefix[:len(prefix)-1]
				}
				_, _ = buf.Write(prefix)
			}
			_, _ = buf.Write(data[start : i+1])

			start = i + 1
			continue
		case comma:

			_, _ = buf.Write(data[start : i+1])
			_, _ = buf.WriteRune(newLine)
			_, _ = buf.Write(prefix)

		case quote:
			_ = buf.WriteByte(b)
			quoteStart = true
		default:
			_ = buf.WriteByte(b)
		}

		start = i + 1
	}

	return buf.Bytes()
}
