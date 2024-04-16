package models

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"time"
)

type Shop struct {
	Name      string
	Adress    string
	Mail      string
	Id        int
	Uid       string
	Categorie string
}

func AddShop(name string, mail string, adresse string, categorie string) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	var uid string
	for i := 0; i < 16; i++ {
		uid += string(charset[seededRand.Intn(len(charset))])
	}
	var pwd string
	for i := 0; i < 20; i++ {
		pwd += string(charset[seededRand.Intn(len(charset))])
	}
	passloghash := sha256.Sum256([]byte(pwd))
	hash := hex.EncodeToString((passloghash[:]))
	SendMail(mail,"Bonjour "+name+"\nMerci pour vous être affilié a nous , voici vos identifiants pour vous connecter a notre site et gérer les livraisons :\n Nom : "+name+"\nMot de passe : "+pwd,"Affiliation TrackTheur")
	stmt, err := DB.Prepare("INSERT INTO Shops(Uid, Name, Passwd, Mail, Adress, Categorie) VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(uid, name, hash, mail, adresse, categorie)
	if err != nil {
		panic(err)
	}
}

func GetAllShops() []Shop {
	rows, err := DB.Query("SELECT id, uid, Name, Mail, Adress, Categorie FROM Shops")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var AllShops []Shop
	for rows.Next() {
		var shop Shop
		err = rows.Scan(&shop.Id, &shop.Uid, &shop.Name, &shop.Mail, &shop.Adress, &shop.Categorie)
		if err != nil {
			panic(err)
		}
		AllShops = append(AllShops, shop)
	}
	return AllShops
}

func GetShop(id int) Shop {
	rows, err := DB.Query("SELECT uid, Name, Mail, Adress, Categorie FROM Shops WHERE id = ?", id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var shop Shop
	for rows.Next() {
		var shop Shop
		err = rows.Scan(&shop.Uid, &shop.Name, &shop.Mail, &shop.Adress, &shop.Categorie)
		if err != nil {
			panic(err)
		}
	}
	shop.Id = id
	return shop
}

func GetShopByUid(uid string) Shop {
	rows, err := DB.Query("SELECT id, Name, Mail, Adress, Categorie FROM Shops WHERE uid = ?", uid)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var shop Shop
	for rows.Next() {
		err = rows.Scan(&shop.Id, &shop.Name, &shop.Mail, &shop.Adress, &shop.Categorie)
		if err != nil {
			panic(err)
		}
	}
	shop.Uid = uid
	return shop
}

func ExistShop(mail string) (bool, string, string) {
	rows, err := DB.Query("SELECT passwd, uid FROM shops WHERE mail = ?", mail)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var passwd string
		var uid string
		_ = rows.Scan(&passwd, &uid)
		return true, passwd, uid
	}
	return false, "", "oui"
}
