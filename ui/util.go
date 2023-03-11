package ui

import (
	"fmt"
	"strings"
)

// "ESC[?25h <- cursor visible code"
func hideCurs() string {
	return "\x1B[?25l"
}
func showCursr() string {
	return "\x1B[?25h"
}

func padding(size int) string {
	if size < 0 {
		return ""
	}
	return strings.Repeat(" ", size)
}

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

func cut(text []rune, size int) string {
	if size > len(text) {
		return string(text)
	}
	return fmt.Sprintf("%s… ", string(text[:size-2]))
}

func Time(ms int) string {
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
