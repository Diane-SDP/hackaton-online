package models

type Colis struct {
	Id          int
	Uid         string
	IdShop      int
	StartAdress string
	FinalAdress string
	Step        int
}

func AddColis(code string, idshop int, startaddr string, finaladdr string) {
	stmt, err := DB.Prepare("INSERT INTO Colis(Uid, IdShop, StartAdress, FinalAdress, Step) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(code, idshop, startaddr, finaladdr, 0)
	if err != nil {
		panic(err)
	}
}

func Exist(code string) bool {
	rows, err := DB.Query("SELECT Uid FROM colis WHERE Uid = ?", code)
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
	rows, err := DB.Query("SELECT Id, IdShop, StartAdress, FinalAdress, Step FROM colis WHERE Uid = ?", code)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var colis Colis
	for rows.Next() {
		err := rows.Scan(&colis.Id, &colis.IdShop, &colis.StartAdress, &colis.FinalAdress, &colis.Step)
		if err != nil {
			panic(err)
		}
	}
	colis.Uid = code
	return colis
}
