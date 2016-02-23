package model

//Tirage : model for tirage item
type Tirage struct {
	De1       int       `json:"de1"`
	De2       int       `json:"de2"`
	De3       int       `json:"de3"`
	De4       int       `json:"de4"`
	De5       int       `json:"de5"`
	Solutions Solutions `json:"solutions"`
}

//Tirages : list of Tirage
type Tirages []Tirage
