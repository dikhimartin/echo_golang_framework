package models

type SettingPrivilege struct {
	Id             	 string `json:"id"`
	Kode_Privilege 	 string `json:"kode_privilege"`
	Kode_Permissions string `json:"kode_permissions"`
	Name_Menu        string `json:"name_menu"`
	Status           string `json:"status"`
	Keterangan       string `json:"keterangan"`
	Permissions      interface{}
}

type SettingPrivileges struct {
	SettingPrivileges []SettingPrivilege `json:"setting_privilege"`
}

type Permissions struct {
	Key              string
	Value 			 string
	Kode_Privilege   string
}

type MasterPermissions struct {
	Key              string
	Value 			 string
	Additional 		 string
	CheckOrUncheck	 string
}

