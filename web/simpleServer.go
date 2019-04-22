package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()	//解析参数，默认是不会解析的
	//注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	fmt.Println(r.Form)	//这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("value", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello diego!")	//这个写入到w的是输出到客户端的
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)	//获取请求的方法
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("src/web/form.html")
		t.Execute(w, token)
	} else {
		//请求的是登录数据，那么执行登录的逻辑判断
		r.ParseForm()
		if token:=r.Form.Get("token");token!="" {
			//验证token的合法性
		} else {
			//不存在token报错
		}
		fmt.Println("username length:", len(r.Form["username"][0]))
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username")))//输出到服务器端
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username")))		//输出到客户端
	}
}
func main() {
	http.HandleFunc("/", sayHelloName)	//设置访问路由
	http.HandleFunc("/login", login)         //设置访问的路由

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("listenAndServer :", err)
	}
}