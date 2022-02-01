package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("mysession"))

type spasession struct {
	ID      int
	Day     int
	Lmtroom int
	Qry     int
}

var (
	templates *template.Template
)

func connect() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:@/myshop")
	checkerr(err)
	return db
}

func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("รัน !!!!")
	var err error
	templates, err = template.ParseGlob("template/*")
	checkerr(err)
	connect()
	http.HandleFunc("/", index)
	http.HandleFunc("/regis", regis)
	// r.HandleFunc("/login", loginGet).Methods("GET")
	http.HandleFunc("/login", loginPost)
	// r.HandleFunc("/test", testGet).Methods("GET")
	http.HandleFunc("/spa", spa)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/sparoom", sparoom)
	http.HandleFunc("/selectdata", selectdata)
	http.HandleFunc("/selectdataforad", selectdataforad)
	http.HandleFunc("/booking", booking)
	http.HandleFunc("/adspa", adspa)
	http.ListenAndServe(":8080", nil)
	// http.HandleFunc("/test", dbselect)
}

func regis(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		templates.ExecuteTemplate(w, "regis.html", nil)
		return
	}
	db := connect()
	username := r.FormValue("username")
	password := r.FormValue("password")
	var user string
	err := db.QueryRow("SELECT username FROM db_user WHERE username=?", username).Scan(&user)
	if err == sql.ErrNoRows {
		_, err = db.Exec("INSERT INTO db_user(username,password) VALUES(? , ?)", username, password)
		if err != nil {
			http.Error(w, "SERVER ERROR 1", 500)
			return
		}
		// templates.ExecuteTemplate(w, "login.html", "กรุณากรอกให้ถูกต้อง")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else if err != nil {
		http.Error(w, "SERVER ERROR 2", 500)
	}

}

func loginPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		templates.ExecuteTemplate(w, "login.html", "กรุณากรอกให้ถูกต้อง")
		return
	}
	db := connect()
	username := r.FormValue("username")
	password := r.FormValue("password")
	var name string
	var pass string
	var sta string
	err := db.QueryRow("SELECT username,password,u_status FROM db_user WHERE username=? AND password=?", username, password).Scan(&name, &pass, &sta)
	if err != nil {
		templates.ExecuteTemplate(w, "login.html", "เข้าสู่ระบบผิดพลาด")
		return
	}
	session, _ := store.Get(r, "session")
	session.Values["username"] = username
	session.Save(r, w)
	if sta == "1" {
		http.Redirect(w, r, "/spa", http.StatusSeeOther)
	} else if sta == "2" {
		http.Redirect(w, r, "/adspa", http.StatusSeeOther)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}

func spa(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	untyped := session.Values["username"]
	username := untyped.(string)
	templates.ExecuteTemplate(w, "spa.html", username)
}

func adspa(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "adminspa.html", nil)
}

func sparoom(w http.ResponseWriter, r *http.Request) {
	db := connect()
	numroom := r.FormValue("numroom")
	spadate := r.FormValue("bookdate")
	var num, date string
	fmt.Println(numroom, spadate)
	err := db.QueryRow("SELECT * FROM db_spa WHERE s_day=?", spadate).Scan(&num, &date)
	if err == sql.ErrNoRows {
		_, err = db.Exec("INSERT INTO db_spa(s_day,s_max) VALUES(?,?);", spadate, numroom)
		if err != nil {
			http.Error(w, "ERROR", 500)
			return
		}
		templates.ExecuteTemplate(w, "adminspa.html", "INSERT Successfully")
	} else if err != nil {
		var re sql.Result
		upd, err := db.Prepare("UPDATE db_spa SET s_max = ? WHERE s_day = ?")
		checkerr(err)
		re, err = upd.Exec(numroom, spadate)
		rowsAff, aa := re.RowsAffected()
		if err != nil || rowsAff != 1 {
			fmt.Println(aa)
			templates.ExecuteTemplate(w, "adminspa.html", "Have some error, Please check")
			return
		} else {
			templates.ExecuteTemplate(w, "adminspa.html", "Update Successfully")
		}
	}
}

func booking(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "booking.html", nil)
	session, _ := store.Get(r, "session")
	untyped := session.Values["username"]
	username := untyped.(string)
	if r.Method == "POST" {
		db := connect()
		var bd bdd
		bookdate := r.FormValue("bookdate")
		err := db.QueryRow("SELECT *,s_max,s_qty FROM db_spa WHERE s_day = ?", bookdate)
		if err != nil {
			_, err := db.Exec("INSERT INTO db_order(o_user,s_day) VALUES(?,?)", username, bookdate)
			checkerr(err)
			err = db.QueryRow("SELECT COUNT(*) as number FROM db_order WHERE s_day = ?", bookdate).Scan(&bd.bdday)
			err2 := db.QueryRow("SELECT s_max,s_qty FROM db_spa WHERE s_day = ?", bookdate).Scan(&bd.bdmax, &bd.bdqty)
			checkerr(err)
			checkerr(err2)
			if bd.bdqty < bd.bdmax {
				var up sql.Result
				upda, err := db.Prepare("UPDATE db_spa SET s_qty = ? WHERE s_day = ?")
				checkerr(err)
				up, err = upda.Exec(bd.bdday, bookdate)
				rowsAff, aa := up.RowsAffected()
				if err != nil || rowsAff != 1 {
					fmt.Println(aa)
				}
				log.Println("รวมวัน ", bd.bdday)
				log.Println("จำนวนห้อง ", bd.bdmax)
				log.Println("จำนวนที่จอง", bd.bdqty)
				templates.ExecuteTemplate(w, "booking.html", "จองเสร็จสมบูรณ์")
			} else {
				templates.ExecuteTemplate(w, "booking.html", "FULL")
			}
		} else {
			templates.ExecuteTemplate(w, "booking.html", "ไม่มีห้องที่ต้องการเช่า")
		}
	}
}

func selectdata(w http.ResponseWriter, r *http.Request) {
	db := connect()
	rows, err := db.Query("SELECT * FROM db_spa Order by s_day ASC")
	checkerr(err)
	var spases []spasession
	for rows.Next() {
		var s spasession
		err = rows.Scan(&s.ID, &s.Day, &s.Qry, &s.Lmtroom)
		if err != nil {
			panic(err)
		}
		spases = append(spases, s)
	}
	templates.ExecuteTemplate(w, "spa.html", spases)
}

func selectdataforad(w http.ResponseWriter, r *http.Request) {
	db := connect()
	rows, err := db.Query("SELECT * FROM db_spa Order by s_day ASC")
	checkerr(err)
	var spases []spasession
	for rows.Next() {
		var s spasession
		err = rows.Scan(&s.ID, &s.Day, &s.Qry, &s.Lmtroom)
		if err != nil {
			panic(err)
		}
		spases = append(spases, s)
	}
	templates.ExecuteTemplate(w, "adminspa.html", spases)
}

type bdd struct {
	bdday int
	bdmax int
	bdqty int
}

func logout(w http.ResponseWriter, r *http.Request) {
	clearSession(w)
	http.Redirect(w, r, "/", http.StatusFound)
}

func clearSession(w http.ResponseWriter) {
	session := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, session)
}
