package models

type SettingGrup struct {
	Id        string `json:"id"`
	Name_Grup string `json:"name_grup"`
	Status    string `json:"status"`
}

type SettingGrups struct {
	SettingGrups []SettingGrup `json:"setting_grup"`
}
