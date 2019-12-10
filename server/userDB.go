package main

import (
	"database/sql"
	"fmt" 
	"strconv"
  _ "github.com/lib/pq"
)

func insertUser(db *sql.DB, u User) {
	fmt.Println("# Inserting values")
	var lastID int
	
	sql := `SELECT id FROM users order by id desc limit 1`
	rows, err := db.Query(sql)
	if rows.Next(){
		err = rows.Scan(&lastID)
		fmt.Println("last inserted id =", lastID)
		checkErr(err)
		// INSERT INTO public.users(
		// 	name, id, email, password)
		// 	VALUES (?, ?, ?, ?);
		sql = `INSERT INTO users(name,id,email,password) VALUES ('`+u.name+`',  `+strconv.Itoa(lastID+1)+`, '`+u.email+`', '`+u.password+`')`
		_, err = db.Exec(sql)
		checkErr(err)
	}else{
		fmt.Println("ERROR ON DB")
	}
}
func getUsers(db *sql.DB){

  rows, err := db.Query("SELECT * FROM users")
  if err != nil {
    // handle this error better than this
    panic(err)
  }
  defer rows.Close()
  for rows.Next() {
    var id int
    var name string
	var email string
	var pass string 
	err = rows.Scan(&name, &id, &email, &pass)
    if err != nil {
      // handle this error
      panic(err)
    }
	fmt.Println(id, name, email, pass)
	
  }
  // get any error encountered during iteration
  err = rows.Err()
  if err != nil {
    panic(err)
  }
}

func findUser(db *sql.DB, id int) User{
  rows, err := db.Query(`SELECT * FROM users WHERE id=`+strconv.Itoa(id))
  if err != nil {
    // handle this error better than this
    panic(err)
  }
  defer rows.Close()
  if rows.Next() {
    var id int
    var name string
	var email string
	var pass string 
	err = rows.Scan(&name, &id, &email, &pass)
    if err != nil {
      // handle this error
      panic(err)
	}
	u := User{
		id: id,
		name: name,
		email: email,
		password: pass,
	}
	return u
  }
  // get any error encountered during iteration
  err = rows.Err()
  if err != nil {
    panic(err)
  }
  return User{}
}

func deleteUser(db *sql.DB, u User){
	fmt.Println("# Deleting")
	sql := `DELETE FROM users WHERE id=`+strconv.Itoa(u.id)
	stmt, err := db.Prepare(sql)
	checkErr(err)

	res, err := stmt.Exec()
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect, "rows changed")

}