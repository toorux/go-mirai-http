# Go Mirai Http

自己使用的Golang实现与 mirai-app-http对接封装（目前仅实现http协议），数据格式请求方式几本与官方文档一致， 开发中...

[mirai-app-http http-adapter 文档](https://docs.mirai.mamoe.net/mirai-api-http/adapter/HttpAdapter.html#http-adapter)  
[mirai-app-http 事件类型 文档](https://docs.mirai.mamoe.net/mirai-api-http/api/EventType.html#%E7%9B%AE%E5%BD%95)  
[mirai-app-http 消息类型 文档](https://docs.mirai.mamoe.net/mirai-api-http/api/MessageType.html#%E6%B6%88%E6%81%AF%E9%93%BE%E7%B1%BB%E5%9E%8B)  

## 许可证
```
Copyright (C) 2019-2022 Mamoe Technologies and contributors.

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as
published by the Free Software Foundation, either version 3 of the
License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
```
本项目同mirai主项目，采用 AGPLv3 协议开源。


## TodoList
<details>
 
### 认证与会话
 - [x] 认证
 - [x] 绑定
 - [x] 释放

### 接收消息与事件
 - [x] 查看队列大小
 - [x] 获取队列头部
 - [x] 获取队列尾部
 - [x] 查看队列头部
 - [x] 查看队列尾部
 - [x] 所有消息与事件数据结构体
### 获取插件信息
 - [x] 获取插件信息
### 缓存操作
 - [x] 通过messageId获取消息
### 获取账号信息
 - [x] 获取好友列表
 - [x] 获取群列表
 - [x] 获取群成员列表
 - [x] 获取Bot资料
 - [x] 获取好友资料
 - [x] 获取群成员资料
 - [x] 获取QQ用户资料
### 消息发送与撤回
 - [x] 发送好友消息
 - [x] 发送群消息
 - [x] 发送临时会话消息
 - [x] 发送头像戳一戳消息
 - [x] 撤回消息
### 文件操作
 - [ ] 查看文件列表
 - [ ] 获取文件信息
 - [ ] 创建文件夹
 - [ ] 删除文件
 - [ ] 移动文件
 - [ ] 重命名文件
### 多媒体内容上传
 - [ ] 图片文件上传
 - [ ] 语音文件上传
 - [ ] 群文件上传
### 账号管理
 - [ ] 删除好友
### 群管理
 - [ ] 禁言群成员
 - [ ] 解除群成员禁言
 - [ ] 移除群成员
 - [ ] 退出群聊
 - [ ] 全体禁言
 - [ ] 解除全体禁言
 - [ ] 设置群精华消息
 - [ ] 获取群设置
 - [ ] 修改群设置
 - [ ] 获取群员设置
 - [ ] 修改群员设置
 - [ ] 修改群员管理员
### 群公告
 - [ ] 获取群公告
 - [ ] 发布群公告
 - [ ] 删除群公告
### 事件处理
 - [ ] 添加好友申请
 - [ ] 用户入群申请
 - [ ] Bot被邀请入群申请
 
 </details>
