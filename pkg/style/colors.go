package style

import "github.com/muesli/termenv"

const (
	// https://github.com/muesli/duf/blob/master/themes.go#L29
	ColorBlue    = "#71BEF2"
	ColorMagenta = "#D290E4"
	ColorRed     = "#E88388"
)

func String(s string, color string) termenv.Style {
	p := termenv.ColorProfile()

	return termenv.String(s).Foreground(p.Color(color))
}

func Blue(s string) termenv.Style {
	return String(s, ColorBlue)
}

func Red(s string) termenv.Style {
	return String(s, ColorRed)
}

func Magenta(s string) termenv.Style {
	return String(s, ColorMagenta)
}
