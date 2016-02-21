package rankcalculator

import (
	"testing"
)

func TestGetPosition(t *testing.T) {
	rank_calculator := GetRankCalculator()
	user_id := 12

  rank1 := GenerateRank(
		11,
		1,
	)
	rank2 := GenerateRank(
		12,
		1,
	)

  charts := []Rank{}
	charts = append(
		charts,
		rank1,
		rank2,
	)

	pos, err := rank_calculator.GetPosition(charts, user_id)

	if pos != 1 {
		t.Error("Expected pos to be 1, got", pos)
	}
	if err != nil {
		t.Error("Expected error to be nil, got", err)
	}

  rank1 = GenerateRank(
		11,
		12,
	)
	rank2 = GenerateRank(
		10,
		12,
	)
	rank3 := GenerateRank(
		15,
		14,
	)
	rank4 := GenerateRank(
		12,
		8,
	)
	charts = []Rank{}
	charts = append(
		charts,
		rank1,
		rank2,
		rank3,
		rank4,
	)

	pos, err = rank_calculator.GetPosition(charts, user_id)

	if pos != 4 {
		t.Error("Expected pos to be 4, got", pos)
	}
	if err != nil {
		t.Error("Expected error to be nil, got", err)
	}


}
