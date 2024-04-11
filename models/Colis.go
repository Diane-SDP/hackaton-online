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
	IdShop      int
	StartAdress string
	FinalAdress string
	StartCity   string
	ClientCity  string
	Step        int
	Distance    float64
}

type Location struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}
type LocationSTR struct {
	LatSTR string `json:"lat"`
	LonSTR string `json:"lon"`
}

func AddColis(code string, idshop int, startaddr string, finaladdr string, startCity string, clientCity string) {
	stmt, err := DB.Prepare("INSERT INTO Colis(Uid, IdShop, StartAdress, FinalAdress, StartCity, ClientCity, Step) VALUES(?, ?, ?, ?, ?,?,?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(code, idshop, startaddr, finaladdr, startCity, clientCity, 0)
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
	println("tarpin grosse PUTE MELISSANDRE")
	url := fmt.Sprintf("https://nominatim.openstreetmap.org/search?format=json&q=%s", url.QueryEscape(address))
	println(url)
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
		println("rien de trouvé")
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
		TheColis.IdShop = IdEntreprise
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
	rows, err := DB.Query("SELECT Id, IdShop, StartAdress, FinalAdress, StartCity, ClientCity, Step FROM Colis WHERE Uid = ?", code)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var colis Colis
	for rows.Next() {
		err := rows.Scan(&colis.Id, &colis.IdShop, &colis.StartAdress, &colis.FinalAdress, &colis.StartCity, &colis.ClientCity, &colis.Step)
		if err != nil {
			panic(err)
		}
	}
	colis.Uid = code
	return colis
}

func GetColisbyid(id int) Colis {
	rows, err := DB.Query("SELECT Id, IdShop, StartAdress, FinalAdress, StartCity, ClientCity, Step FROM Colis WHERE id = ?", id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var colis Colis
	for rows.Next() {
		err := rows.Scan(&colis.Id, &colis.IdShop, &colis.StartAdress, &colis.FinalAdress, &colis.StartCity, &colis.ClientCity, &colis.Step)
		if err != nil {
			panic(err)
		}
	}
	colis.Id = id
	return colis
}
