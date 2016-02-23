package model

//Solution : model for solution item
type Solution struct {
	Calculs  Calculs `json:"calculs"`
	Resultat int     `json:"resultat"`
}

//Solutions : list of solutions
type Solutions []Solution
