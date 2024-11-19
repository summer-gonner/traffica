create table sms_coupon_history
(
    id              bigint auto_increment
        primary key,
    coupon_id       bigint       not null comment '优惠券id',
    member_id       bigint       not null comment '会员id',
    coupon_code     varchar(64)  not null comment '优惠码',
    member_nickname varchar(64)  not null comment '领取人昵称',
    get_type   tinyint not null comment '获取类型：0->后台赠送；1->主动获取',
    create_time     datetime     not null comment '领取时间',
    use_status tinyint not null comment '使用状态：0->未使用；1->已使用；2->已过期',
    use_time        datetime     not null comment '使用时间',
    order_id        bigint       not null comment '订单编号',
    order_sn        varchar(100) not null comment '订单号码'
)
    comment '优惠券使用、领取历史表';

create index idx_coupon_id
    on sms_coupon_history (coupon_id);

create index idx_member_id
    on sms_coupon_history (member_id);

INSERT INTO sms_coupon_history (id, coupon_id, member_id, coupon_code, member_nickname, get_type, create_time, use_status, use_time, order_id, order_sn) VALUES (2, 2, 1, '4931048380330002', 'koobe', 1, current_time, 1, current_time, 12, '201809150101000001');
INSERT INTO sms_coupon_history (id, coupon_id, member_id, coupon_code, member_nickname, get_type, create_time, use_status, use_time, order_id, order_sn) VALUES (3, 3, 1, '4931048380330003', 'koobe', 1, current_time, 0, current_time, 12, ' ');
INSERT INTO sms_coupon_history (id, coupon_id, member_id, coupon_code, member_nickname, get_type, create_time, use_status, use_time, order_id, order_sn) VALUES (4, 4, 1, '4931048380330004', 'koobe', 1, current_time, 0, current_time, 12, ' ');
