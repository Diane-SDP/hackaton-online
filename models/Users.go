package models

type User struct {
	Id int
	Passwd string
	Uid string
	Name string
	Entreprise int
	Admin int
}

func GetUserByUid(uid string)User {
	rows, err := DB.Query("SELECT id, psswd, uid, name, entreprise, admin FROM users WHERE uid = ?", uid)
	if err != nil {
		panic(err)
	}
	var psswd string
	var id int
	var name string
	var entreprise int
	var admin int
	// defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &psswd, &uid, &name, &entreprise, &admin)
		if err != nil {
			panic(err)
		}
	}
	user := User{
		Id:       id,
		Passwd: psswd,
		Uid:      uid,
		Entreprise: entreprise,
		Admin:    admin,
	}
	return user

}
