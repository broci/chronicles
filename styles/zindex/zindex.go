package zindex

type Zindex struct {
	MobileStepper int
	Menu          int
	AppBar        int
	DrawerOverlay int
	NavDrawer     int
	DialogOverlay int
	Dialog        int
	Layer         int
	Popover       int
	Snackbar      int
	Tooltip       int
}

func New() Zindex {
	return Zindex{
		MobileStepper: 900,
		Menu:          1000,
		AppBar:        1100,
		DrawerOverlay: 1200,
		NavDrawer:     1300,
		DialogOverlay: 1400,
		Dialog:        1500,
		Layer:         2000,
		Popover:       2100,
		Snackbar:      2900,
		Tooltip:       3000,
	}
}
