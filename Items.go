package alfred_models

type Item struct {
	Uid          string         `json:"uid"`
	Valid        bool           `json:"valid"`
	Title        string         `json:"title"`
	Subtitle     string         `json:"subtitle"`
	Arg          string         `json:"arg"`
	Autocomplete string         `json:"autocomplete"`
	Icon         Icon           `json:"icon"`
	Text         Text           `json:"text"`
	Mods         map[string]Mod `json:"mods"`
	Quicklook	string			`json:"quicklookurl"`
}

type Items struct {
	Items []Item `json:"items"`
}

func NewResponse() *Items {
	return new(Items)
}

func NewItem() *Item {
	return &Item{
		Uid:          "",
		Valid:        true,
		Title:        "",
		Subtitle:     "",
		Arg:          "",
		Autocomplete: "",
		Icon: Icon{
			Type: "filetype",
			Path: "icon.png",
		},
		Text: Text{
			Copy: "",
			Largetype: "",
		},
		Mods: make(map[string]Mod),
	}
}
