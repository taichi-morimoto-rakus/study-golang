package result

import (
	"io"
	"sort"
	"strconv"
	"sync"
)

type Line struct {
	Text string
	No   int
}

type Result struct {
	sync.Mutex
	Data map[string][]Line
}

// 検索結果はこのグローバル変数に保存される
var Store = &Result{Data: make(map[string][]Line, 100)}

// ファイル名、一致した行の内容、行番号を渡すと var Store に保存する関数
func Set(fileName, txt string, no int) {
	Store.Lock()
	defer Store.Unlock()

	if _, ok := Store.Data[fileName]; !ok {
		Store.Data[fileName] = make([]Line, 0, 10)
	}
	Store.Data[fileName] = append(Store.Data[fileName], Line{txt, no})
}

func RenderFiles(w io.Writer) {
	for _, file := range Store.Files() {
		io.WriteString(w, file+"\n")
	}
}

func RenderWithContent(w io.Writer) {
	for index, file := range Store.Files() {
		io.WriteString(w, file+"\n")
		for _, line := range Store.Data[file] {
			io.WriteString(w, strconv.Itoa(line.No)+": "+line.Text+"\n")
		}

		if index < len(Store.Data)-1 {
			io.WriteString(w, "\n")
		}
	}
}

// 保存されているファイル名を昇順でソートした上で []string として返す関数
func (r *Result) Files() []string {
	files := make([]string, 0, len(r.Data))
	for k := range r.Data {
		files = append(files, k)
	}

	sort.Strings(files)
	return files
}

func Reset() {
	Store = &Result{Data: make(map[string][]Line, 100)}
}
