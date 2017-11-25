package colors

// keys
const (
	C50             = "50"
	C100            = "100"
	C200            = "200"
	C300            = "300"
	C400            = "400"
	C500            = "500"
	C600            = "600"
	C700            = "700"
	C800            = "800"
	C900            = "900"
	A100            = "A100"
	A200            = "A200"
	A400            = "A400"
	A700            = "A700"
	ContrastDefault = "contrast-default"
)

// Color maps color palette.
type Color map[string]string

// Amber color pallete.
func Amber() Color {
	return Color{
		C50:             "#fff8e1",
		C100:            "#ffecb3",
		C200:            "#ffe082",
		C300:            "#ffd54f",
		C400:            "#ffca28",
		C500:            "#ffc107",
		C600:            "#ffb300",
		C700:            "#ffa000",
		C800:            "#ff8f00",
		C900:            "#ff6f00",
		A100:            "#ffe57f",
		A200:            "#ffd740",
		A400:            "#ffc400",
		A700:            "#ffab00",
		ContrastDefault: "dark",
	}
}

// Blue color pallete.
func Blue() Color {
	return Color{
		C50:             "#e3f2fd",
		C100:            "#bbdefb",
		C200:            "#90caf9",
		C300:            "#64b5f6",
		C400:            "#42a5f5",
		C500:            "#2196f3",
		C600:            "#1e88e5",
		C700:            "#1976d2",
		C800:            "#1565c0",
		C900:            "#0d47a1",
		A100:            "#82b1ff",
		A200:            "#448aff",
		A400:            "#2979ff",
		A700:            "#2962ff",
		ContrastDefault: "light",
	}
}

// BlueGray color pallete.
func BlueGray() Color {
	return Color{
		C50:             "#eceff1",
		C100:            "#cfd8dc",
		C200:            "#b0bec5",
		C300:            "#90a4ae",
		C400:            "#78909c",
		C500:            "#607d8b",
		C600:            "#546e7a",
		C700:            "#455a64",
		C800:            "#37474f",
		C900:            "#263238",
		A100:            "#cfd8dc",
		A200:            "#b0bec5",
		A400:            "#78909c",
		A700:            "#455a64",
		ContrastDefault: "light",
	}
}
