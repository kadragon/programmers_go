// https://programmers.co.kr/learn/courses/30/lessons/42583

package main

func p42583(bridgeLength int, weight int, truckWeights []int) int {
	type truck struct {
		end    int
		weight int
	}

	var inPlace []truck
	time := 0
	sumW := 0

	for _, v := range truckWeights {
		if sumW+v > weight {
			for sumW+v > weight {
				if inPlace[0].end > time {
					time = inPlace[0].end
				}
				sumW -= inPlace[0].weight
				inPlace = inPlace[1:]
			}
		}

		inPlace = append(inPlace, truck{time + bridgeLength, v})
		sumW += v
		time++
	}

	return inPlace[len(inPlace)-1].end + 1
}
