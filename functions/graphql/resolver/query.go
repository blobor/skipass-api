package resolver

import "context"

type QueryResolver struct{}

type SkipassQueryArgs struct {
	TicketNumber *string
}

func (_ *QueryResolver) Hello() string {
	return "Hello, world!"
}

func (_ *QueryResolver) Skipass(_ context.Context, args SkipassQueryArgs) (*SkipassResolver, error) {
	return NewSkipass(args)
}
