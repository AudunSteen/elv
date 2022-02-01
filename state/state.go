package state

func ViewState() string {
	// Sjekke data som er lagret ("kylling til venstre", "rev til venstre")
	return "[kylling rev korn hs ---\\_________\\/ _______________/---]"
}
func ViewState() string {
	// kylling rev korn og mann hopper i båten")
	return "[V----------\\kylling rev korn hs\\/ ----------//---Ø]"
}
func ViewState() string {
	// kylling rev korn og mann går opp på østsiden av elven")
	return "[V----------\\___________\\/ kylling rev korn hs//---Ø]"
}
3

func PutInBoat() string {
	return "rev"
}

func CrossRiver() string {
	return "test"
}
