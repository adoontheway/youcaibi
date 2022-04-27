package main

import (
	"encoding/json"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func homeHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 获取session和user
	cname, err1 := r.Cookie("username")
	sid, err2 := r.Cookie("session")
	if err1 != nil || err2 != nil {
		// 不存在的话就走登录流程
		h := &HomePage{
			Name: "ado",
		}
		t, err := template.ParseFiles("../template/home.html")
		if err != nil {
			log.Printf("Parse template home.html error:%v", err)
			return
		}
		t.Execute(w, h)
		return
	}
	// 存在的话就重定向到用户主页
	// TODO 要不要检查session和user的合法性
	if len(cname.Value) != 0 && len(sid.Value) != 0 {
		http.Redirect(w, r, "/userhome", http.StatusFound)
		return
	}

}

func userHomeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// check if is visitor
	cname, err1 := r.Cookie("username")
	_, err2 := r.Cookie("session")
	if err1 != nil || err2 != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	fname := r.FormValue("username")
	var p *UserPage
	if len(cname.Value) != 0 {
		p = &UserPage{Name: cname.Value}
	} else if len(fname) != 0 {
		p = &UserPage{Name: fname}
	}

	t, err := template.ParseFiles("../template/userhome.html")
	if err != nil {
		log.Fatal("Parseing userhome.html template errro: %v", err)
		return
	}
	t.Execute(w, p)
}

func apiHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if r.Method != http.MethodPost {
		re, _ := json.Marshal(ErrorRequestNotRecognized)
		io.WriteString(w, string(re))
		return
	}

	res, _ := ioutil.ReadAll(r.Body)
	apibody := &ApiBody{}
	if err := json.Unmarshal(res, apibody); err != nil {
		re, _ := json.Marshal(ErrorRequestBodyParseFailed)
		io.WriteString(w, string(re))
		return
	}
	request(apibody, w, r)
	defer r.Body.Close()
}
