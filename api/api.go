package api

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// ConnectDB ConnectDB
func ConnectDB() {
	var err error
	db, err = sql.Open("mysql", "urlShort:urlshort@/urlShort")
	erring(err)
}

// R Route
func R(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		bin, err := ioutil.ReadFile("./page/index.html")
		erring(err)

		var count int
		db.QueryRow("SELECT COUNT(*) as count FROM shorts").Scan(&count)
		erring(err)

		str := string(bin)
		str = strings.ReplaceAll(str, "{{.}}", strconv.Itoa(count))

		w.Header().Add("Content-Type", "text/html")
		fmt.Fprint(w, str)
	} else if r.URL.Path == "/api" {
		if r.URL.Query().Get("do") != "create" {
			w.WriteHeader(400)
			fmt.Fprint(w, "API: need 'do=create'")
			return
		}

		if r.URL.Query().Get("short") != "" && r.URL.Query().Get("long") != "" {
			var long string
			db.QueryRow("SELECT `long` FROM shorts where short = ?", r.URL.Query().Get("short")).Scan(&long)
			if long != "" {
				w.WriteHeader(403)
				fmt.Fprint(w, "<script>alert('"+r.URL.Query().Get("short")+"는 이미 있는 단축주소 입니다');window.location.replace('/')</script>")
				return
			}

			_, err := db.Query("INSERT INTO shorts (short, `long`) VALUES (?, ?)", r.URL.Query().Get("short"), r.URL.Query().Get("long"))
			if err != nil {
				w.WriteHeader(502)
				fmt.Fprint(w, err.Error())
			} else {
				fmt.Fprintf(w, "<script>alert(\"short.kro.kr%s를 %s에 연결하였습니다\");window.location.replace('/')</script>", r.URL.Query().Get("short"), r.URL.Query().Get("long"))
			}
		}

		w.WriteHeader(400)
		fmt.Fprint(w, "API: need 'short=(url:encoded)&long=(url:encoded)'\nex) short=%2Fnotosans&long=https%3A%2F%2Fwww.google.co.kr%2Fget%2Fnoto")
	} else {
		var long string
		db.QueryRow("SELECT `long` FROM shorts where short = ?", r.URL.Path).Scan(&long)
		if long != "" {
			http.Redirect(w, r, long, 302)
			return
		}

		w.WriteHeader(404)
		http.ServeFile(w, r, "./page/nohere.html")
	}
}

func checkCount(rows *sql.Rows) (count int) {
	for rows.Next() {
		err := rows.Scan(&count)
		erring(err)
	}
	return count
}

func erring(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
