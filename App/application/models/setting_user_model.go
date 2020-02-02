package models

type SettingUser struct {
	Id          string `json:"id"`
	Id_group 	string `json:"id_group"`
	Name_grup 	string `json:"name_grup"`
	Full_name   string `json:"full_name"`
	Username 	string `json:"username"`
	Email 		string `json:"email"`
	Password 	string `json:"password"`
	Gender 		string `json:"gender"`
	Salt        string `json:"salt"`
	Add_Date    string `json:"add_Date"`
	Update_Date string `json:"update_date"`
	Status      string `json:"status"`
	Image    	string `json:"image"`
	Additional  string `json:"additional"`
}

type SettingUsers struct {
	SettingUsers []SettingUser `json:"setting_user"`
}


