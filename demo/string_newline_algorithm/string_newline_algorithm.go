package string_newline_algorithm

import (
	"strings"
	"unicode"
)

// WrapTextToLines 将文本按照指定的最大长度进行换行处理
// text: 需要处理的输入文本
// maxLength: 每行允许的最大长度（对于中文和全角字符按2个长度计算）
// 返回值: 按照指定长度换行后的文本行数组
func WrapTextToLines(text string, maxLength int) []string {
	// 特殊情况快速处理
	switch {
	case text == "":
		return []string{""}
	case text == " ":
		return []string{" "}
	case len(text) <= maxLength:
		if !strings.ContainsRune(text, '\n') {
			return []string{text}
		}
	}

	var (
		currentLine      strings.Builder // 当前正在构建的行
		currentWord      strings.Builder // 当前正在构建的单词
		currentLineWidth int             // 当前行的实际宽度（考虑中文全角）
		currentWordWidth int             // 当前单词的实际宽度
		spaceNeeded      bool            // 标记是否需要在添加下一个单词前加空格
	)

	// 预分配内存以提高性能
	estimatedLines := (len(text) / maxLength) + 1
	wrappedLines := make([]string, 0, estimatedLines)
	currentLine.Grow(maxLength)
	currentWord.Grow(maxLength)

	// 遍历文本中的每个字符
	for _, c := range text {
		// 处理换行符：保存当前单词和当前行，开始新的一行
		if c == '\n' {
			if currentWord.Len() > 0 {
				if spaceNeeded {
					currentLine.WriteByte(' ')
					currentLineWidth++
				}
				currentLine.WriteString(currentWord.String())
				currentLineWidth += currentWordWidth
				currentWord.Reset()
				currentWordWidth = 0
			}
			if currentLine.Len() > 0 {
				wrappedLines = append(wrappedLines, currentLine.String())
				currentLine.Reset()
				currentLineWidth = 0
				spaceNeeded = false
			}
			continue
		}

		// 处理空格：将已缓存的单词添加到当前行
		if unicode.IsSpace(c) {
			if currentWord.Len() > 0 {
				if currentLineWidth > 0 && currentLineWidth+1+currentWordWidth > maxLength {
					wrappedLines = append(wrappedLines, currentLine.String())
					currentLine.Reset()
					currentLineWidth = 0
					spaceNeeded = false
				}
				if spaceNeeded {
					currentLine.WriteByte(' ')
					currentLineWidth++
				}
				currentLine.WriteString(currentWord.String())
				currentLineWidth += currentWordWidth
				currentWord.Reset()
				currentWordWidth = 0
				spaceNeeded = true
			}
			continue
		}

		// 处理单词超长情况：直接将当前单词作为单独的一行
		charWidth := CountRuneWidth(c)
		if currentWordWidth+charWidth > maxLength {
			if currentLine.Len() > 0 {
				wrappedLines = append(wrappedLines, currentLine.String())
				currentLine.Reset()
				currentLineWidth = 0
				spaceNeeded = false
			}
			if currentWord.Len() > 0 {
				wrappedLines = append(wrappedLines, currentWord.String())
				currentWord.Reset()
				currentWordWidth = 0
			}
			currentWord.WriteRune(c)
			currentWordWidth = charWidth
			continue
		}

		// 将字符添加到当前单词缓存中
		currentWord.WriteRune(c)
		currentWordWidth += charWidth
	}

	// 处理最后一个单词
	if currentWord.Len() > 0 {
		if currentLineWidth > 0 && currentLineWidth+1+currentWordWidth > maxLength {
			wrappedLines = append(wrappedLines, currentLine.String())
			currentLine.Reset()
			currentLineWidth = 0
			spaceNeeded = false
		}
		if spaceNeeded {
			currentLine.WriteByte(' ')
		}
		currentLine.WriteString(currentWord.String())
	}

	// 处理最后一行
	if currentLine.Len() > 0 {
		wrappedLines = append(wrappedLines, currentLine.String())
	}

	return wrappedLines
}
