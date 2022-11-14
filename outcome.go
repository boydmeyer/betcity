package betcity

type Outcome struct {
	ID		int64
	Label	string
	Odds	int64
}

func NewOutcome(id int64, label string, odds int64) (*Outcome, error) {
	return &Outcome{ID: id, Label: label, Odds: odds}, nil
}