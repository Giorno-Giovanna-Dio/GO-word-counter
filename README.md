# 主要特性：
- 檔案讀取：使用 os.Open 開啟檔案
- 逐行掃描：使用 bufio.Scanner 高效讀取大型檔案

# 三種統計：
- 行數：每讀取一行就累加
- 單字數：使用 strings.Fields 分割空白字元來計算
- 字元數：計算每行的字元長度
- 錯誤處理：檢查檔案開啟和讀取錯誤
- 清晰輸出：格式化的統計結果顯示
# 使用方式：
- 首先編譯程式：
`cd word-counter`
`go build -o word-counter main.go`
- 然後執行：
`./word-counter <檔案路徑>`

