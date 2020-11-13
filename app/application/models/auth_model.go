package models


type GetDataLogin struct {
	Id_users       	string `json:"id_users"`
	Id_group       	string `json:"id_group"`
	Name_users     	string `json:"name_users"`
	Name_group     	string `json:"name_group"`    
    Jti    			string
    Exp    	   	 	string
    Image    	   	string
    Extension    	string
}