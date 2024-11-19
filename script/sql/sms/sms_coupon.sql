create table sms_coupon
(
    id            bigint auto_increment
        primary key,
    type         tinyint not null comment '优惠券类型；0->全场赠券；1->会员赠券；2->购物赠券；3->注册赠券',
    name          varchar(100)   not null comment '名称',
    platform     tinyint not null comment '使用平台：0->全部；1->移动；2->PC',
    count         int            not null comment '数量',
    amount       bigint  not null comment '金额',
    per_limit     int            not null comment '每人限领张数',
    min_point    bigint  not null comment '使用门槛；0表示无门槛',
    start_time    datetime       not null comment '开始时间',
    end_time      datetime       not null comment '结束时间',
    use_type     tinyint not null comment '使用类型：0->全场通用；1->指定分类；2->指定商品',
    note          varchar(200)   not null comment '备注',
    publish_count int            not null comment '发行数量',
    use_count     int            not null comment '已使用数量',
    receive_count int            not null comment '领取数量',
    enable_time   datetime       not null comment '可以领取的日期',
    code          varchar(64)    not null comment '优惠码',
    member_level tinyint not null comment '可领取的会员类型：0->无限时'
)
    comment '优惠券表';

INSERT INTO sms_coupon (id, type, name, platform, count, amount, per_limit, min_point, start_time, end_time, use_type, note, publish_count, use_count, receive_count, enable_time, code, member_level) VALUES (2, 0, '全品类通用券', 0, 92, 10.00, 1, 100.00, current_time, '2025-12-04 10:14:28', 0, '满100减10', 100, 0, 8, current_time, 'string', 0);
INSERT INTO sms_coupon (id, type, name, platform, count, amount, per_limit, min_point, start_time, end_time, use_type, note, publish_count, use_count, receive_count, enable_time, code, member_level) VALUES (3, 0, '小米手机专用券', 0, 92, 50.00, 1, 1000.00, current_time, '2025-12-04 10:14:28', 2, '小米手机专用优惠券', 100, 0, 8, current_time, 'string', 0);
INSERT INTO sms_coupon (id, type, name, platform, count, amount, per_limit, min_point, start_time, end_time, use_type, note, publish_count, use_count, receive_count, enable_time, code, member_level) VALUES (4, 0, '手机品类专用券', 0, 92, 300.00, 1, 2000.00, current_time, '2025-12-04 10:14:28', 1, '手机分类专用优惠券', 100, 0, 8, current_time, 'string', 0);
INSERT INTO sms_coupon (id, type, name, platform, count, amount, per_limit, min_point, start_time, end_time, use_type, note, publish_count, use_count, receive_count, enable_time, code, member_level) VALUES (7, 0, 'T恤分类专用优惠券', 0, 93, 50.00, 1, 500.00, current_time, '2025-12-04 10:14:28', 1, '满500减50', 100, 0, 7, current_time, 'string', 0);
INSERT INTO sms_coupon (id, type, name, platform, count, amount, per_limit, min_point, start_time, end_time, use_type, note, publish_count, use_count, receive_count, enable_time, code, member_level) VALUES (8, 0, '新优惠券', 0, 100, 100.00, 1, 1000.00, current_time, '2025-12-04 10:14:28', 0, '测试', 100, 0, 1, current_time, 'string', 0);
INSERT INTO sms_coupon (id, type, name, platform, count, amount, per_limit, min_point, start_time, end_time, use_type, note, publish_count, use_count, receive_count, enable_time, code, member_level) VALUES (9, 0, '全品类通用券', 0, 100, 5.00, 1, 100.00, current_time, '2025-12-04 10:14:28', 0, '测试', 100, 0, 1, current_time, 'string', 0);
INSERT INTO sms_coupon (id, type, name, platform, count, amount, per_limit, min_point, start_time, end_time, use_type, note, publish_count, use_count, receive_count, enable_time, code, member_level) VALUES (10, 0, '全品类通用券', 0, 100, 15.00, 1, 200.00, current_time, '2025-12-04 10:14:28', 0, '测试', 100, 0, 1, current_time, 'string', 0);
INSERT INTO sms_coupon (id, type, name, platform, count, amount, per_limit, min_point, start_time, end_time, use_type, note, publish_count, use_count, receive_count, enable_time, code, member_level) VALUES (11, 0, '全品类通用券', 0, 1000, 50.00, 1, 1000.00, current_time, '2025-12-04 10:14:28', 0, '测试', 1000, 0, 0, current_time, 'string', 0);
INSERT INTO sms_coupon (id, type, name, platform, count, amount, per_limit, min_point, start_time, end_time, use_type, note, publish_count, use_count, receive_count, enable_time, code, member_level) VALUES (12, 0, '移动端全品类通用券', 1, 1, 10.00, 1, 100.00, current_time, '2025-12-04 10:14:28', 0, '测试', 100, 0, 0, current_time, 'string', 0);
INSERT INTO sms_coupon (id, type, name, platform, count, amount, per_limit, min_point, start_time, end_time, use_type, note, publish_count, use_count, receive_count, enable_time, code, member_level) VALUES (19, 0, '手机分类专用', 0, 100, 100.00, 1, 1000.00, current_time, '2025-12-04 10:14:28', 1, '手机分类专用', 100, 0, 0, current_time, 'string', 0);
INSERT INTO sms_coupon (id, type, name, platform, count, amount, per_limit, min_point, start_time, end_time, use_type, note, publish_count, use_count, receive_count, enable_time, code, member_level) VALUES (20, 0, '小米手机专用', 0, 100, 200.00, 1, 1000.00, current_time, '2025-12-04 10:14:28', 2, '小米手机专用', 100, 0, 0, current_time, 'string', 0);
INSERT INTO sms_coupon (id, type, name, platform, count, amount, per_limit, min_point, start_time, end_time, use_type, note, publish_count, use_count, receive_count, enable_time, code, member_level) VALUES (21, 0, 'xxx11', 0, 100, 10.02, 1, 100.00, current_time, '2025-12-04 10:14:28', 2, '测试', 100, 0, 0, current_time, '0000', 0);
INSERT INTO sms_coupon (id, type, name, platform, count, amount, per_limit, min_point, start_time, end_time, use_type, note, publish_count, use_count, receive_count, enable_time, code, member_level) VALUES (23, 0, '有效期测试', 0, 100, 10.00, 1, 100.00, current_time, '2025-12-04 10:14:28', 0, '测试', 100, 0, 0, current_time, 'string', 0);
INSERT INTO sms_coupon (id, type, name, platform, count, amount, per_limit, min_point, start_time, end_time, use_type, note, publish_count, use_count, receive_count, enable_time, code, member_level) VALUES (25, 0, '小米测试', 0, 10, 100.00, 100, 10.00, current_time, '2025-12-04 10:14:28', 0, '123', 10, 0, 0, current_time, '213', 0);
