# simple-tiktok

## 抖音项目服务端

具体功能内容参考 app 说明文档

[app说明文档](https://bytedance.feishu.cn/docs/doccnM9KkBAdyDhg8qaeGlIz7S7)

[接口文档](https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18900033)

工程无其他依赖，直接编译运行即可

```shell
go build && ./simple-demo
```

### 功能说明

| 功能项 | 说明 | 备注 |
| ------ | ------ | --- |
| 视频 Feed 流 | 支持所有用户刷抖音，按投稿时间倒序推出 | <input type="checkbox" checked=true> |
| 视频投稿 | 登录用户可以自己拍视频投稿 | <input type="checkbox" checked=true> |
| 个人信息 | 查看自己的基本信息和投稿列表，注册用户流程简化 | <input type="checkbox" checked=true>
| 点赞列表 | 在个人主页能够查看点赞视频列表 | <input type="checkbox">
| 用户评论 | 登录用户可以对视频点赞，并在视频下进行评论 | <input type="checkbox">
| 关注列表、粉丝列表 | 登录用户可以关注其他用户，能够在个人信息页查看本人的关注数和粉丝数，点击打开关注列表和粉丝列表。 | <input type="checkbox" checked=true>

* 用户登录数据保存在 mysql 中，使用前请提前配置数据库。
* 视频上传后会保存到本地 public 目录中，访问时用 127.0.0.1:8080/static/video_name 即可
* 安装 app 后，请打开首页，在右下角双击 “**我**” ，配置后端路径。

### 测试数据

测试数据写在 demo_data.go 中，用于列表接口的 mock 测试