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

// WrapTextToLines 将文本按照指定长度进行换行
func WrapTextToLines(text string, maxLength int) (wrappedLines []string) {
        if text == " " {
		return []string{" "}
	}
	var (
		currentLine      strings.Builder
		currentLineWidth int
	)
	for _, c := range text {
		// 处理换行符
		if c == '\n' {
			if currentLine.Len() > 0 {
				wrappedLines = append(wrappedLines, currentLine.String())
				currentLine.Reset()
				currentLineWidth = 0
			}
			continue
		}
		// 处理空格
		if unicode.IsSpace(c) {
			if currentLine.Len() > 0 {
				currentLine.WriteRune(' ')
			}
			continue
		}
		// 计算当前字符宽度
		charWidth := CountRuneWidth(c)
		// 如果添加当前字符会超出最大长度，先保存当前行
		if currentLineWidth+charWidth > maxLength {
			if currentLine.Len() > 0 {
				wrappedLines = append(wrappedLines, currentLine.String())
				currentLine.Reset()
				currentLineWidth = 0
			}
		}
		// 写入当前字符
		currentLine.WriteRune(c)
		currentLineWidth += charWidth
	}
	// 处理最后一行
	if currentLine.Len() > 0 {
		wrappedLines = append(wrappedLines, currentLine.String())
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
