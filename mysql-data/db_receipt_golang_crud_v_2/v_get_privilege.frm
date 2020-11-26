TYPE=VIEW
query=select `db_receipt_golang_crud_v_2`.`tb_setting_privilege_detail`.`id` AS `id`,`db_receipt_golang_crud_v_2`.`tb_setting_privilege`.`id` AS `id_setting_privilege`,`db_receipt_golang_crud_v_2`.`tb_setting_privilege`.`name_menu` AS `name_menu`,`db_receipt_golang_crud_v_2`.`tb_setting_privilege`.`code_privilege` AS `code_privilege`,concat(concat(`db_receipt_golang_crud_v_2`.`tb_setting_privilege`.`code_privilege`,\'_\',`db_receipt_golang_crud_v_2`.`tb_setting_privilege_detail`.`permissions`)) AS `code_permissions`,`db_receipt_golang_crud_v_2`.`tb_setting_privilege`.`status` AS `STATUS`,`db_receipt_golang_crud_v_2`.`tb_setting_privilege_detail`.`permissions` AS `permissions` from (`db_receipt_golang_crud_v_2`.`tb_setting_privilege_detail` left join `db_receipt_golang_crud_v_2`.`tb_setting_privilege` on((`db_receipt_golang_crud_v_2`.`tb_setting_privilege_detail`.`id_setting_privilege` = `db_receipt_golang_crud_v_2`.`tb_setting_privilege`.`id`)))
md5=5c4ca3c17d64613859ca7929d3f32560
updatable=0
algorithm=0
definer_user=user
definer_host=%
suid=1
with_check_option=0
timestamp=2020-11-25 11:05:40
create-version=1
source=SELECT\n				tb_setting_privilege_detail.id AS id,\n				tb_setting_privilege.id AS id_setting_privilege,\n				tb_setting_privilege.name_menu AS name_menu,\n				tb_setting_privilege.code_privilege AS code_privilege,\n				concat( concat( tb_setting_privilege.code_privilege, \'_\', tb_setting_privilege_detail.permissions ) ) AS code_permissions,\n				tb_setting_privilege.STATUS AS STATUS,\n				tb_setting_privilege_detail.permissions AS permissions \n			FROM\n				tb_setting_privilege_detail\n				LEFT JOIN tb_setting_privilege ON tb_setting_privilege_detail.id_setting_privilege = tb_setting_privilege.id
client_cs_name=utf8mb4
connection_cl_name=utf8mb4_general_ci
view_body_utf8=select `db_receipt_golang_crud_v_2`.`tb_setting_privilege_detail`.`id` AS `id`,`db_receipt_golang_crud_v_2`.`tb_setting_privilege`.`id` AS `id_setting_privilege`,`db_receipt_golang_crud_v_2`.`tb_setting_privilege`.`name_menu` AS `name_menu`,`db_receipt_golang_crud_v_2`.`tb_setting_privilege`.`code_privilege` AS `code_privilege`,concat(concat(`db_receipt_golang_crud_v_2`.`tb_setting_privilege`.`code_privilege`,\'_\',`db_receipt_golang_crud_v_2`.`tb_setting_privilege_detail`.`permissions`)) AS `code_permissions`,`db_receipt_golang_crud_v_2`.`tb_setting_privilege`.`status` AS `STATUS`,`db_receipt_golang_crud_v_2`.`tb_setting_privilege_detail`.`permissions` AS `permissions` from (`db_receipt_golang_crud_v_2`.`tb_setting_privilege_detail` left join `db_receipt_golang_crud_v_2`.`tb_setting_privilege` on((`db_receipt_golang_crud_v_2`.`tb_setting_privilege_detail`.`id_setting_privilege` = `db_receipt_golang_crud_v_2`.`tb_setting_privilege`.`id`)))
