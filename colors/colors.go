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
	Black           = "black"
	White           = "white"
	Transparent     = "transparent"
	FullBlack       = "full-black"
	DarkBlack       = "dark-black"
	LightBlack      = "light-black"
	MinBlack        = "min-black"
	FaintBlack      = "faint-black"
	FullWhite       = "full-white"
	DarkWhite       = "dark-white"
	LightWhite      = "light-white"
)

// Contrast can either be dark or light.
type Contrast string

// contrast options
const (
	LightContrast Contrast = "light"
	DarkContrast  Contrast = "dark"
)

func (c Contrast) Value() string {
	return string(c)
}

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

// Brown color pallete.
func Brown() Color {
	return Color{
		C50:             "#efebe9",
		C100:            "#d7ccc8",
		C200:            "#bcaaa4",
		C300:            "#a1887f",
		C400:            "#8d6e63",
		C500:            "#795548",
		C600:            "#6d4c41",
		C700:            "#5d4037",
		C800:            "#4e342e",
		C900:            "#3e2723",
		A100:            "#d7ccc8",
		A200:            "#bcaaa4",
		A400:            "#8d6e63",
		A700:            "#5d4037",
		ContrastDefault: "light",
	}
}

// Common color pallete.
func Common() CommonColors {
	return CommonColors{
		Black:       "#000",
		White:       "#fff",
		Transparent: "rgba(0, 0, 0, 0)",
		FullBlack:   "rgba(0, 0, 0, 1)",
		DarkBlack:   "rgba(0, 0, 0, 0.87)",
		LightBlack:  "rgba(0, 0, 0, 0.54)",
		MinBlack:    "rgba(0, 0, 0, 0.26)",
		FaintBlack:  "rgba(0, 0, 0, 0.12)",
		FullWhite:   "rgba(255, 255, 255, 1)",
		DarkWhite:   "rgba(255, 255, 255, 0.87)",
		LightWhite:  "rgba(255, 255, 255, 0.54)",
	}
}

// Cyan color pallete.
func Cyan() Color {
	return Color{
		C50:             "#e0f7fa",
		C100:            "#b2ebf2",
		C200:            "#80deea",
		C300:            "#4dd0e1",
		C400:            "#26c6da",
		C500:            "#00bcd4",
		C600:            "#00acc1",
		C700:            "#0097a7",
		C800:            "#00838f",
		C900:            "#006064",
		A100:            "#84ffff",
		A200:            "#18ffff",
		A400:            "#00e5ff",
		A700:            "#00b8d4",
		ContrastDefault: "dark",
	}
}

// DeepOrange color pallete.
func DeepOrange() Color {
	return Color{
		C50:             "#fbe9e7",
		C100:            "#ffccbc",
		C200:            "#ffab91",
		C300:            "#ff8a65",
		C400:            "#ff7043",
		C500:            "#ff5722",
		C600:            "#f4511e",
		C700:            "#e64a19",
		C800:            "#d84315",
		C900:            "#bf360c",
		A100:            "#ff9e80",
		A200:            "#ff6e40",
		A400:            "#ff3d00",
		A700:            "#dd2c00",
		ContrastDefault: "light",
	}
}

// DeepPurple color pallete.
func DeepPurple() Color {
	return Color{
		C50:             "#ede7f6",
		C100:            "#d1c4e9",
		C200:            "#b39ddb",
		C300:            "#9575cd",
		C400:            "#7e57c2",
		C500:            "#673ab7",
		C600:            "#5e35b1",
		C700:            "#512da8",
		C800:            "#4527a0",
		C900:            "#311b92",
		A100:            "#b388ff",
		A200:            "#7c4dff",
		A400:            "#651fff",
		A700:            "#6200ea",
		ContrastDefault: "light",
	}
}

// Green color pallete.
func Green() Color {
	return Color{
		C50:             "#e8f5e9",
		C100:            "#c8e6c9",
		C200:            "#a5d6a7",
		C300:            "#81c784",
		C400:            "#66bb6a",
		C500:            "#4caf50",
		C600:            "#43a047",
		C700:            "#388e3c",
		C800:            "#2e7d32",
		C900:            "#1b5e20",
		A100:            "#b9f6ca",
		A200:            "#69f0ae",
		A400:            "#00e676",
		A700:            "#00c853",
		ContrastDefault: "dark",
	}
}

// Grey color pallete.
func Grey() Color {
	return Color{
		C50:             "#fafafa",
		C100:            "#f5f5f5",
		C200:            "#eeeeee",
		C300:            "#e0e0e0",
		C400:            "#bdbdbd",
		C500:            "#9e9e9e",
		C600:            "#757575",
		C700:            "#616161",
		C800:            "#424242",
		C900:            "#212121",
		A100:            "#d5d5d5",
		A200:            "#aaaaaa",
		A400:            "#303030",
		A700:            "#616161",
		ContrastDefault: "dark",
	}
}

// Indigo color pallete.
func Indigo() Color {
	return Color{
		C50:             "#e8eaf6",
		C100:            "#c5cae9",
		C200:            "#9fa8da",
		C300:            "#7986cb",
		C400:            "#5c6bc0",
		C500:            "#3f51b5",
		C600:            "#3949ab",
		C700:            "#303f9f",
		C800:            "#283593",
		C900:            "#1a237e",
		A100:            "#8c9eff",
		A200:            "#536dfe",
		A400:            "#3d5afe",
		A700:            "#304ffe",
		ContrastDefault: "light",
	}
}

