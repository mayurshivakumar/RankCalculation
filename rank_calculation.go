package rankcalculator

import (
	"RankCalculation/Godeps/_workspace/src/github.com/pmylund/sortutil"
	"errors"
)

type IRankCalculator interface {
	GetUsersInPosition(
		ranks []Rank,
		pos int,
	) ([]int, error)
	GetPosition(
		ranks []Rank,
		user_id int,
	) (int, error)
}

type RankCalculator struct {
	IRankCalculator
}

func GetRankCalculator() *RankCalculator {
	var rank_calculator = RankCalculator{}
	return &rank_calculator
}

// Given a list of ranks and position, returns the user id
func (rank_calculator *RankCalculator) GetUsersInPosition(ranks []Rank, pos int) ([]int, error) {
	var rank Rank
	var ids []int

	sortutil.DescByField(ranks, "Value")

	for i := 0; i < len(ranks); i++ {
		rank = ranks[i]
		user_pos, err := rank_calculator.GetPosition(ranks, rank.User_id)
		if err != nil {
			return ids, errors.New("Position does not exist in rankings provided")
		}
		if pos == user_pos {
			ids = append(ids, rank.User_id)
		}
	}

	if len(ids) > 0 {
		return ids, nil
	}

	return ids, errors.New("Position does not exist in rankings provided")
}

// Given a list of ranks, returns the users position in those ranks
func (rank_calculator *RankCalculator) GetPosition(ranks []Rank, user_id int) (int, error) {
	var rank Rank
	user_value := 0

	sortutil.DescByField(ranks, "Value")

	for i := 0; i < len(ranks); i++ {
		rank = ranks[i]
		if rank.User_id == user_id {
			user_value = rank.Value
		}
	}

	if user_value != 0 {
		for i := 0; i < len(ranks); i++ {
			rank = ranks[i]
			if rank.Value == user_value {
				return i + 1, nil
			}
		}
	}

	return 0, errors.New("User does not exist in rankings provided")
}
