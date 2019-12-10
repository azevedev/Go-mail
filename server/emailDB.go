package main

import (
	"database/sql"
	"fmt" 
	"strings"
	"strconv"
  _ "github.com/lib/pq"
)

func insertEmail(db *sql.DB, e Email) {
	fmt.Println("# Inserting values")
	var lastID int
	
	sql := `SELECT id FROM emails order by id desc limit 1`
	rows, err := db.Query(sql)
	if rows.Next(){
		err = rows.Scan(&lastID)
		fmt.Println("last inserted id =", lastID)
		checkErr(err)
	
		sql = `INSERT INTO emails("to",subject,id,"from",created_at, user_id) VALUES ('`+e.to+`', '`+e.subject+`', `+strconv.Itoa(lastID+1)+`, '`+e.from+`', '`+e.date+`', `+strconv.Itoa(e.user_id)+`)`
		_, err = db.Exec(sql)
		checkErr(err)
	}else{
		fmt.Println("ERROR ON DB")
	}
}
func getEmails(db *sql.DB){

  rows, err := db.Query("SELECT * FROM emails")
  if err != nil {
    // handle this error better than this
    panic(err)
  }
  defer rows.Close()
  for rows.Next() {
    var id int
    var sub string
	var to string
	var from string 
	var date string
	var user_id int
	err = rows.Scan(&to, &sub, &id, &from, &date, &user_id)
	s := strings.Split(date,"T")
	date = s[0]
    if err != nil {
      // handle this error
      panic(err)
    }
	fmt.Println(id, from, sub, to, date, user_id)
	
  }
  // get any error encountered during iteration
  err = rows.Err()
  if err != nil {
    panic(err)
  }
}

func findEmail(db *sql.DB, id int) Email{
  rows, err := db.Query(`SELECT * FROM emails WHERE id=`+strconv.Itoa(id))
  if err != nil {
    // handle this error better than this
    panic(err)
  }
  defer rows.Close()
  if rows.Next() {
    var id int
    var sub string
	var to string
	var from string 
	var date string
	var user_id int
	err = rows.Scan(&to, &sub, &id, &from, &date, &user_id)
	s := strings.Split(date,"T")
	date = s[0]
    if err != nil {
      // handle this error
      panic(err)
	}
	e := Email{
		id: id,
		subject: sub,
		to: to,
		from: from,
		date: date,
		user_id: user_id,
	}
	return e
  }
  // get any error encountered during iteration
  err = rows.Err()
  if err != nil {
    panic(err)
  }
  return Email{}
}

func deleteEmail(db *sql.DB, e Email){
	fmt.Println("# Deleting")
	sql := `DELETE FROM emails WHERE id=`+strconv.Itoa(e.id)
	stmt, err := db.Prepare(sql)
	checkErr(err)

	res, err := stmt.Exec()
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect, "rows changed")

}