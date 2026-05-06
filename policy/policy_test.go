package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	tests := []struct {
		name         string
		signal       Signal
		wantScore    int
		wantDecision string
	}{
		{name: "case_1", signal: Signal{Demand: 78, Capacity: 82, Latency: 23, Risk: 18, Weight: 7}, wantScore: 96, wantDecision: "review"},
		{name: "case_2", signal: Signal{Demand: 70, Capacity: 86, Latency: 8, Risk: 13, Weight: 11}, wantScore: 179, wantDecision: "accept"},
		{name: "case_3", signal: Signal{Demand: 67, Capacity: 98, Latency: 19, Risk: 20, Weight: 12}, wantScore: 115, wantDecision: "review"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Score(tc.signal); got != tc.wantScore {
				t.Fatalf("score = %d, want %d", got, tc.wantScore)
			}
			if got := Classify(tc.signal); got != tc.wantDecision {
				t.Fatalf("decision = %s, want %s", got, tc.wantDecision)
			}
		})
	}
}
