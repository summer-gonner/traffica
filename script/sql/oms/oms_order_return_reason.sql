create table oms_order_return_reason
(
    id          bigint auto_increment
        primary key,
    name        varchar(100) not null comment '退货类型',
    sort        int          not null comment '排序',
    status tinyint not null comment '状态：0->不启用；1->启用',
    create_time datetime     default CURRENT_TIMESTAMP not null comment '创建时间'
)
    comment '退货原因表';

INSERT INTO oms_order_return_reason (id, name, sort, status, create_time) VALUES (1, '质量问题', 1, 1, current_time);
INSERT INTO oms_order_return_reason (id, name, sort, status, create_time) VALUES (2, '尺码太大', 1, 1, current_time);
INSERT INTO oms_order_return_reason (id, name, sort, status, create_time) VALUES (3, '颜色不喜欢', 1, 1, current_time);
INSERT INTO oms_order_return_reason (id, name, sort, status, create_time) VALUES (4, '7天无理由退货', 1, 1, current_time);
INSERT INTO oms_order_return_reason (id, name, sort, status, create_time) VALUES (5, '价格问题', 1, 0, current_time);
INSERT INTO oms_order_return_reason (id, name, sort, status, create_time) VALUES (12, '发票问题', 0, 1, current_time);
INSERT INTO oms_order_return_reason (id, name, sort, status, create_time) VALUES (13, '其他问题', 0, 1, current_time);
INSERT INTO oms_order_return_reason (id, name, sort, status, create_time) VALUES (14, '物流问题', 0, 1, current_time);
INSERT INTO oms_order_return_reason (id, name, sort, status, create_time) VALUES (15, '售后问题', 0, 1, current_time);
