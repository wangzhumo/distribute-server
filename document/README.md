## Apk分发

### 数据

#### APK
```shell
create table apkinfo(
    id int auto_increment comment 'version发布的Id，主键自增',
    apk_icon varchar(255) comment 'icon图片',
    apk_name varchar(20) comment 'apk的名字',
    appid varchar(50) not null comment 'apk的包名',
    last_release int comment 'apk的最后一个Release版本',
    primary key (id)
) CHARACTER SET utf8 COLLATE utf8_general_ci comment 'Apk发布版本';
```


#### 版本

```shell
create table versions(
    id int auto_increment comment 'version发布的Id，主键自增',
    name varchar(20) not null comment '当前Apk的版本名',
    code int default -1 not null comment '当前Apk的VersionCode',
    changelog varchar(255) comment '更改内容日志',
    downloads int default 0 not null comment '下载次数',
    apk_url varchar(255) comment '下载地址',
    qrcode_url varchar(255) comment '下载地址的Qrcode',
    apkid int not null comment 'apk的ID',
    release boolean default false comment '是否是一个Release版本',
    primary key (id)
) CHARACTER SET utf8 COLLATE utf8_general_ci comment 'Apk发布版本';

```


### 接口

#### 返回数据

##### 普通接口
```json
{
    "code":0,
    "message":"",
    "data":{}
}
```
- code  0 成功  
- message  失败时一般返回错误原因
- data  具体返回数据

##### 分页接口
```json
{
    "last_id":0,
    "more":false,
    "data":[],
    "total":123
}
```


#### 1.上传接口

@POST
> v1/apk/upload

@Params
> file


#### 2.获取版本列表
@GET
> vl/info/list

@Params
> size  `int` 每页数量
> 
> last_id  `long`
>   
> type  `string`  release 正式包  debug(默认值) debug

@Ruturn
```json
{
    "last_id":0,
    "more":false,
    "data":[

    ],
    "total":213
}

```
#### 3.获取详细信息
@GET
> v1/info/apkinfo

@Params
> id `int` apk的id

@Return
```json
{
    "code":0,
    "message":"",
    "data":{
        "id":0,
        "code":0,
        "changelog":"",
        "downloads":"",
        "apk_url":"",
        "qrcode_url":"",
        "app_icon":"",
        "release":false
    }
}
```

#### 4.标记release
@GET
> v1/info/mark

@Params
> id `int` apk的id
> release `bool` 

@Return
```json
{
    "code":0,
    "message":"",
    "data":{
        "id":0,
        "code":0,
        "changelog":"",
        "downloads":"",
        "apk_url":"",
        "qrcode_url":"",
        "app_icon":"",
        "release":false
    }
}
```

#### 5.创建应用
@GET
> v1/apk/create

@Params
> name `string` apk的名字
> 
> icon `file`  apk的icon文件

@Return
```json
{
    "code":0,
    "message":"ok",
    "data":{}
}
```