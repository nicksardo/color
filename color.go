package color

import (
	"bytes"
	"io"
	"strconv"
	"strings"
)

type Color int
type Effect int

const (
	foreground       = 30
	background       = 40
	None       Color = 0
)

const (
	Black Color = iota + 1
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

const (
	BlackHi Color = iota + 60
	RedBri
	GreenHi
	YellowHi
	BlueHi
	MagentaHi
	CyanHi
	WhiteHi
)

const (
	reset Effect = iota
	Bold
	Faint
	Italic
	Underline
	BlinkSlow
	BlinkRapid
	ReverseVideo
	Concealed
	CrossedOut
)

type Palette struct {
	Foreground Color
	Background Color
	Effects    []Effect
	b          *bytes.Buffer
	reset      bool
}

func (p *Palette) Clear() {
	p.Foreground = None
	p.Background = None
	p.Effects = []Effect{}
}

func (p *Palette) ApplyTo(w io.Writer) {
	b := bytes.Buffer{}

	strs := make([]string, 0, len(p.Effects)+2)
	if p.Foreground != None {
		i := strconv.Itoa(foreground + int(p.Foreground) - 1)
		strs = append(strs, i)
	}
	if p.Background != None {
		i := strconv.Itoa(background + int(p.Background) - 1)
		strs = append(strs, i)
	}
	for _, v := range p.Effects {
		i := strconv.Itoa(int(v))
		strs = append(strs, i)
	}

	if len(strs) > 0 {
		b.WriteString("\x1b[")
		s := strings.Join(strs, ";")
		b.WriteString(s + "m\x1b[K")
	} else {
		b.WriteString("\x1b[0m\x1b[K")
	}

	w.Write(b.Bytes())
}
