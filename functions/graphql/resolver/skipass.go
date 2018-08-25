package resolver

import (
	"github.com/blobor/skipass-api/functions/graphql/bapi"
	"github.com/graph-gophers/graphql-go"
	"net/http"
)

type SkipassResolver struct {
	skipass bapi.Skipass
}

func NewSkipass(args SkipassQueryArgs) (*SkipassResolver, error) {
	client := bapi.NewClient(http.DefaultClient)

	skipass, err := client.Skipass(*args.TicketNumber)

	if err != nil {
		return nil, err
	}

	return &SkipassResolver{skipass}, nil
}

func (s *SkipassResolver) Plan() string {
	return s.skipass.Plan
}

func (s *SkipassResolver) CardNumber() string {
	return s.skipass.CardNumber
}

func (s *SkipassResolver) TicketNumber() string {
	return s.skipass.TicketNumber
}

func (s *SkipassResolver) PurchaseDate() graphql.Time {
	return graphql.Time{Time: s.skipass.PurchaseDate}
}
