package core

func (m *Maths) Max(max, i int32) int32 {
	if max < i {
		return i
	} else {
		return max
	}
}
