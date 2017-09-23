package alfred_models

const (
	Shift = "shift"
	Cmd   = "cmd"
	Fn    = "fn"
	Alt   = "alt"
	Ctrl  = "ctrl"
)

var Modifiers = [5]string{Shift, Cmd, Fn, Alt, Ctrl}

type Icon struct {
	Type string `json:"type"`
	Path string `json:"path"`
}

type Text struct {
	Copy      string `json:"copy"`
	Largetype string `json:"largetype"`
}

type Mod struct {
	Valid    bool   `json:"valid"`
	Arg      string `json:"arg"`
	Subtitle string `json:"subtitle"`
}

func NewMod(arg string, subtitle string) *Mod {
	if len(arg) == 0 && len(subtitle) == 0 {
		return &Mod{
			Valid: false,
			Arg: "",
			Subtitle: "",
		}
	}
	return &Mod{
		Valid:    true,
		Arg:      arg,
		Subtitle: subtitle,
	}
}