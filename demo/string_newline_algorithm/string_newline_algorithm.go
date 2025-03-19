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

    // 文本特征分析
    textLen := len(text)

    // 短文本（小于200字符）或包含大量换行符的文本使用Alpha版本
    if textLen < 200 || (strings.Count(text, "\n") > textLen/30 && textLen < 1000) {
        return WrapTextToLinesSimple(text, maxLength)
    }

    // 其他情况使用Beta版本
    return WrapTextToLinesCharByChar(text, maxLength)
}


// WrapTextToLinesSimple 将文本按照指定的最大长度进行换行
// 此算法采用简化处理逻辑，适合短文本和带换行符的文本
func WrapTextToLinesSimple(text string, maxLen int) []string {
    if text == "" || maxLen <= 0 {
        return []string{}
    }

    var result []string

    // 先按换行符分割文本
    paragraphs := strings.Split(text, "\n")

    for _, paragraph := range paragraphs {
        // 如果段落为空，添加一个空行
        if paragraph == "" {
            result = append(result, "")
            continue
        }

        var currentLine string
        var currentLineWidth int

        words := strings.Split(paragraph, " ")

        for _, word := range words {
            // 计算单词的宽度
            wordWidth := 0
            for _, r := range word {
                wordWidth += CountRuneWidth(r)
            }

            // 处理超长单词
            if wordWidth > maxLen {
                // 如果当前行不为空，先添加当前行
                if currentLine != "" {
                    result = append(result, currentLine)
                    currentLine = ""
                    currentLineWidth = 0
                }

                // 逐字符处理超长单词
                var longWordLine string
                var longWordLineWidth int

                for _, r := range word {
                    charWidth := CountRuneWidth(r)

                    // 如果添加当前字符会超过最大长度，添加当前行并开始新行
                    if longWordLineWidth+charWidth > maxLen {
                        result = append(result, longWordLine)
                        longWordLine = string(r)
                        longWordLineWidth = charWidth
                    } else {
                        longWordLine += string(r)
                        longWordLineWidth += charWidth
                    }
                }

                // 添加超长单词的最后一部分
                if longWordLine != "" {
                    currentLine = longWordLine
                    currentLineWidth = longWordLineWidth
                }
            } else {
                // 正常处理不超长的单词
                if currentLine == "" {
                    currentLine = word
                    currentLineWidth = wordWidth
                } else {
                    // 计算添加空格和单词后的总宽度
                    spaceWidth := 1 // 空格的宽度为1
                    totalWidth := currentLineWidth + spaceWidth + wordWidth

                    if totalWidth <= maxLen {
                        // 如果添加空格和单词后不超过最大长度，直接添加
                        currentLine += " " + word
                        currentLineWidth = totalWidth
                    } else {
                        // 如果超过最大长度，将当前行添加到结果，并开始新行
                        result = append(result, currentLine)
                        currentLine = word
                        currentLineWidth = wordWidth
                    }
                }
            }
        }

        // 添加段落的最后一行
        if currentLine != "" {
            result = append(result, currentLine)
        }
    }

    return result
}

// WrapTextToLinesCharByChar 将文本按照指定的最大长度进行换行，保持单词的完整性
// 此算法逐字符处理，适合处理长文本、中文文本和混合文本
func WrapTextToLinesCharByChar(text string, maxLength int) []string {
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
