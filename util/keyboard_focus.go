package util

import (
	"sync"

	"github.com/gernest/keycode"
	"honnef.co/go/js/dom"
)

var state = &sync.Map{}

var focusKeys = map[string]bool{
	"tab":   true,
	"enter": true,
	"space": true,
	"esc":   true,
	"up":    true,
	"down":  true,
	"left":  true,
	"right": true,
}

func isFocusKey(e dom.KeyboardEvent) bool {
	code := e.KeyCode
	if code == 0 {
		code = e.CharCode
	}
	return focusKeys[keycode.String(code)]
}

// ListenForFocusKeys lists for focus keys.
func ListenForFocusKeys(w dom.Window) {
	if l, ok := state.Load("listening"); ok {
		t := l.(bool)
		if !t {
			w.AddEventListener("keyup", false, func(e dom.Event) {
				t := e.(dom.KeyboardEvent)
				if isFocusKey(t) {
					state.Store("focusKeyPressed", true)
				}
			})
			state.Store("listening", true)
		}
	}
}

// FocusKeyPressed returns true when the focus key is pressed.
func FocusKeyPressed() bool {
	if v, ok := state.Load("focusKeyPressed"); ok {
		return v.(bool)
	}
	return false
}
