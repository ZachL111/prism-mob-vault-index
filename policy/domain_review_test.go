package policy

import "testing"

func TestDomainReviewLane(t *testing.T) {
	item := DomainReview{Signal: 44, Slack: 52, Drag: 23, Confidence: 53}
	if got := DomainReviewScore(item); got != 124 {
		t.Fatalf("score = %d", got)
	}
	if got := DomainReviewLane(item); got != "watch" {
		t.Fatalf("lane = %s", got)
	}
}
