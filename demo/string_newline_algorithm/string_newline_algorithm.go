package string_newline_algorithm

import (
	"strings"
	"unicode"
)

func CountStringRuneWidth(s string) (res int) {
	for _, r := range s {
		res += CountRuneWidth(r)
	}
	return
}

// SplitStringIntoWords 将字符串按照空格分割成单词
func SplitStringIntoWords(s string) []string {
	words := make([]string, 0, len(s)/5)
	var wordBegin int
	wordPending := false
	for i, c := range s {
		if unicode.IsSpace(c) {
			if wordPending {
				words = append(words, s[wordBegin:i])
				wordPending = false
			}
			continue
		}
		if !wordPending {
			wordBegin = i
			wordPending = true
		}
	}
	if wordPending {
		words = append(words, s[wordBegin:])
	}
	return words
}

// WrapTextToLines 将文本按照指定长度进行换行
func WrapTextToLines(text string, maxLength int) (wrappedLines []string) {
	if text == " " {
		return []string{" "}
	}
	words := SplitStringIntoWords(text)
	if len(words) == 0 {
		return []string{""}
	}

	var currentLength int
	var currentWords []string
	for _, word := range words {
		wordWidth := CountStringRuneWidth(word)
		if wordWidth > maxLength {
			cutWords := SplitTextByLength(word, maxLength)
			for _, cutWord := range cutWords {
				cutWordWidth := CountStringRuneWidth(cutWord)
				if currentLength+cutWordWidth+1 > maxLength {
					wrappedLines = append(wrappedLines, strings.Join(currentWords, " "))
					currentLength = cutWordWidth
					currentWords = []string{cutWord}
				} else {
					currentLength += cutWordWidth + 1
					currentWords = append(currentWords, cutWord)
				}
			}
			continue
		}
		if currentLength+wordWidth+1 > maxLength {
			wrappedLines = append(wrappedLines, strings.Join(currentWords, " "))
			currentLength = wordWidth
			currentWords = []string{word}
		} else {
			currentLength += wordWidth + 1
			currentWords = append(currentWords, word)
		}
	}

	if currentLength > 0 {
		wrappedLines = append(wrappedLines, strings.Join(currentWords, " "))
	}
	return wrappedLines
}

// SplitTextByLength 将文本按照指定长度进行分割
func SplitTextByLength(text string, maxLength int) (lines []string) {
	var currentLine strings.Builder
	var currentLength int
	for _, r := range text {
		rWidth := CountRuneWidth(r)
		currentLength += rWidth
		if currentLength < maxLength {
			// 拼接该字符
			currentLine.WriteRune(r)
		} else if currentLength == maxLength {
			// 该字符正好填满一行
			currentLine.WriteRune(r)
			lines = append(lines, currentLine.String())
			// 重新开始一行
			currentLine.Reset()
			currentLength = 0
		} else {
			// 该字符超出一行
			lines = append(lines, currentLine.String())
			// 以该字符开始一行
			currentLine.Reset()
			currentLine.WriteRune(r)
			currentLength = rWidth
		}
	}
	// 处理最后一行
	if currentLength > 0 {
		lines = append(lines, currentLine.String())
	}
	return lines
}

func IsChineseRune(r rune) bool {
	return unicode.Is(unicode.Han, r)
}

func CountRuneWidth(r rune) int {
	if IsChineseRune(r) {
		return 2
	}
	return 1
}
