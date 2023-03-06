# Took

抖声后端实现

## 模块划分

接口详细说明文档见：https://bytedance.feishu.cn/docs/doccnKrCsU5Iac6eftnFBdsXTof

### 用户模块

- /douyin/user/register/ - 用户注册接口

  注册新用户，需保证用户名唯一

- /douyin/user/login/ - 用户登录接口

  通过用户名和密码进行登录

- /douyin/user/ - 用户信息

  返回指定用户信息

- /douyin/relation/action/ - 关注操作

  登录用户对其它用户进行关注或取关

- /douyin/relation/follow/list/ - 用户关注列表

  返回登录用户关注的所有用户列表

- /douyin/relation/follower/list/ - 用户粉丝列表

  返回登录用户的粉丝列表

- /douyin/relation/friend/list/ - 用户好友列表

  返回与登录用户互关的用户列表



### 视频模块

- /douyin/feed/ - 视频流接口

  返回按投稿时间倒序的视频列表。

- /douyin/publish/action/ - 视频投稿

  登录用户上传视频

- /douyin/publidh/list/ - 发布列表

  返回登录用户的所有视频投稿

- /douyin/favorite/action/ - 赞操作

  登录用户对视频点赞或取消点赞

- /douyin/favorite/list/ - 喜欢列表

  登录用户的所有点赞视频列表

- /douyin/comment/action/ - 评论操作

  登录用户对视频进行评论

- /douyin/comment/list/ - 视频评论列表

  查看视频的所有评论，按发布时间倒序



### 聊天模块

- /douyin/message/list/ - 聊天记录

  返回当前登录用户与其他指定用户的聊天消息记录

- /douyin/message/action/ - 消息操作

  登录用户对消息的相关操作（目前只支持消息发送）

