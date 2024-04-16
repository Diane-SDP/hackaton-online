package models

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"net/url"
	"strconv"
)

type Colis struct {
	Id          int
	Uid         string
	Shop        Shop
	StartAdress string
	FinalAdress string
	StartCity   string
	ClientCity  string
	Step        int
	Distance    int
	Position    int
	Mail string
}

type Location struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}
type LocationSTR struct {
	LatSTR string `json:"lat"`
	LonSTR string `json:"lon"`
}

func AddColis(code string, idshop int, startaddr string, finaladdr string, startCity string, clientCity string,Email string) {
	stmt, err := DB.Prepare("INSERT INTO Colis(Uid, IdShop, StartAdress, FinalAdress, StartCity, ClientCity, Step,Distance,Position,Email) VALUES(?, ?, ?, ?, ?,?,?,?,?,?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	distance := GetDistanceFromStart(startCity, clientCity)
	Distance := int(distance)
	var position int
	if distance <= 100 {
		position = int(distance / 3)
	} else if distance <= 500 {
		position = int(distance*0.08 + 25.21)
	} else if distance <= 1000 {
		position = int(distance*0.068 + 31)
	} else {
		position = 99
	}
	_, err = stmt.Exec(code, idshop, startaddr, finaladdr, startCity, clientCity, 0, Distance, position,Email)
	if err != nil {
		panic(err)
	}
}

func GetDistanceFromStart(start string, destination string) float64 {
	startCoor := getLocationCoordinates(start)
	destinationCoor := getLocationCoordinates(destination)

	radlat1 := float64(math.Pi * startCoor.Lat / 180)
	radlat2 := float64(math.Pi * destinationCoor.Lat / 180)

	theta := float64(startCoor.Lon - destinationCoor.Lon)
	radtheta := float64(math.Pi * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)
	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / math.Pi
	dist = dist * 60 * 1.1515
	dist = dist * 1.609344

	return dist
}

func getLocationCoordinates(address string) Location {
	url := fmt.Sprintf("https://nominatim.openstreetmap.org/search?format=json&q=%s", url.QueryEscape(address))
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var locations []Location
	var locationSTR []LocationSTR

	err = json.NewDecoder(resp.Body).Decode(&locationSTR)
	if err != nil {
		panic(err)
	}

	for _, elt := range locationSTR {
		var loc Location
		loc.Lat, _ = strconv.ParseFloat(elt.LatSTR, 64)
		loc.Lon, _ = strconv.ParseFloat(elt.LonSTR, 64)
		locations = append(locations, loc)
	}

	if len(locations) == 0 {
		return Location{}
	}

	return locations[0]
}

func GetAllColis() []Colis {
	rows, err := DB.Query("SELECT id, uid, StartAdress, FinalAdress, StartCity, ClientCity, step, idShop FROM Colis")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var AllColis []Colis
	for rows.Next() {
		var id int
		var uid string
		var StartAdress string
		var FinalAdress string
		var step int
		var IdEntreprise int
		var TheColis Colis
		err = rows.Scan(&id, &uid, &StartAdress, &FinalAdress, &TheColis.StartCity, &TheColis.ClientCity, &step, &IdEntreprise)
		if err != nil {
			panic(err)
		}
		TheColis.Id = id
		TheColis.Uid = uid
		TheColis.StartAdress = StartAdress
		TheColis.FinalAdress = FinalAdress
		TheColis.Step = step
		TheColis.Shop = GetShop(IdEntreprise)
		AllColis = append(AllColis, TheColis)
	}
	return AllColis
}

func GetColisOf(idshop int) []Colis {
	rows, err := DB.Query("SELECT id, uid, StartAdress, FinalAdress, StartCity, ClientCity, step, idShop FROM Colis WHERE idShop = ?", idshop)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var AllColis []Colis
	for rows.Next() {
		var id int
		var uid string
		var StartAdress string
		var FinalAdress string
		var step int
		var IdEntreprise int
		var TheColis Colis
		err = rows.Scan(&id, &uid, &StartAdress, &FinalAdress, &TheColis.StartCity, &TheColis.ClientCity, &step, &IdEntreprise)
		if err != nil {
			panic(err)
		}
		TheColis.Id = id
		TheColis.Uid = uid
		TheColis.StartAdress = StartAdress
		TheColis.FinalAdress = FinalAdress
		TheColis.Step = step
		TheColis.Shop = GetShop(IdEntreprise)
		AllColis = append(AllColis, TheColis)
	}
	return AllColis
}

func Exist(code string) bool {
	rows, err := DB.Query("SELECT Uid FROM Colis WHERE Uid = ?", code)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var uid string
	for rows.Next() {
		err := rows.Scan(&uid)
		if err != nil {
			panic(err)
		}
		return true
	}
	return false
}

func GetColis(code string) Colis {
	rows, err := DB.Query("SELECT Id, IdShop, StartAdress, FinalAdress, StartCity, ClientCity, Step ,Distance,Position FROM Colis WHERE Uid = ?", code)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var colis Colis
	var idShop int
	for rows.Next() {
		err := rows.Scan(&colis.Id, &idShop, &colis.StartAdress, &colis.FinalAdress, &colis.StartCity, &colis.ClientCity, &colis.Step, &colis.Distance, &colis.Position)
		if err != nil {
			panic(err)
		}
	}
	// distance := GetDistanceFromStart(colis.StartCity, colis.ClientCity)
	// colis.Distance = int(distance)
	// var position int
	// if distance <= 100 {
	// 	position = int(distance / 3)
	// } else if distance <= 500 {
	// 	position = int(distance*0.08 + 25.21)
	// } else if distance <= 1000 {
	// 	position = int(distance*0.068 + 31)
	// } else {
	// 	position = 99
	// }
	// println(position)
	println("truc",colis.Distance, colis.Position)
	colis.Shop = GetShop(idShop)
	colis.Uid = code
	return colis
}

func GetColisbyid(id int) Colis {
	rows, err := DB.Query("SELECT Id,Uid, IdShop, StartAdress, FinalAdress, StartCity, ClientCity, Step,Distance,Position,Email FROM Colis WHERE id = ?", id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var colis Colis
	var IdEntreprise int
	for rows.Next() {
		err := rows.Scan(&colis.Id, &colis.Uid,&IdEntreprise, &colis.StartAdress, &colis.FinalAdress, &colis.StartCity, &colis.ClientCity, &colis.Step,&colis.Distance,&colis.Position,&colis.Mail)
		if err != nil {
			panic(err)
		}
	}
	colis.Shop = GetShop(IdEntreprise)
	colis.Id = id
	return colis
}
