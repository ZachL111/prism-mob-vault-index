package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	signal := Signal{Demand: 78, Capacity: 82, Latency: 23, Risk: 18, Weight: 7}
	if got := Score(signal); got != 96 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 70, Capacity: 86, Latency: 8, Risk: 13, Weight: 11}
	if got := Score(signal); got != 179 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "accept" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 67, Capacity: 98, Latency: 19, Risk: 20, Weight: 12}
	if got := Score(signal); got != 115 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
}
