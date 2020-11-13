package models

type User struct {
	ID 				int `json:"id"`
	NamaDepan 		string `json:"namaDepan"`
	NamaBelakang 	string `json:"namaBelakang"`
	Email 			string `json:"email"`
}

//type user
type Users []User