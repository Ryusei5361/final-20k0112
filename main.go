package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
)

//テンプレートの生成
var primenumber = template.Must(template.New("index").Parse(`<!DOCTYPE html>
<html>
	<body>
	<h1>素数判定</h1>
		<form action="/">
			<label for="num">判定したい数</input>
			<input type="number" name="num" min="1" value="1">
			<input type="submit" value="判定">
		</form>
		<div>
		<h2><判定結果></h2>
		<li>{{.}}</li>
	</body>
</html>`))


type prime struct{
	number int
	judge  string
}

//メインルーチン
func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

//ハンドラの登録
func run() error {
	http.HandleFunc("/", handler)
	return http.ListenAndServe(":8080", nil)
}

//ハンドラ
func handler(w http.ResponseWriter, r *http.Request) {
	//入力した自然数を取得し、int型に変換
	num, err := strconv.Atoi(r.FormValue("num"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//テンプレートに埋め込む
	if err := primenumber.Execute(w, inputN(num)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//結果として表示される文字列を作成
func (p *prime) String() string {
	number := strconv.Itoa(p.number)
	judge  := p.judge
	return number + "　:　" + judge
}

//素数判定の関数をもとに出力する文を決定
func inputN(n int) *prime {
	switch Prime(n) {
	case true:
		return &prime{number: n, judge: "素数です。"}
	default:
		return &prime{number: n, judge: "素数ではありません。"}
	}
}

//素数を判定する関数
func Prime(n int) bool {
	check := true
	switch n {
	case 1:
		return false
	case 2:
		return true
	default:
		for i := 2; i < n; i++ {
			if n % i == 0 {
				check = false
				break
			}
		}
	}
	return check
}

