package style

import "github.com/muesli/termenv"

const (
	// https://github.com/muesli/duf/blob/master/themes.go#L29
	ColorBlue    = "#71BEF2"
	ColorCyan    = "#66C2CD"
	ColorGray    = "#B9BFCA"
	ColorGreen   = "#A8CC8C"
	ColorMagenta = "#D290E4"
	ColorRed     = "#E88388"
	ColorYellow  = "#DBAB79"
)

func String(s string, color string) termenv.Style {
	p := termenv.ColorProfile()

	return termenv.String(s).Foreground(p.Color(color))
}

func Blue(s string) termenv.Style {
	return String(s, ColorBlue)
}

func Cyan(s string) termenv.Style {
	return String(s, ColorCyan)
}

func Gray(s string) termenv.Style {
	return String(s, ColorGray)
}

func Green(s string) termenv.Style {
	return String(s, ColorGreen)
}

func Red(s string) termenv.Style {
	return String(s, ColorRed)
}

func Magenta(s string) termenv.Style {
	return String(s, ColorMagenta)
}

func Yellow(s string) termenv.Style {
	return String(s, ColorYellow)
}
