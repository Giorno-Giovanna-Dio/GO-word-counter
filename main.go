package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 檢查命令列參數
	if len(os.Args) < 2 {
		fmt.Println("使用方式: ./word-counter <檔案路徑>")
		fmt.Println("範例: ./word-counter example.txt")
		os.Exit(1)
	}

	// 取得檔案路徑
	filePath := os.Args[1]

	// 開啟檔案
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("錯誤：無法開啟檔案 '%s': %v\n", filePath, err)
		os.Exit(1)
	}
	defer file.Close()

	// 初始化計數器
	lineCount := 0
	wordCount := 0
	charCount := 0

	// 使用 Scanner 逐行讀取
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineCount++

		// 計算字元數（包含空白）
		charCount += len(line)

		// 計算單字數
		// 使用 Fields 分割空白字元來計算單字
		words := strings.Fields(line)
		wordCount += len(words)
	}

	// 檢查讀取過程中是否有錯誤
	if err := scanner.Err(); err != nil {
		fmt.Printf("錯誤：讀取檔案時發生問題: %v\n", err)
		os.Exit(1)
	}

	// 輸出結果
	fmt.Printf("\n檔案統計結果: %s\n", filePath)
	fmt.Println("─────────────────────────")
	fmt.Printf("行數：%d\n", lineCount)
	fmt.Printf("單字數：%d\n", wordCount)
	fmt.Printf("字元數：%d\n", charCount)
	fmt.Println("─────────────────────────")
}
