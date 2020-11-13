package models

type SettingUserGrup struct {
	Id          		string `json:"id"`
	Id_setting_user     string `json:"id_setting_user"`
	Id_setting_grup     string `json:"id_setting_grup"`
	Status      		string `json:"status"`
	Add_Date    		string `json:"add_Date"`
	Update_Date 		string `json:"update_date"`
}

type SettingUserGrups struct {
	SettingUserGrups []SettingUserGrup `json:"setting_user_grup"`
}
