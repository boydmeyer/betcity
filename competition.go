package betcity

import (
	"fmt"
	"strconv"

	"github.com/xlab/tablewriter"
)

type Competition struct{
	ID			int64
	Name		string
	Group		string
	GroupId		int64
	BetOffers	[]BetOffer
	State		string
}

func NewCompetition(id int64, name string, group string, groupId int64, state string) (*Competition, error) {
	return &Competition{
		ID: id,
		Name: name,
		Group: group,
		GroupId: groupId,
		State: state,
	}, nil
}

func (c *Competition) AddBetOffer(bo *BetOffer) {
	c.BetOffers = append(c.BetOffers, *bo)
}

func (c *Competition) Render() {
	for _, offer := range c.BetOffers {
		table := tablewriter.CreateTable()
		table.AddHeaders(offer.Name, "Name", "Odds")
		for i, o := range offer.Outcomes {
			table.AddRow(strconv.Itoa(i+1), o.Label, o.Odds)
		}
		fmt.Println(table.Render())
	}

}

type RawCompetitionResult struct {
	Events           []struct {
		BetOffers []struct {
			ID        int64     `json:"id"`
			Criterion struct {
				ID           int    `json:"id"`
				Label        string `json:"label"`
				EnglishLabel string `json:"englishLabel"`
				Order        []int  `json:"order"`
			} `json:"criterion"`
			Outcomes []struct {
				ID             int64   `json:"id"`
				Label          string  `json:"label"`
				Odds           int64   `json:"odds"`
			} `json:"outcomes"`
		} `json:"betOffers"`
		Event struct {
			EnglishName    string `json:"englishName"`
			Group          string `json:"group"`
			GroupID        int64  `json:"groupId"`
			ID             int64  `json:"id"`
			Name           string `json:"name"`
			Sport string   `json:"sport"`
			Start string   `json:"start"`
			State string   `json:"state"`
		} `json:"event"`
	}
}
