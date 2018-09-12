
## 常用命令
```
go install  
go get -u ""  下载包
go fmt 统一代码风格和模版
go test package  运行当前包目录下tests
```

## plugin
github.com/julienschmidt/httprouter
github.com/go-sql-driver/mysql

## golang的test
```
最佳实践，不是必须的
main.go
main_test.go
```

## API
1. 统一接口
2. 无状态
3. 可缓存
4. 分层
5. cs模式

#### 设计原则
1. 以URL风格设计api
2. 通过不同的method(GET,POST,PUT,DELETE)区分资源的CRUD
3. 返回码符合http资源描述的规定

#### 用户
1. 创建用户： /user method:post, sc: 201 400 500
2. 登录：/user/:username method:post, sc:200 400 500
3. 获取用户的基本信息：/user/:username method:get,sc:200,400,401,403,500
4. 用户注销：/user/:username method:delete,sc: 204,400,401,403,500

#### 用户资源
1. list all videos : /user/:username/videos method:get, sc:200,400,500
2. get one video: /user/:username/videos/:vid-id method:get,sc:200,400,500
3. delete on video: /user/:username/videos/:vid-id method:delete,sc:204,400,401,403,500

#### 评论
1. show comments: /videos/:vid-id/comments method:get,sc:200,400,500
2. post a comment: /videos/:vid-id/comments method:post,sc:201,400,500
3. delete a comment:/videos/:vid-id/comments/:comment-id method:delete,sc:204,400,401,403,500

### db
```sql
create table users(
  id int not null auto_increment,
  login_name VARCHAR(64),
  pwd text,
  PRIMARY key(`id`),
  UNIQUE(`login_name`)
);

 create table video_info (
  `id` VARCHAR(64) not NULL,
  `author_id` INT,
  `name` text,
  `display_ctime` text,
  `create_time` datetime,
  PRIMARY key(`id`)
);

CREATE TABLE comments(
id VARCHAR(64) not NULL ,
video_id VARCHAR(64),
author_id int,
content text,
time datetime,
PRIMARY key(id)
);

create table sessions (
  session_id VARCHAR(100) not null,
  ttl tinytext,
  login_name varchar(64),
  PRIMARY key(session_id)
);


```

## Session

## streamserver 模块

#### bucket token 算法
限流

## scheduler 异步任务模块

1. restful api
2. Timer
3. 生产者/消费者模型


## Web 前端模块
