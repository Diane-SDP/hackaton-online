package models

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"time"
)

type User struct {
	Id         int
	Passwd     string
	Uid        string
	Name       string
	Entreprise int
	Admin      int
}

func GetUserByUid(uid string) User {
	rows, err := DB.Query("SELECT id, passwd, uid, name, admin FROM users WHERE uid = ?", uid)
	if err != nil {
		panic(err)
	}
	var psswd string
	var id int
	var name string
	var admin int
	// defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &psswd, &uid, &name, &admin)
		if err != nil {
			panic(err)
		}
	}
	user := User{
		Id:         id,
		Passwd:     psswd,
		Uid:        uid,
		Admin:      admin,
	}
	return user

}
func AddUser(pseudo string, psswd string,email string) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	var uid string
	for i := 0; i < 16; i++ {
		uid += string(charset[seededRand.Intn(len(charset))])
	}
	stmt, err := DB.Prepare("INSERT INTO users(uid, name, passwd,Mail,admin ) VALUES(?,?, ?, ?, ?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(uid, pseudo, psswd,email, 1)
	if err != nil {
		panic(err)
	}
}

func ExistAccount(Pseudo string) (bool, string, string) {
	rows, err := DB.Query("SELECT name , passwd, uid FROM users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var each_pseudo string
		var each_psswd string
		var uid string
		_ = rows.Scan(&each_pseudo, &each_psswd, &uid)
		if each_pseudo == Pseudo {
			return true, each_psswd, uid
		}
	}
	return false, "", "oui"
}

func Getuid(Pseudo string) string {
	rows, err := DB.Query("SELECT uid FROM users WHERE name = ?", Pseudo)
	if err != nil {
		panic(err)
	}
	var uid string
	// defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&uid)
		if err != nil {
			panic(err)
		}
	}
	return uid
}

func CreateDefaultAdmin(){
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	var uid string
	for i := 0; i < 16; i++ {
		uid += string(charset[seededRand.Intn(len(charset))])
	}
	mdp := sha256.Sum256([]byte("admin"))
	mdphash :=  hex.EncodeToString((mdp[:]))
	stmt, err := DB.Prepare("INSERT INTO users(uid, name, passwd,admin ) VALUES(?, ?, ?, ?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(uid, "admin", mdphash, 1)
	if err != nil {
		panic(err)
	}
}
func DeleteDefaultAdmin(){
	_,err := DB.Exec("Delete from users where name = ?","admin")
	if err != nil {
		panic(err)
	}
}

func ExistAdmin()bool{
    var count int
    err := DB.QueryRow("SELECT COUNT(id) from users").Scan(&count)
    if err != nil {
        panic(err)
    }
    return count > 0
}