// LightBlue color pallete.
func LightBlue() Color {
	return Color{
		C50:             "#e1f5fe",
		C100:            "#b3e5fc",
		C200:            "#81d4fa",
		C300:            "#4fc3f7",
		C400:            "#29b6f6",
		C500:            "#03a9f4",
		C600:            "#039be5",
		C700:            "#0288d1",
		C800:            "#0277bd",
		C900:            "#01579b",
		A100:            "#80d8ff",
		A200:            "#40c4ff",
		A400:            "#00b0ff",
		A700:            "#0091ea",
		ContrastDefault: "dark",
	}
}

// LightGreen color pallete.
func LightGreen() Color {
	return Color{
		C50:             "#f1f8e9",
		C100:            "#dcedc8",
		C200:            "#c5e1a5",
		C300:            "#aed581",
		C400:            "#9ccc65",
		C500:            "#8bc34a",
		C600:            "#7cb342",
		C700:            "#689f38",
		C800:            "#558b2f",
		C900:            "#33691e",
		A100:            "#ccff90",
		A200:            "#b2ff59",
		A400:            "#76ff03",
		A700:            "#64dd17",
		ContrastDefault: "dark",
	}
}

// Lime color pallete.
func Lime() Color {
	return Color{
		C50:             "#f9fbe7",
		C100:            "#f0f4c3",
		C200:            "#e6ee9c",
		C300:            "#dce775",
		C400:            "#d4e157",
		C500:            "#cddc39",
		C600:            "#c0ca33",
		C700:            "#afb42b",
		C800:            "#9e9d24",
		C900:            "#827717",
		A100:            "#f4ff81",
		A200:            "#eeff41",
		A400:            "#c6ff00",
		A700:            "#aeea00",
		ContrastDefault: "dark",
	}
}

// Orange color pallete.
func Orange() Color {
	return Color{
		C50:             "#fff3e0",
		C100:            "#ffe0b2",
		C200:            "#ffcc80",
		C300:            "#ffb74d",
		C400:            "#ffa726",
		C500:            "#ff9800",
		C600:            "#fb8c00",
		C700:            "#f57c00",
		C800:            "#ef6c00",
		C900:            "#e65100",
		A100:            "#ffd180",
		A200:            "#ffab40",
		A400:            "#ff9100",
		A700:            "#ff6d00",
		ContrastDefault: "dark",
	}
}

// Pink color pallete.
func Pink() Color {
	return Color{
		C50:             "#fce4ec",
		C100:            "#f8bbd0",
		C200:            "#f48fb1",
		C300:            "#f06292",
		C400:            "#ec407a",
		C500:            "#e91e63",
		C600:            "#d81b60",
		C700:            "#c2185b",
		C800:            "#ad1457",
		C900:            "#880e4f",
		A100:            "#ff80ab",
		A200:            "#ff4081",
		A400:            "#f50057",
		A700:            "#c51162",
		ContrastDefault: "light",
	}
}

// Purple color pallete.
func Purple() Color {
	return Color{
		C50:             "#f3e5f5",
		C100:            "#e1bee7",
		C200:            "#ce93d8",
		C300:            "#ba68c8",
		C400:            "#ab47bc",
		C500:            "#9c27b0",
		C600:            "#8e24aa",
		C700:            "#7b1fa2",
		C800:            "#6a1b9a",
		C900:            "#4a148c",
		A100:            "#ea80fc",
		A200:            "#e040fb",
		A400:            "#d500f9",
		A700:            "#aa00ff",
		ContrastDefault: "light",
	}
}

// Red color pallete.
func Red() Color {
	return Color{
		C50:             "#ffebee",
		C100:            "#ffcdd2",
		C200:            "#ef9a9a",
		C300:            "#e57373",
		C400:            "#ef5350",
		C500:            "#f44336",
		C600:            "#e53935",
		C700:            "#d32f2f",
		C800:            "#c62828",
		C900:            "#b71c1c",
		A100:            "#ff8a80",
		A200:            "#ff5252",
		A400:            "#ff1744",
		A700:            "#d50000",
		ContrastDefault: "light",
	}
}

// Teal color pallete.
func Teal() Color {
	return Color{
		C50:             "#e0f2f1",
		C100:            "#b2dfdb",
		C200:            "#80cbc4",
		C300:            "#4db6ac",
		C400:            "#26a69a",
		C500:            "#009688",
		C600:            "#00897b",
		C700:            "#00796b",
		C800:            "#00695c",
		C900:            "#004d40",
		A100:            "#a7ffeb",
		A200:            "#64ffda",
		A400:            "#1de9b6",
		A700:            "#00bfa5",
		ContrastDefault: "light",
	}
}

// Yellow color pallete.
func Yellow() Color {
	return Color{
		C50:             "#fffde7",
		C100:            "#fff9c4",
		C200:            "#fff59d",
		C300:            "#fff176",
		C400:            "#ffee58",
		C500:            "#ffeb3b",
		C600:            "#fdd835",
		C700:            "#fbc02d",
		C800:            "#f9a825",
		C900:            "#f57f17",
		A100:            "#ffff8d",
		A200:            "#ffff00",
		A400:            "#ffea00",
		A700:            "#ffd600",
		ContrastDefault: "dark",
	}
}

// CommonColors common color pallete.
type CommonColors struct {
	Black       string
	White       string
	Transparent string
	FullBlack   string
	DarkBlack   string
	LightBlack  string
	MinBlack    string
	FaintBlack  string
	FullWhite   string
	DarkWhite   string
	LightWhite  string
}
