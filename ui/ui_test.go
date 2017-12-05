package ui

import "testing"
import "fmt"

func TestUI_Parse(t *testing.T) {
	src := `<div>
	<todo key="value">
{{if .}} . {{end}}
</todo>
	</div>

`
	u := &UI{}
	err := u.Parse([]byte(src))
	if err != nil {
		t.Fatal(err)
	}
	h, err := u.HTML()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(h)
}
