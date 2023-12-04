create table ums_member_login_log
(
    id          bigint auto_increment
        primary key,
    member_id   bigint      not null comment '会员id',
    create_time datetime    not null comment '登录时间',
    ip          varchar(64) not null comment '登录ip',
    city        varchar(64) not null comment '登录城市',
    login_type  int(1)      not null comment '登录类型：0->PC；1->android;2->ios;3->小程序',
    province    varchar(64) not null comment '登录省份'
)
    comment '会员登录记录';

