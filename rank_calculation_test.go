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

	pos, err = rank_calculator.GetPosition(charts, 122)

	if pos != 0 {
		t.Error("Expected pos to be 0, got", pos)
	}
	if err == nil {
		t.Error("Expected error to be, got User does not exist in rankings provided")
	}
}

func TestGetUsersInPosition(t *testing.T) {
	rank_calculator := GetRankCalculator()
	var err error
	var ids []int
	charts := []Rank{}

	rank1 := GenerateRank(
		11,
		2,
	)
	rank2 := GenerateRank(
		12,
		2,
	)
	charts = append(
		charts,
		rank1,
		rank2,
	)

	ids, err = rank_calculator.GetUsersInPosition(charts, 1)

	if len(ids) != 2 {
		t.Error("Expected ids to be [11,12], got", ids)
	}
	if len(ids) == 0 {
		t.Error("Expected ids to be [11,12], got", ids)
	}
	if err != nil {
		t.Error("Expected err to be nil, got", ids)
	}

	rank1 = GenerateRank(
		11,
		1,
	)
	rank2 = GenerateRank(
		12,
		1,
	)
	charts = []Rank{}
	charts = append(
		charts,
		rank1,
		rank2,
	)

	ids, err = rank_calculator.GetUsersInPosition(charts, 2)

	if len(ids) != 0 {
		t.Error("Expected ids to be [ ], got", ids)
	}
	if err == nil {
		t.Error("err to be Position does not exist in rankings provided got", err)
	}

	rank1 = GenerateRank(
		11,
		10,
	)
	rank2 = GenerateRank(
		12,
		9,
	)
	rank3 := GenerateRank(
		13,
		9,
	)
	rank4 := GenerateRank(
		14,
		6,
	)
	rank5 := GenerateRank(
		15,
		5,
	)
	rank6 := GenerateRank(
		16,
		1,
	)
	charts = []Rank{}
	charts = append(
		charts,
		rank1,
		rank2,
		rank3,
		rank4,
		rank5,
		rank6,
	)

	ids, err = rank_calculator.GetUsersInPosition(charts, 4)

	if len(ids) != 1 {
		t.Error("Expected ids to be [14], got", ids)
	}
	if ids[0] != 14 {
		t.Error("Expected ids to be [14], got", ids)
	}
	if err != nil {
		t.Error("Expected err to be nil, got", ids)
	}
}
