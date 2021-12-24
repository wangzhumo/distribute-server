create table apkinfo
(
    id           int auto_increment comment 'version发布的Id，主键自增'
        primary key,
    apk_icon     varchar(255) null comment 'icon图片',
    apk_name     varchar(20)  null comment 'apk的名字',
    app_id       varchar(50)  not null comment 'apk的包名',
    last_release int          null comment 'apk的最后一个Release版本'
)
    comment 'Apk发布版本';

create table versions
(
    id              int auto_increment comment 'version发布的Id，主键自增'
        primary key,
    name            varchar(20)           not null comment '当前Apk的版本名',
    code            int        default -1 not null comment '当前Apk的VersionCode',
    changelog       varchar(255)          null comment '更改内容日志',
    downloads       int        default 0  not null comment '下载次数',
    apk_url         varchar(255)          null comment '下载地址',
    qrcode_url      varchar(255)          null comment '下载地址的Qrcode',
    apk_id          int                   not null comment 'apk的ID',
    release_version tinyint(1) default 0  null comment '是否是一个Release版本',
    timestamp       bigint     default 0  not null
)
    comment 'Apk发布版本';


