package betcity

type BetOffer struct{
	ID			int64
	Name		string
	Outcomes 	[]Outcome
}

func NewBetOffer(id int64, name string) (*BetOffer, error) {
	return &BetOffer{ID: id, Name: name}, nil
}

func (b *BetOffer) AddOutcome(o *Outcome) {
	b.Outcomes = append(b.Outcomes, *o)
}

