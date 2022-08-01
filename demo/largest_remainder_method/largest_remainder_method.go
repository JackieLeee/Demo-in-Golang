package main

import (
	"fmt"
	"math"
)

func main() {
	list := []int64{33, 33, 33}
	fmt.Println(GetPercentValue(list, 2))
}

// GetPercentValue list为数值列表，round为保留的小数位数，返回的是百分比数组
func GetPercentValue(list []int64, round int) []float64 {
	res := make([]float64, len(list))
	if len(list) == 0 {
		return res
	}
	var sum int64
	for _, v := range list {
		sum += v
	}
	if sum <= 0 {
		return res
	}
	digits := math.Pow(10, float64(round))
	votesPerQuota := make([]float64, 0, len(list))
	for i := range list {
		votesPerQuota = append(votesPerQuota, float64(list[i])/float64(sum)*digits*100)
	}

	targetSeats := digits * 100
	seats := make([]float64, 0, len(list))
	for i := range votesPerQuota {
		seats = append(seats, math.Floor(votesPerQuota[i]))
	}

	var currentSum float64
	for i := range seats {
		currentSum += seats[i]
	}

	remainder := make([]float64, 0, len(list))
	for i := range votesPerQuota {
		remainder = append(remainder, votesPerQuota[i]-seats[i])
	}

	for currentSum < targetSeats {
		var max float64
		var maxIdx int
		for i := range remainder {
			if remainder[i] > max {
				max = remainder[i]
				maxIdx = i
			}
		}
		seats[maxIdx]++
		remainder[maxIdx] = 0
		currentSum++
	}

	for i := range seats {
		res[i] = seats[i] / digits
	}
	return res
}
