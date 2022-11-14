package betcity

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sort"
)

type Client struct {}

// Creates a new Client struct
func New() (*Client, error) {
	return &Client{}, nil
}

// Returns a splice of competitions
func (c *Client) GetCompetitions() ([]Competition, error) {
	endpoint := "https://eu-offering.kambicdn.org/offering/v2018/betcitynl/listView/formula_1/race/all/all/competitions.json?lang=nl_NL&market=NL"
	req, err := http.Get(endpoint)
    if err != nil {
		return []Competition{}, err
    }

	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return []Competition{}, err
	}

	var raw RawCompetitionResult
	json.Unmarshal(data, &raw)

	results := []Competition{}
	for _, e := range raw.Events {
		c, _ := NewCompetition(e.Event.ID, e.Event.EnglishName, e.Event.Group, e.Event.GroupID, e.Event.State)
		for _, offer := range e.BetOffers {
			bo, _ := NewBetOffer(offer.ID, offer.Criterion.Label)
			for _, outcome := range offer.Outcomes {
				o, _ := NewOutcome(outcome.ID, outcome.Label, outcome.Odds)
				bo.AddOutcome(o)
			}
			c.AddBetOffer(bo)
		}
		for _, offer := range e.BetOffers {
			sort.Slice(offer.Outcomes, func(p, q int) bool { return offer.Outcomes[p].Odds < offer.Outcomes[q].Odds })
		}
		results = append(results, *c)
	}
	return results, nil
}