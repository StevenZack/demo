package redis

type SortedSet struct {
}

func (s *SortedSet) ZAdd(score float64, value string) int {
	return 1
}

func (s *SortedSet) ZRange(from,to float64,withScores bool)  {
	
}
