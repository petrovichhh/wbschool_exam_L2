package main

var (
	start      = 10
	end        = 100
	strategies = []Strategy{
		&PublicTransportStrategy{},
		&RoadStrategy{},
		&WalkStrategy{},
	}
)

func main() {
	nav := &Navigator{}
	for _, strategy := range strategies {
		nav.SetStrategy(strategy)
		nav.Route(start, end)
	}
}
