package string_newline_algorithm

import (
	"math"
	"strings"

	"github.com/mattn/go-runewidth"
)

// wrapWords 换行算法
func wrapWords(words []string, spc, lim, pen int) [][]string {
	n := len(words)

	length := make([][]int, n)
	for i := 0; i < n; i++ {
		length[i] = make([]int, n)
		length[i][i] = runewidth.StringWidth(words[i])
		for j := i + 1; j < n; j++ {
			length[i][j] = length[i][j-1] + spc + runewidth.StringWidth(words[j])
		}
	}
	nbrk := make([]int, n)
	cost := make([]int, n)
	for i := range cost {
		cost[i] = math.MaxInt32
	}
	for i := n - 1; i >= 0; i-- {
		if length[i][n-1] <= lim {
			cost[i] = 0
			nbrk[i] = n
		} else {
			for j := i + 1; j < n; j++ {
				d := lim - length[i][j-1]
				c := d*d + cost[j]
				if length[i][j-1] > lim {
					c += pen // too-long lines get a worse penalty
				}
				if c < cost[i] {
					cost[i] = c
					nbrk[i] = j
				}
			}
		}
	}
	var lines [][]string
	i := 0
	for i < n {
		lines = append(lines, words[i:nbrk[i]])
		i = nbrk[i]
	}
	return lines
}

// splitLineAndKeepWord 将字符串按照指定长度拆分成多行，并保证单词完整（但如果单个单词超过指定长度，则无法保证单词完整）
func splitLineAndKeepWord(s string, lim int) (lines []string) {
	// 将字符串按照换行符拆分成多行
	words := strings.Split(strings.Replace(s, "\n", " ", -1), " ")
	var newWords []string
	for i := range words {
		w := words[i]
		// 将太长的单词拆分成若干个最大长度为lim的单词
		for len(w) > lim {
			newWords = append(newWords, w[:lim])
			w = w[lim:]
		}
		// 将拆分后的单词加入新的单词列表
		if len(w) > 0 {
			newWords = append(newWords, w)
		}
	}
	// 拆分后的单词重新组合成行
	for _, line := range wrapWords(newWords, 1, lim, 1e5) {
		lines = append(lines, strings.Join(line, " "))
	}
	return lines
}
