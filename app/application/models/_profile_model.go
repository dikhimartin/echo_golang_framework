package models


type GetDataProfile struct {
	Id_users       	string `json:"id_users"`
	Id_group       	string `json:"id_group"`
	Name_users     	string `json:"name_users"`
	Name_group     	string `json:"name_group"`    
    Email           string `json:"email"`    
    Telephone       string `json:"telephone"`    
    Address         string `json:"address"`    
    Image    	   	string
    Extension    	string
}