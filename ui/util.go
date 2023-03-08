package ui

import (
	"fmt"
	"strings"
)

//	func fgRGB(r, g, b int) string {
//		return fmt.Sprintf("\x1b[38;2;%v;%v;%vm", r, g, b)
//	}
//
//	func bgRGB(r, g, b int) string {
//		return fmt.Sprintf("\x1b[48;2;%v;%v;%vm", r, g, b)
//	}
//
//	func fgDefault() string {
//		return "\x1b[1;39m"
//	}
//
//	func bgDefault() string {
//		return "\x1b[1;49m"
//	}

// "ESC[?25h <- cursor visible code"
func showCursr() string {
	return "\x1B[?25h"
}

func padding(size int) string {
	if size < 0 {
		return ""
	}
	return strings.Repeat(" ", size)
}

// indexed from 1 WTF!?
func setXY(x, y int) string {
	return fmt.Sprintf("\x1B[%v;%vH", y, x)
}

func italic(text string) string {
	return fmt.Sprintf("\x1B[3m%v\x1b[23m", text)
}

func bold(text string) string {
	return fmt.Sprintf("\x1B[1m%v\x1b[22m", text)
}

func clearLine() string {
	return "\x1B[2K"
}

func cut(text string, size int) string {
	if size > len(text) {
		return text
	}
	return fmt.Sprintf("%v…", text[:size-1])
}

func time(ms int) string {
	t := ""
	secs := ms / 1000
	min := secs / 60
	sec := secs % 60
	if min < 10 {
		t = fmt.Sprintf("0%v", min)
	} else {
		t = fmt.Sprintf("%v", min)
	}
	if sec < 10 {
		t = fmt.Sprintf("%v:0%v", t, sec)
	} else {
		t = fmt.Sprintf("%v:%v", t, sec)
	}
	return t
}
