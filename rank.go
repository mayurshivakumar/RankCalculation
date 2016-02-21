package rankcalculator

type Rank struct {
	User_id int `json:"user_id"`
	Value   int `json:"data_value"`
}

// Takes user_id and value and returns a Rank struct
func GenerateRank(user_id, value int) Rank {
	var rank Rank
	rank.User_id = user_id
	rank.Value = value
	return rank
}
