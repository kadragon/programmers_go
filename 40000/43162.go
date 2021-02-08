// https://programmers.co.kr/learn/courses/30/lessons/43162

package main

// Queue Data Struct
type Queue struct {
	Data  []int
	Start int
	End   int
}

func (q *Queue) inq(k int) {
	q.Data = append(q.Data, k)
	q.End++
}

func (q *Queue) deq() int {
	if q.Start < q.End {
		r := q.Data[q.Start]
		q.Start++
		return r
	}
	return -1
}

func p43162(n int, computers [][]int) int {
	var queue Queue
	check := make([]int, n)
	result := 0

	for i := 0; i < n; i++ {
		if check[i] == 0 {
			result++
			queue.inq(i)
			for queue.Start < queue.End {
				t := queue.deq()
				check[t] = 1

				for j := 0; j < n; j++ {
					if computers[t][j] == 1 && check[j] == 0 {
						queue.inq(j)
					}
				}
			}
		}
	}
	return result
}
