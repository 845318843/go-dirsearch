package utils

import (
	"fmt"
	"godir/common"
	"strings"
)

type Bar struct {
	percent uint8
	total   uint32
	current uint32
	grath   string
	rate    string
}

func NewBarOption(start, total uint32) *Bar {
	var b Bar
	b.total = total
	b.current = start
	b.current = 0
	if b.grath == "" {
		b.grath = "#"
	}
	b.percent = b.Getpercent()
	return &b
}

func (b *Bar) Getpercent() uint8 {
	return uint8(float32(b.current) / float32(b.total) * 100)
}

func (b *Bar) GetCurrent() uint32 {
	return b.current
}

func (b *Bar) Play() {
	b.percent = b.Getpercent()
	b.rate = strings.Repeat(b.grath, int(b.percent)/4)
	common.Glock.Lock()
	fmt.Printf("\r[%-25s]  %d%%   %d/%d", b.rate, b.percent, b.current, b.total)
	common.Glock.Unlock()
}

func (b *Bar) AddCurOne() {

	b.current = b.current + 1
	if b.current > b.total {
		b.current = b.total
	}
}

func (b *Bar) Finish() {
	fmt.Println()
}

func (b *Bar) GetTotal() uint32 {
	return b.total
}
