package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/hello", hellohandler)
	http.HandleFunc("/now", nowhandler)
	http.HandleFunc("/webfortune", fortunehandler)

	http.ListenAndServe(":8080", nil)
}
func hellohandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "こんにちは from Cocespace !")
}
func nowhandler(w http.ResponseWriter, r *http.Request) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	fmt.Fprintln(w, (time.Now().In(jst)).Format("2006年01月02日 15:04:05"))
}

func fortunehandler(w http.ResponseWriter, r *http.Request) {
	seed := time.Now().UnixNano()
	d := rand.New(rand.NewSource(seed))
	var n int32 = d.Int31n(5)
	var fortune string
	if n == 0 {
		fortune = "大吉"
	} else if n == 1 {
		fortune = "吉"
	} else if n == 2 {
		fortune = "中吉"
	} else if n == 3 {
		fortune = "小吉"
	} else if n == 4 {
		fortune = "凶"
	}
	fmt.Fprint(w, "あなたの今日の運勢は",fortune,"です")
}
