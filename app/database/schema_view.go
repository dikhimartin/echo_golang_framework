package database


func v_get_user() string{
	schema := `SELECT
				tb_setting_user.id AS id,
				tb_setting_user_grup.id_setting_grup AS id_setting_grup,
				tb_setting_user.full_name AS full_name,
				tb_setting_user.gender AS gender,
				tb_setting_user.email AS email,
				tb_setting_user.telephone AS telephone,
				tb_setting_user.address AS address,
				tb_setting_user.username AS username,
				tb_setting_grup.name_grup AS name_grup,
				tb_setting_user.STATUS AS STATUS,
				tb_setting_user.PASSWORD AS PASSWORD,
				tb_setting_user.auth_token AS auth_token,
				tb_setting_user.image AS image,
				tb_setting_user.additional AS extension 
			FROM
				tb_setting_user
				JOIN tb_setting_user_grup ON tb_setting_user.id = tb_setting_user_grup.id_setting_user
				JOIN tb_setting_grup ON tb_setting_user_grup.id_setting_grup = tb_setting_grup.id`
	return schema
}

func v_get_grup() string{
	schema := `SELECT
				tb_setting_grup.id AS id,
				tb_setting_grup_privilege.id AS id_setting_grup_privilege,
				tb_setting_grup.name_grup AS name_grup,
				tb_setting_grup.status AS status 
			FROM
	 			tb_setting_grup LEFT JOIN tb_setting_grup_privilege ON   tb_setting_grup.id = tb_setting_grup_privilege.id_setting_grup`
	return schema
}

func v_get_user_grup() string {
	schema := `SELECT
				tb_setting_user_grup.id AS id,
				tb_setting_user.id AS id_setting_user,
				tb_setting_grup.id AS id_setting_grup,
				tb_setting_user.username AS username,
				tb_setting_user.full_name AS full_name,
				tb_setting_user.email AS email,
				tb_setting_user.image AS image,
				tb_setting_grup.name_grup AS name_grup,
				tb_setting_grup.STATUS AS STATUS,
				tb_setting_user.additional AS extension,
				tb_setting_user.auth_token AS auth_token,
				tb_setting_user.password AS password 
			FROM
				tb_setting_user_grup
				JOIN tb_setting_grup ON tb_setting_grup.id = tb_setting_user_grup.id_setting_grup
				JOIN tb_setting_user ON tb_setting_user.id = tb_setting_user_grup.id_setting_user`
	return schema
}

func v_get_privilege() string{
	schema := `SELECT
				tb_setting_privilege_detail.id AS id,
				tb_setting_privilege.id AS id_setting_privilege,
				tb_setting_privilege.name_menu AS name_menu,
				tb_setting_privilege.code_privilege AS code_privilege,
				concat( concat( tb_setting_privilege.code_privilege, '_', tb_setting_privilege_detail.permissions ) ) AS kode_permissions,
				tb_setting_privilege.STATUS AS STATUS,
				tb_setting_privilege_detail.permissions AS permissions 
			FROM
				tb_setting_privilege_detail
				LEFT JOIN tb_setting_privilege ON tb_setting_privilege_detail.id_setting_privilege = tb_setting_privilege.id`
	return schema
}

func v_get_grup_privilege() string{
	schema := `SELECT
				tb_setting_grup_privilege.id AS id,
				tb_setting_grup_privilege.id_setting_grup AS id_setting_grup,
				tb_setting_grup.name_grup AS name_grup,
				tb_setting_grup_privilege.remarks AS remarks,
				tb_setting_grup_privilege.STATUS AS STATUS,
				tb_setting_grup_privilege.created_at AS created_at,
				tb_setting_grup_privilege.updated_at AS updated_at 
			FROM
				tb_setting_grup_privilege
				JOIN tb_setting_grup ON tb_setting_grup_privilege.id_setting_grup = tb_setting_grup.id`
	return schema
}

func v_get_grup_privilege_detail()string{
	schema := `SELECT
				tb_setting_grup_privilege_detail.id AS id,
				tb_setting_grup_privilege.id_setting_grup AS id_setting_grup,
				tb_setting_grup_privilege_detail.id_setting_grup_privilege AS id_setting_grup_privilege,
				tb_setting_grup.name_grup AS name_grup,
				tb_setting_grup_privilege_detail.code_permissions AS code_permissions 
			FROM
				tb_setting_grup_privilege_detail
				JOIN tb_setting_grup_privilege ON tb_setting_grup_privilege_detail.id_setting_grup_privilege = tb_setting_grup_privilege.id
				JOIN tb_setting_grup ON tb_setting_grup_privilege.id_setting_grup = tb_setting_grup.id`
	return schema
}
