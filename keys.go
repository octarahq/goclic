package menu

const (
	KeyTab         byte = 9
	KeyEnter       byte = 13
	KeyEsc         byte = 27
	KeySpace       byte = 32
	KeyBackspace   byte = 127
	KeyCtrlC       byte = 3
	KeyCtrlD       byte = 4
	KeyCtrlV       byte = 22
	KeyOpenBracket byte = 91
	KeySeqLen      byte = 3
)

const (
	ArrowUp    byte = 65
	ArrowDown  byte = 66
	ArrowRight byte = 67
	ArrowLeft  byte = 68
)

const (
	KeyInsert   byte = 50
	KeyDelete   byte = 51
	KeyPageUp   byte = 53
	KeyPageDown byte = 54
	KeyHome     byte = 72
	KeyEnd      byte = 70
)

const (
	KeyF1  byte = 80
	KeyF2  byte = 81
	KeyF3  byte = 82
	KeyF4  byte = 83
	KeyF5  byte = 53
	KeyF6  byte = 17
	KeyF7  byte = 18
	KeyF8  byte = 19
	KeyF9  byte = 20
	KeyF10 byte = 21
	KeyF11 byte = 23
	KeyF12 byte = 24
)

func IsPrintable(key []byte) bool {
	if len(key) != 1 {
		return false
	}

	return key[0] >= 32 && key[0] <= 126
}

func IsAlpha(key []byte) bool {
	if len(key) != 1 {
		return false
	}
	b := key[0]
	return (b >= 65 && b <= 90) || (b >= 97 && b <= 122)
}

func IsNumeric(key []byte) bool {
	if len(key) != 1 {
		return false
	}
	return key[0] >= 48 && key[0] <= 57
}
