package largest_remainder_method

import (
	"math"
)

// GetPercentValue 最大余额法求百分比（所有百分比加起来必定为100%）
// list: 数据列表
// round: 四舍五入的取整位数
// setMin4NoZeroValue: 设置为true时，当某个非0值百分比为0时，将其设置为当前精度下的最小值
// 返回值为与list长度相同的数组，其内每个元素为对应list数组元素所占的百分比
func GetPercentValue(list []int64, round int, setMin4NoZeroVal bool) []float64 {
	res := make([]float64, len(list))
	if len(list) == 0 || round < 0 {
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

	var resSum float64
	for i := range seats {
		val := seats[i] / digits
		// 如果原来的占比为0，但数值不为0，则填充默认值
		if setMin4NoZeroVal && val == 0 && list[i] > 0 {
			val = math.Pow(10, float64(-round))
		}
		res[i] = val
		resSum += res[i]
	}
	// 如果结果不是100%且不算0%，则重新计算补充剩余的百分比
	if resSum != 100 && resSum != 0 {
		res = reGetPercentValue(res, round)
	}
	return res
}

func reGetPercentValue(list []float64, round int) []float64 {
	res := make([]float64, len(list))
	if len(list) == 0 {
		return res
	}
	var sum float64
	for _, v := range list {
		sum += v
	}
	if sum <= 0 {
		return res
	}
	digits := math.Pow(10, float64(round))
	votesPerQuota := make([]float64, 0, len(list))
	for i := range list {
		votesPerQuota = append(votesPerQuota, list[i]/sum*digits*100)
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
