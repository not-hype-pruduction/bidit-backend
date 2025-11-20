package cards

import (
	"math/rand"
)

type generatedCards struct {
	A, K, Q, J     int8
	A_, K_, Q_, J_ int8
}

func generateCardsWithPoints(
	myPointsMin int32,
	myPointsMax int32,
	partnerPointsMin int32,
	partnerPointsMax int32,
) (*generatedCards, error) {

	aRange := randomRange(0, 4)
	kRange := randomRange(0, 4)
	qRange := randomRange(0, 4)
	jRange := randomRange(0, 4)

	a_Range := randomRange(0, 4)
	k_Range := randomRange(0, 4)
	q_Range := randomRange(0, 4)
	j_Range := randomRange(0, 4)

	for _, a := range aRange {
		for _, k := range kRange {
			for _, q := range qRange {
				for _, j := range jRange {
					if a+k+q+j > 13 {
						continue
					}

					myPoints := int32(a*4 + k*3 + q*2 + j*1)
					if myPoints < myPointsMin || myPoints > myPointsMax {
						continue
					}

					for _, a_ := range a_Range {
						if a_ > 4-a {
							continue
						}
						for _, k_ := range k_Range {
							if k_ > 4-k {
								continue
							}
							for _, q_ := range q_Range {
								if q_ > 4-q {
									continue
								}
								for _, j_ := range j_Range {
									if j_ > 4-j {
										continue
									}

									if a_+k_+q_+j_ > 13 {
										continue
									}

									partnerPoints := int32(a_*4 + k_*3 + q_*2 + j_*1)
									if partnerPoints < partnerPointsMin || partnerPoints > partnerPointsMax {
										continue
									}

									return &generatedCards{
										A: a, K: k, Q: q, J: j,
										A_: a_, K_: k_, Q_: q_, J_: j_,
									}, nil
								}
							}
						}
					}
				}
			}
		}
	}

	return nil, IMPOSSIBLECARDCOMBINATION
}

// randomRange создает случайную перестановку чисел от min до max включительно
func randomRange(min, max int8) []int8 {
	size := max - min + 1
	result := make([]int8, size)

	for i := int8(0); i < size; i++ {
		result[i] = min + i
	}

	// Тасуем массив (Fisher-Yates shuffle)
	for i := size - 1; i > 0; i-- {
		j := int8(rand.Intn(int(i + 1)))
		result[i], result[j] = result[j], result[i]
	}

	return result
}
