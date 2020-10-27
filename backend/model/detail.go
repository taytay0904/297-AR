package model

type DetailInput struct {
	ID string `json:"id"`
}

type DetailOutput struct {
	Weight   string `json:"weight"`
	Height   string `json:"height"`
	Built    string `json:"built"`
	Designer string `json:"desinger"`
}
