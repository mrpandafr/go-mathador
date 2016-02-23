package model

//Calcul : model for calcul item
type Calcul struct {
	Operande1 int    `json:"operande1"`
	Operande2 int    `json:"operande2"`
	Resultat  int    `json:"resultat"`
	Signe     string `json:"signe"`
}

//Calculs : list of calculs
type Calculs []Calcul
