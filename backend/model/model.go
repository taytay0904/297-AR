package model

type Message struct {
	Message string `json:"message" example:"message"`
}

type ModelOutput struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	ModelPath     string `json:"modelPath"`
	DetailPath    string `json:"detailPath"`
	SoundPath     string `json:"soundPath"`
	AnimationPath string `json:"animationPath"`
}

type Model struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	ModelPath     string `json:"modelPath"`
	DetailPath    string `json:"detailPath"`
	SoundPath     string `json:"soundPath"`
	AnimationPath string `json:"animationPath"`
}
