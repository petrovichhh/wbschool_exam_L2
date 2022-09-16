package main

import "fmt"

type Strategy interface {
	Route(startPoint int, endPoint int)
}

type Navigator struct {
	Strategy
}

func (nav *Navigator) SetStrategy(str Strategy) {
	nav.Strategy = str
}

type RoadStrategy struct {
}

func (r *RoadStrategy) Route(startPoint int, endPoint int) {
	avgSpeed := 30
	trafficJam := 2
	total := endPoint - startPoint
	totalTime := total * trafficJam
	fmt.Printf("Road A: [%d] to B: [%d] Avg speed: [%d] TrafficJam: [%d] Total: [%d] TotalTime: [%d] min\n",
		startPoint, endPoint, avgSpeed, trafficJam, total, totalTime)

}

type PublicTransportStrategy struct {
}

func (p *PublicTransportStrategy) Route(startPoint int, endPoint int) {
	avgSpeed := 40
	total := endPoint - startPoint
	totalTime := total * 40
	fmt.Printf("PublicTransport A: [%d] to B: [%d] Avg speed: [%d] Total: [%d] TotalTime: [%d] min\n",
		startPoint, endPoint, avgSpeed, total, totalTime)

}

type WalkStrategy struct {
}

func (w *WalkStrategy) Route(startPoint int, endPoint int) {
	avgSpeed := 4
	total := endPoint - startPoint
	totalTime := total * 60
	fmt.Printf("Walk A: [%d] to B: [%d] Avg speed: [%d] Total: [%d] TotalTime: [%d] min\n",
		startPoint, endPoint, avgSpeed, total, totalTime)

}
