TYPE=VIEW
query=select `db_receipt_golang_crud_v_2`.`tb_setting_grup_privilege`.`id` AS `id`,`db_receipt_golang_crud_v_2`.`tb_setting_grup_privilege`.`id_setting_grup` AS `id_setting_grup`,`db_receipt_golang_crud_v_2`.`tb_setting_grup`.`name_grup` AS `name_grup`,`db_receipt_golang_crud_v_2`.`tb_setting_grup_privilege`.`remarks` AS `remarks`,`db_receipt_golang_crud_v_2`.`tb_setting_grup_privilege`.`status` AS `STATUS`,`db_receipt_golang_crud_v_2`.`tb_setting_grup_privilege`.`created_at` AS `created_at`,`db_receipt_golang_crud_v_2`.`tb_setting_grup_privilege`.`updated_at` AS `updated_at` from (`db_receipt_golang_crud_v_2`.`tb_setting_grup_privilege` join `db_receipt_golang_crud_v_2`.`tb_setting_grup` on((`db_receipt_golang_crud_v_2`.`tb_setting_grup_privilege`.`id_setting_grup` = `db_receipt_golang_crud_v_2`.`tb_setting_grup`.`id`)))
md5=ed0aee85b2115da0db752d845c25cbeb
updatable=1
algorithm=0
definer_user=user
definer_host=%
suid=1
with_check_option=0
timestamp=2020-11-25 11:05:40
create-version=1
source=SELECT\n				tb_setting_grup_privilege.id AS id,\n				tb_setting_grup_privilege.id_setting_grup AS id_setting_grup,\n				tb_setting_grup.name_grup AS name_grup,\n				tb_setting_grup_privilege.remarks AS remarks,\n				tb_setting_grup_privilege.STATUS AS STATUS,\n				tb_setting_grup_privilege.created_at AS created_at,\n				tb_setting_grup_privilege.updated_at AS updated_at \n			FROM\n				tb_setting_grup_privilege\n				JOIN tb_setting_grup ON tb_setting_grup_privilege.id_setting_grup = tb_setting_grup.id
client_cs_name=utf8mb4
connection_cl_name=utf8mb4_general_ci
view_body_utf8=select `db_receipt_golang_crud_v_2`.`tb_setting_grup_privilege`.`id` AS `id`,`db_receipt_golang_crud_v_2`.`tb_setting_grup_privilege`.`id_setting_grup` AS `id_setting_grup`,`db_receipt_golang_crud_v_2`.`tb_setting_grup`.`name_grup` AS `name_grup`,`db_receipt_golang_crud_v_2`.`tb_setting_grup_privilege`.`remarks` AS `remarks`,`db_receipt_golang_crud_v_2`.`tb_setting_grup_privilege`.`status` AS `STATUS`,`db_receipt_golang_crud_v_2`.`tb_setting_grup_privilege`.`created_at` AS `created_at`,`db_receipt_golang_crud_v_2`.`tb_setting_grup_privilege`.`updated_at` AS `updated_at` from (`db_receipt_golang_crud_v_2`.`tb_setting_grup_privilege` join `db_receipt_golang_crud_v_2`.`tb_setting_grup` on((`db_receipt_golang_crud_v_2`.`tb_setting_grup_privilege`.`id_setting_grup` = `db_receipt_golang_crud_v_2`.`tb_setting_grup`.`id`)))
