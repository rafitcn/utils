package color

import "fmt"

const (
	BLACK = iota + 30
	RED
	GREEN
	YELLOW
	BLUE
	MAGENTA
	CYAN
	WHITE
)

const (
	TypeNone      = 0
	TypeHighlight = 1
	TypeUnderline = 4
	TypeShine     = 5
	TypeReverse   = 7
	TypeHidden    = 8
)

func Black(msg string, conf ...int) string {
	return Colorful(msg, setColor(BLACK, conf...)...)
}

func Red(msg string, conf ...int) string {
	return Colorful(msg, setColor(RED, conf...)...)
}

func Green(msg string, conf ...int) string {
	return Colorful(msg, setColor(GREEN, conf...)...)
}

func Yellow(msg string, conf ...int) string {
	return Colorful(msg, setColor(YELLOW, conf...)...)
}

func Blue(msg string, conf ...int) string {
	return Colorful(msg, setColor(BLUE, conf...)...)
}

func Magenta(msg string, conf ...int) string {
	return Colorful(msg, setColor(MAGENTA, conf...)...)
}

func Cyan(msg string, conf ...int) string {
	return Colorful(msg, setColor(CYAN, conf...)...)
}

func White(msg string, conf ...int) string {
	return Colorful(msg, setColor(WHITE, conf...)...)
}

// Colorful set msg color
// conf[0] is foreground, 30-37
// conf[1] is showType, 0-8
// conf[2] is background, 40-47
func Colorful(msg string, conf ...int) string {
	if len(conf) == 0 {
		return msg
	}
	var showType, foreground, background int
	switch len(conf) {
	case 1:
		foreground = conf[0]
	case 2:
		foreground = conf[0]
		showType = conf[1]
	default:
		foreground = conf[0]
		showType = conf[1]
		background = conf[2]
	}

	return fmt.Sprintf("%c[%d;%d;%dm%s%c[0m", 0x1B, showType, background, foreground, msg, 0x1B)
}

func setColor(color int, conf ...int) []int {
	switch len(conf) {
	case 0:
		return []int{color, 0, 0}
	case 1:
		return []int{color, conf[0], 0}
	default:
		return []int{color, conf[0], conf[1]}
	}
}
