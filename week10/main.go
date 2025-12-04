package main

import (
	"fmt"
	"html"
	"net/http"
	"os"
	"runtime"
)

const saveFile1 = "public/memo1.txt" // データファイルの保存先
const saveFile2 = "public/memo2.txt"
const saveFile3 = "public/memo3.txt"


func main() {
	fmt.Printf("Go version: %s\n", runtime.Version())

	http.HandleFunc("/hello", hellohandler)
	http.HandleFunc("/memo", memo)
	http.HandleFunc("/mwrite", mwrite)

	fmt.Println("Launch server...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to launch server: %v", err)
	}
}

func hellohandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "こんにちは from Codespace !")
}

func memo(w http.ResponseWriter, r *http.Request) {
	// データファイルを開く
	text, err := os.ReadFile(saveFile1)
	if err != nil {
		text = []byte("ここにメモを記入してください。")
	}
	text1, err := os.ReadFile(saveFile1)
	if err != nil {
		text1 = []byte("未記入")
	}
	text2, err := os.ReadFile(saveFile2)
	if err != nil {
		text2 = []byte("未記入")
	}
	text3, err := os.ReadFile(saveFile3)
	if err != nil {
		text3 = []byte("未記入")
	}
	// HTMLのフォームを返す
	htmlText := html.EscapeString(string(text))
	Text1 := html.EscapeString(string(text1))
	Text2 := html.EscapeString(string(text2))
	Text3 := html.EscapeString(string(text3))
	s := "<html>" +
	"<style>textarea { width:99%; height:200px; }</style>" +
	"<form method='get' action='/mwrite'>" +
	"<textarea name='text'>" + htmlText + "</textarea>" +
	"保存場所："+"<br>"+
    "<input type='radio' name='wmemo' value='+' checked> memo1" +"<br>"+Text1+"<br>"+
    "<input type='radio' name='wmemo' value='-' > memo2" +"<br>"+Text2+"<br>"+
    "<input type='radio' name='wmemo' value='*' > memo3"+"<br>"+Text3+"<br>"+
	"<input type='submit' value='保存' /></form></html>"
	w.Write([]byte(s))
r.ParseForm()
}

func mwrite(w http.ResponseWriter, r *http.Request) {
	// 投稿されたフォームを解析
	r.ParseForm()
	if len(r.Form["text"]) == 0 { // 値が書き込まれてない時
		w.Write([]byte("フォームから投稿してください。"))
		return
	}
	text := r.Form["text"][0]
	// データファイルへ書き込む
	switch r.Form["wmemo"][0]{
	case "+":
		os.WriteFile(saveFile1, []byte(text), 0644)
	case"-":
		os.WriteFile(saveFile2, []byte(text), 0644)
	case "*":
		os.WriteFile(saveFile3, []byte(text), 0644)
	}

	fmt.Println("save: " + text)
	// ルートページへリダイレクトして戻る --- (*4)
	http.Redirect(w, r, "/memo", 301)
}
