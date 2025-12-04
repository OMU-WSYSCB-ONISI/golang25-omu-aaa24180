package main

import (
	"fmt"
	"net/http"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Go version: %s\n", runtime.Version())

	http.Handle("/", http.FileServer(http.Dir("public/")))
	http.HandleFunc("/hello", hellohandler)
	http.HandleFunc("/enq", enqhandler)
	http.HandleFunc("/fdump", fdump)
	http.HandleFunc("/cal00", cal00handler)
	http.HandleFunc("/cal01", calpmhandler)
	http.HandleFunc("/sum", sumhandler)
	http.HandleFunc("/bmi", bmicalc)
	http.HandleFunc("/cal02", calpmhandler2)
	http.HandleFunc("/ave", avehandler)
	fmt.Println("Launch server...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to launch server: %v", err)
	}
}

func hellohandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "こんにちは from Codespace !")
}

func fdump(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println("errorだよ")
	}
	// フォームはマップとして利用でき以下で内容を確認できる．
	for k, v := range r.Form {
		fmt.Printf("%v : %v\n", k, v)
	}
}

func enqhandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println("errorだよ")
	}
	// r.FormValue("name")として，フォーム中name欄の値を得る
	fmt.Fprintln(w, r.FormValue("name")+"さん，ご協力ありがとうございます.\n年齢は"+r.FormValue("age")+"で，性別は"+r.FormValue("gend")+"で，出身地は"+r.FormValue("birthplace")+"ですね")
}

func cal00handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println("errorだよ")
	}
	price, _ := strconv.Atoi(r.FormValue("price"))
	num, _ := strconv.Atoi(r.FormValue("num"))
	fmt.Fprint(w, "合計金額は ")
	fmt.Fprintln(w, price*num)
}

func calpmhandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println("errorだよ")
	}
	x, _ := strconv.Atoi(r.FormValue("x"))
	y, _ := strconv.Atoi(r.FormValue("y"))
	switch r.FormValue("cal0") {
	case "+":
		fmt.Fprintln(w, x+y)
	case "-":
		fmt.Fprintln(w, x-y)
	}
}

func sumhandler(w http.ResponseWriter, r *http.Request) {
	var sum, tt int
	if err := r.ParseForm(); err != nil {
		fmt.Println("errorだよ")
	}
	tokuten := strings.Split(r.FormValue("dd"), ",")
	fmt.Println(tokuten)
	for i := range tokuten {
		tt, _ = strconv.Atoi(tokuten[i])
		sum += tt
	}
	fmt.Fprintln(w, sum)
	fmt.Println(sum)
}
func bmicalc(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println("errorだよ")
	}
	weight, _ := strconv.Atoi(r.FormValue("weight"))
	height, _ := strconv.Atoi(r.FormValue("height"))

	fmt.Fprintln(w, float32(weight)/(float32(height*height)/10000.0))

}
func calpmhandler2(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println("errorだよ")
	}
	x, _ := strconv.Atoi(r.FormValue("x"))
	y, _ := strconv.Atoi(r.FormValue("y"))
	switch r.FormValue("cal2") {
	case "+":
		fmt.Fprintln(w, x+y)
	case "-":
		fmt.Fprintln(w, x-y)
	case "*":
		fmt.Fprintln(w, float32(x)*float32(y))
	case "/":
		fmt.Fprintln(w, float32(x)/float32(y))
	}
}
func avehandler(w http.ResponseWriter, r *http.Request) {
	var sum, tt int
	var k [10]int = [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	if err := r.ParseForm(); err != nil {
		fmt.Println("errorだよ")
	}
	tokuten := strings.Split(r.FormValue("ddd"), ",")
	fmt.Println(tokuten)
	for i := range tokuten {
		tt, _ = strconv.Atoi(tokuten[i])
		if tt < 10 {
			k[0]++
		} else if tt < 20 {
			k[1]++
		} else if tt < 30 {
			k[2]++
		} else if tt < 40 {
			k[3]++
		} else if tt < 50 {
			k[4]++
		} else if tt < 60 {
			k[5]++
		} else if tt < 70 {
			k[6]++
		} else if tt < 80 {
			k[7]++
		} else if tt < 90 {
			k[8]++
		} else {
			k[9]++
		}
	}
	sum += tt

	fmt.Fprint(w, "平均：")
	fmt.Fprintln(w, float32(sum)/float32(len(tokuten)))
	fmt.Fprintln(w, "")

	fmt.Fprintln(w, 0)
	var s int
	for s = 0; s < len(k); s++ {
		var i int
		for i = 0; i < k[s]; i++ {
			fmt.Fprint(w, "*")
		}
		fmt.Fprintln(w, "")
		fmt.Fprintln(w, (s+1)*10)
	}
	fmt.Println(sum)
}
