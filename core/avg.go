package core

func (a *Maths) AVG(nums ...int64) float64 {
	var sum int64
	for _, num := range nums {
		sum += num
	}
	return float64(sum) / float64(len(nums))
}
