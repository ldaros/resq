package color

type Color int

const (
	Black Color = iota
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

func (c Color) String() string {
	return [...]string{"Black", "Red", "Green", "Yellow", "Blue", "Magenta", "Cyan", "White"}[c]
}

func (c Color) Code() string {
	return [...]string{"30", "31", "32", "33", "34", "35", "36", "37"}[c]
}

// Colorize a string with the given color
func Colorize(text string, color Color) string {
	return "\033[" + color.Code() + "m" + text + "\033[0m"
}

// Color for http status codes
func GetHTTPStatusColors(s int) Color {
	// Informational
	if s >= 100 && s < 200 {
		return Cyan
	}

	// Success
	if s >= 200 && s < 300 {
		return Green
	}

	// Redirection
	if s >= 300 && s < 400 {
		return Yellow
	}

	// Client error
	if s >= 400 && s < 500 {
		return Yellow
	}

	// Server error
	if s >= 500 && s < 600 {
		return Red
	}

	return White
}

func GetExecutionTimeColor(t_milis int64) Color {
	// under 5 second
	if t_milis < 1000 {
		return Green
	}

	// under 10 seconds
	if t_milis < 5000 {
		return Yellow
	}

	return Red
}
