create table sys_dept
(
    id          bigint auto_increment comment '编号'
        primary key,
    dept_name   varchar(50)                            not null comment '部门名称',
    dept_status tinyint      default 1                 not null comment '部门状态',
    dept_sort   tinyint      default 0                 not null comment '部门排序',
    parent_id   bigint                                 not null comment '上级机构ID，一级机构为0',
    leader      varchar(64)  default ''                not null comment '负责人',
    phone       varchar(11)  default ''                not null comment '电话号码',
    email       varchar(24)  default ''                not null comment '邮箱',
    remark      varchar(255) default ''                not null comment '备注信息',
    is_deleted  tinyint      default 0                 not null comment '是否删除  0：否  1：是',
    parent_ids  varchar(64)                            not null comment '上级机构IDs，一级机构为0',
    create_by   varchar(50)                            not null comment '创建者',
    create_time datetime     default CURRENT_TIMESTAMP not null comment '创建时间',
    update_by   varchar(50)  default ''                not null comment '更新者',
    update_time datetime                               null on update CURRENT_TIMESTAMP comment '更新时间'
)
    comment '部门信息表';

INSERT INTO sys_dept (id, dept_name, dept_status, dept_sort, parent_id, leader, phone, email, remark, is_deleted, parent_ids, create_by, create_time, update_by, update_time) VALUES (1, '总公司一', 1, 1, 0, 'liufeihua', '18613030352', '1002219331@qq.com', '测试', 0, '0', 'admin', '2024-05-30 17:16:46', '', null);
INSERT INTO sys_dept (id, dept_name, dept_status, dept_sort, parent_id, leader, phone, email, remark, is_deleted, parent_ids, create_by, create_time, update_by, update_time) VALUES (2, '电商部', 1, 1, 1, 'liufeihua', '18613030352', '1002219331@qq.com', '测试', 0, '0', 'admin', '2024-05-30 17:16:46', '', null);
INSERT INTO sys_dept (id, dept_name, dept_status, dept_sort, parent_id, leader, phone, email, remark, is_deleted, parent_ids, create_by, create_time, update_by, update_time) VALUES (4, '开发一', 1, 1, 2, 'liufeihua', '18613030352', '1002219331@qq.com', '测试', 0, '0', 'admin', '2024-05-30 17:16:46', '', null);
INSERT INTO sys_dept (id, dept_name, dept_status, dept_sort, parent_id, leader, phone, email, remark, is_deleted, parent_ids, create_by, create_time, update_by, update_time) VALUES (6, '开发二', 1, 1, 2, 'liufeihua', '18613030352', '1002219331@qq.com', '测试', 0, '0', 'admin', '2024-05-30 17:16:46', '', null);
INSERT INTO sys_dept (id, dept_name, dept_status, dept_sort, parent_id, leader, phone, email, remark, is_deleted, parent_ids, create_by, create_time, update_by, update_time) VALUES (3, '运营部', 1, 1, 1, 'liufeihua', '18613030352', '1002219331@qq.com', '测试', 0, '0', 'admin', '2024-05-30 17:16:46', '', null);
INSERT INTO sys_dept (id, dept_name, dept_status, dept_sort, parent_id, leader, phone, email, remark, is_deleted, parent_ids, create_by, create_time, update_by, update_time) VALUES (5, '运营部一', 1, 1, 3, 'liufeihua', '18613030352', '1002219331@qq.com', '测试', 0, '0', 'admin', '2024-05-30 17:16:46', '', null);
INSERT INTO sys_dept (id, dept_name, dept_status, dept_sort, parent_id, leader, phone, email, remark, is_deleted, parent_ids, create_by, create_time, update_by, update_time) VALUES (7, '财务部门', 1, 1, 1, 'liufeihua', '18613030352', '1002219331@qq.com', '测试', 0, '0', 'admin', '2024-05-30 17:16:46', '', null);
INSERT INTO sys_dept (id, dept_name, dept_status, dept_sort, parent_id, leader, phone, email, remark, is_deleted, parent_ids, create_by, create_time, update_by, update_time) VALUES (8, 'test', 1, 1, 0, 'liufeihua', '18613030352', '1002219331@qq.com', '测试', 0, '0', 'admin', '2024-05-30 17:16:46', '', null);
