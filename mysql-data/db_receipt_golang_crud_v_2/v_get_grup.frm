TYPE=VIEW
query=select `db_receipt_golang_crud_v_2`.`tb_setting_grup`.`id` AS `id`,`db_receipt_golang_crud_v_2`.`tb_setting_grup_privilege`.`id` AS `id_setting_grup_privilege`,`db_receipt_golang_crud_v_2`.`tb_setting_grup`.`name_grup` AS `name_grup`,`db_receipt_golang_crud_v_2`.`tb_setting_grup`.`status` AS `status` from (`db_receipt_golang_crud_v_2`.`tb_setting_grup` left join `db_receipt_golang_crud_v_2`.`tb_setting_grup_privilege` on((`db_receipt_golang_crud_v_2`.`tb_setting_grup`.`id` = `db_receipt_golang_crud_v_2`.`tb_setting_grup_privilege`.`id_setting_grup`)))
md5=d24c8ac80e1b3269a589cfbceed2237e
updatable=0
algorithm=0
definer_user=user
definer_host=%
suid=1
with_check_option=0
timestamp=2020-11-25 11:05:40
create-version=1
source=SELECT\n				tb_setting_grup.id AS id,\n				tb_setting_grup_privilege.id AS id_setting_grup_privilege,\n				tb_setting_grup.name_grup AS name_grup,\n				tb_setting_grup.status AS status \n			FROM\n	 			tb_setting_grup LEFT JOIN tb_setting_grup_privilege ON   tb_setting_grup.id = tb_setting_grup_privilege.id_setting_grup
client_cs_name=utf8mb4
connection_cl_name=utf8mb4_general_ci
view_body_utf8=select `db_receipt_golang_crud_v_2`.`tb_setting_grup`.`id` AS `id`,`db_receipt_golang_crud_v_2`.`tb_setting_grup_privilege`.`id` AS `id_setting_grup_privilege`,`db_receipt_golang_crud_v_2`.`tb_setting_grup`.`name_grup` AS `name_grup`,`db_receipt_golang_crud_v_2`.`tb_setting_grup`.`status` AS `status` from (`db_receipt_golang_crud_v_2`.`tb_setting_grup` left join `db_receipt_golang_crud_v_2`.`tb_setting_grup_privilege` on((`db_receipt_golang_crud_v_2`.`tb_setting_grup`.`id` = `db_receipt_golang_crud_v_2`.`tb_setting_grup_privilege`.`id_setting_grup`)))
