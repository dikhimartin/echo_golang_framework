TYPE=VIEW
query=select `db_receipt_golang_crud_v_2`.`tb_setting_grup_privilege_detail`.`id` AS `id`,`db_receipt_golang_crud_v_2`.`tb_setting_grup_privilege`.`id_setting_grup` AS `id_setting_grup`,`db_receipt_golang_crud_v_2`.`tb_setting_grup_privilege_detail`.`id_setting_grup_privilege` AS `id_setting_grup_privilege`,`db_receipt_golang_crud_v_2`.`tb_setting_grup`.`name_grup` AS `name_grup`,`db_receipt_golang_crud_v_2`.`tb_setting_grup_privilege_detail`.`code_permissions` AS `code_permissions` from ((`db_receipt_golang_crud_v_2`.`tb_setting_grup_privilege_detail` join `db_receipt_golang_crud_v_2`.`tb_setting_grup_privilege` on((`db_receipt_golang_crud_v_2`.`tb_setting_grup_privilege_detail`.`id_setting_grup_privilege` = `db_receipt_golang_crud_v_2`.`tb_setting_grup_privilege`.`id`))) join `db_receipt_golang_crud_v_2`.`tb_setting_grup` on((`db_receipt_golang_crud_v_2`.`tb_setting_grup_privilege`.`id_setting_grup` = `db_receipt_golang_crud_v_2`.`tb_setting_grup`.`id`)))
md5=432fe28409614abcb86dd7172b019909
updatable=1
algorithm=0
definer_user=user
definer_host=%
suid=1
with_check_option=0
timestamp=2020-11-25 11:05:40
create-version=1
source=SELECT\n				tb_setting_grup_privilege_detail.id AS id,\n				tb_setting_grup_privilege.id_setting_grup AS id_setting_grup,\n				tb_setting_grup_privilege_detail.id_setting_grup_privilege AS id_setting_grup_privilege,\n				tb_setting_grup.name_grup AS name_grup,\n				tb_setting_grup_privilege_detail.code_permissions AS code_permissions \n			FROM\n				tb_setting_grup_privilege_detail\n				JOIN tb_setting_grup_privilege ON tb_setting_grup_privilege_detail.id_setting_grup_privilege = tb_setting_grup_privilege.id\n				JOIN tb_setting_grup ON tb_setting_grup_privilege.id_setting_grup = tb_setting_grup.id
client_cs_name=utf8mb4
connection_cl_name=utf8mb4_general_ci
view_body_utf8=select `db_receipt_golang_crud_v_2`.`tb_setting_grup_privilege_detail`.`id` AS `id`,`db_receipt_golang_crud_v_2`.`tb_setting_grup_privilege`.`id_setting_grup` AS `id_setting_grup`,`db_receipt_golang_crud_v_2`.`tb_setting_grup_privilege_detail`.`id_setting_grup_privilege` AS `id_setting_grup_privilege`,`db_receipt_golang_crud_v_2`.`tb_setting_grup`.`name_grup` AS `name_grup`,`db_receipt_golang_crud_v_2`.`tb_setting_grup_privilege_detail`.`code_permissions` AS `code_permissions` from ((`db_receipt_golang_crud_v_2`.`tb_setting_grup_privilege_detail` join `db_receipt_golang_crud_v_2`.`tb_setting_grup_privilege` on((`db_receipt_golang_crud_v_2`.`tb_setting_grup_privilege_detail`.`id_setting_grup_privilege` = `db_receipt_golang_crud_v_2`.`tb_setting_grup_privilege`.`id`))) join `db_receipt_golang_crud_v_2`.`tb_setting_grup` on((`db_receipt_golang_crud_v_2`.`tb_setting_grup_privilege`.`id_setting_grup` = `db_receipt_golang_crud_v_2`.`tb_setting_grup`.`id`)))
