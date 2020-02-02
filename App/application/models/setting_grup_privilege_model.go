package models

type SettingGrupPrivilege struct {
	Id             		   string `json:"id"`
	Id_setting_grup        string `json:"id_setting_grup"`
	Name_grup        	   string `json:"name_grup"`
	Kode_Privilege 		   string `json:"kode_privilege"`
	Hak_Akses              string `json:"hak_akses"`
	Created_at             string `json:"created_at"`
	Updated_at             string `json:"updated_at"`
	Status           	   string `json:"status"`
	Keterangan     		   string `json:"keterangan"`
}

type SettingGrupPrivileges struct {
	SettingGrupPrivileges []SettingGrupPrivilege `json:"setting_grup_privilege"`
}

