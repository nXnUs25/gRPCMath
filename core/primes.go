package core

/*
k = 2
N = 210
while N > 1:
    if N % k == 0:   // if k evenly divides into N
        print k      // this is a factor
        N = N / k    // divide N by k so that we have the rest of the number left.
    else:
        k = k + 1
*/

func (p *Maths) PrimariesAI(number int) (results []int) {
	k := 2
	N := number

	for N > 1 {
		if N%k == 0 {
			results = append(results, k)
			N = N / k
		} else {
			k++
		}
	}
	return
}
