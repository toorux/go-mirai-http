package model

import (
	"encoding/json"
)

// MessageType 消息类型
type MessageType string

const (
	TypeFriendMessage      MessageType = "FriendMessage"      // 好友消息
	TypeGroupMessage       MessageType = "GroupMessage"       // 群组消息
	TypeTempMessage        MessageType = "TempMessage"        // 群临时消息
	TypeStrangerMessage    MessageType = "StrangerMessage"    // 陌生人消息
	TypeOtherClientMessage MessageType = "OtherClientMessage" // 其他客户端消息

	TypeFriendSyncMessage   MessageType = "FriendSyncMessage"   // 同步好友消息
	TypeGroupSyncMessage    MessageType = "GroupSyncMessage"    // 同步群组消息
	TypeTempSyncMessage     MessageType = "TempSyncMessage"     // 同步群临时消息
	TypeStrangerSyncMessage MessageType = "StrangerSyncMessage" // 同步陌生人消息

	TypeBotOnlineEvent         MessageType = "BotOnlineEvent"         // Bot登录成功
	TypeBotOfflineEventActive  MessageType = "BotOfflineEventActive"  // Bot主动离线
	TypeBotOfflineEventForce   MessageType = "BotOfflineEventForce"   // Bot被挤下线
	TypeBotOfflineEventDropped MessageType = "BotOfflineEventDropped" // Bot被服务器断开或因网络问题而掉线
	TypeBotReloginEvent        MessageType = "BotReloginEvent"        // Bot主动重新登录

	TypeFriendInputStatusChangedEvent MessageType = "FriendInputStatusChangedEvent" // 好友输入状态改变
	TypeFriendNickChangedEvent        MessageType = "FriendNickChangedEvent"        // 好友昵称改变

	TypeBotGroupPermissionChangeEvent MessageType = "BotGroupPermissionChangeEvent" // Bot在群里的权限被改变. 操作人一定是群主
	TypeBotMuteEvent                  MessageType = "BotMuteEvent"                  // Bot被禁言
	TypeBotUnmuteEvent                MessageType = "BotUnmuteEvent"                // Bot被取消禁言
	TypeBotJoinGroupEvent             MessageType = "BotJoinGroupEvent"             // Bot加入了一个新群
	TypeBotLeaveEventActive           MessageType = "BotLeaveEventActive"           // Bot主动退出一个群
	TypeBotLeaveEventKick             MessageType = "BotLeaveEventKick"             // Bot被踢出一个群

	TypeGroupRecallEvent                     MessageType = "GroupRecallEvent"                     // 群消息撤回
	TypeFriendRecallEvent                    MessageType = "FriendRecallEvent"                    // 好友消息撤回
	TypeNudgeEvent                           MessageType = "NudgeEvent"                           // 戳一戳
	TypeGroupNameChangeEvent                 MessageType = "GroupNameChangeEvent"                 // 某个群名改变
	TypeGroupEntranceAnnouncementChangeEvent MessageType = "GroupEntranceAnnouncementChangeEvent" // 某群入群公告改变
	TypeGroupMuteAllEvent                    MessageType = "GroupMuteAllEvent"                    // 全员禁言
	TypeGroupAllowAnonymousChatEvent         MessageType = "GroupAllowAnonymousChatEvent"         // 匿名聊天
	TypeGroupAllowConfessTalkEvent           MessageType = "GroupAllowConfessTalkEvent"           // 坦白说

	TypeGroupAllowMemberInviteEvent   MessageType = "GroupAllowMemberInviteEvent"   // 允许群员邀请好友加群
	TypeMemberJoinEvent               MessageType = "MemberJoinEvent"               // 新人入群的事件
	TypeMemberLeaveEventKick          MessageType = "MemberLeaveEventKick"          // 成员被踢出群（该成员不是Bot）
	TypeMemberLeaveEventQuit          MessageType = "MemberLeaveEventQuit"          // 成员主动离群（该成员不是Bot）
	TypeMemberCardChangeEvent         MessageType = "MemberCardChangeEvent"         // 群名片改动
	TypeMemberSpecialTitleChangeEvent MessageType = "MemberSpecialTitleChangeEvent" // 群头衔改动
	TypeMemberPermissionChangeEvent   MessageType = "MemberPermissionChangeEvent"   // 成员权限改变的事件（该成员不是Bot）
	TypeMemberMuteEvent               MessageType = "MemberMuteEvent"               // 群成员被禁言事件（该成员不是Bot）
	TypeMemberUnmuteEvent             MessageType = "MemberUnmuteEvent"             // 群成员被取消禁言事件（该成员不是Bot）
	TypeMemberHonorChangeEvent        MessageType = "MemberHonorChangeEvent"        // 群员称号改变

	TypeNewFriendRequestEvent           MessageType = "NewFriendRequestEvent"           // 添加好友申请
	TypeMemberJoinRequestEvent          MessageType = "MemberJoinRequestEvent"          // 用户入群申请（Bot需要有管理员权限）
	TypeBotInvitedJoinGroupRequestEvent MessageType = "BotInvitedJoinGroupRequestEvent" // Bot被邀请入群申请

	TypeOtherClientOnlineEvent  MessageType = "OtherClientOnlineEvent"  // 其他客户端上线
	TypeOtherClientOfflineEvent MessageType = "OtherClientOfflineEvent" // 其他客户端下线

	TypeCommandExecutedEvent MessageType = "CommandExecutedEvent" // 命令被执行
)

// Message 消息类型
type Message struct {
	Type         MessageType
	Sender       Sender         `json:"sender,omitempty"`
	Subject      Subject        `json:"subject,omitempty"`
	MessageChain []MessageChain `json:"messageChain,omitempty"`

	Qq uint64 `json:"qq,omitempty"`

	Friend    Friend `json:"friend,omitempty"`
	Inputting bool   `json:"inputting,omitempty"`

	From string `json:"from,omitempty"`
	To   string `json:"to,omitempty"`

	Origin  string `json:"origin,omitempty"`
	Current string `json:"current,omitempty"`
	Group   Group  `json:"group,omitempty"`

	DurationSeconds int `json:"durationSeconds,omitempty"` //禁言时长，单位为秒
	Operator        any `json:"operator,omitempty"`        //操作的管理员或群主信息

	Invitor Operator `json:"invitor,omitempty"`

	AuthorId  uint64 `json:"authorId,omitempty"`
	MessageId int    `json:"messageId,omitempty"`
	Time      int    `json:"time,omitempty"`

	FromId uint64 `json:"fromId,omitempty"`
	Action string `json:"action,omitempty"`
	Suffix string `json:"suffix,omitempty"`
	Target uint64 `json:"target,omitempty"`

	IsByBot bool `json:"isByBot,omitempty"`

	Member Operator `json:"member,omitempty"`

	Honor string `json:"honor,omitempty"`

	EventId int    `json:"eventId,omitempty"`
	GroupId uint64 `json:"groupId,omitempty"`
	Nick    string `json:"nick,omitempty"`
	Message string `json:"message,omitempty"`

	GroupName string `json:"groupName,omitempty"`

	Client Client `json:"client,omitempty"`
	Kind   int    `json:"kind,omitempty"`

	Name string    `json:"name,omitempty"`
	Args []CmdArgs `json:"args,omitempty"`
}

func (m Message) GetOperator() Operator {
	if m.Type == TypeFriendRecallEvent {
		return Operator{}
	} else {
		str, err := json.Marshal(m.Operator)
		if err != nil {
			return Operator{}
		}
		var oper Operator
		err2 := json.Unmarshal(str, &oper)
		if err2 != nil {
			return Operator{}
		}
		return oper
	}
}

func (m Message) GetOperatorFriend() uint64 {
	if m.Type != TypeFriendRecallEvent {
		return 0
	} else {
		v, ok := m.Operator.(uint64)
		if ok {
			return v
		} else {
			return 0
		}
	}
}

// Sender 消息发送目标类型
type Sender struct {
	Id uint64 `json:"id,omitempty"`

	// 好友消息 和 陌生人消息
	Nickname string `json:"nickname,omitempty"`
	Remark   string `json:"remark,omitempty"`

	// 群消息 和 群临时会话
	MemberName         string `json:"memberName,omitempty"`
	SpecialTitle       string `json:"specialTitle,omitempty"`
	Permission         string `json:"permission,omitempty"`
	JoinTimestamp      int    `json:"joinTimestamp,omitempty"`
	LastSpeakTimestamp int    `json:"lastSpeakTimestamp,omitempty"`
	MuteTimeRemaining  int    `json:"muteTimeRemaining,omitempty"`
	Group              Group  `json:"group,omitempty"`

	// 其他客户端消息
	Platform string `json:"platform,omitempty"`
}

// Subject 消息发送目标类型
type Subject struct {
	Id uint64 `json:"id,omitempty"`

	// 好友消息 和 陌生人消息
	Nickname string `json:"nickname,omitempty"`
	Remark   string `json:"remark,omitempty"`

	// 群消息 和 群临时会话
	MemberName         string `json:"memberName,omitempty"`
	SpecialTitle       string `json:"specialTitle,omitempty"`
	Permission         string `json:"permission,omitempty"`
	JoinTimestamp      int    `json:"joinTimestamp,omitempty"`
	LastSpeakTimestamp int    `json:"lastSpeakTimestamp,omitempty"`
	MuteTimeRemaining  int    `json:"muteTimeRemaining,omitempty"`
	Group              Group  `json:"group,omitempty"`
	Kind               string `json:"kind,omitempty"`
}

// Group 消息发送目标来自的群信息
type Group struct {
	Id         uint64 `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Permission string `json:"permission,omitempty"`
}

type Friend struct {
	Id       uint64 `json:"id,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Remark   string `json:"remark,omitempty"`
}

type Operator struct {
	Id                 uint64 `json:"id,omitempty"`
	MemberName         string `json:"memberName,omitempty"`
	Permission         string `json:"permission,omitempty"`
	SpecialTitle       string `json:"specialTitle,omitempty"`
	JoinTimestamp      int    `json:"joinTimestamp,omitempty"`
	LastSpeakTimestamp int    `json:"lastSpeakTimestamp,omitempty"`
	MuteTimeRemaining  int    `json:"muteTimeRemaining,omitempty"`
	Group              Group  `json:"group,omitempty"`
}

type Client struct {
	Id       int    `json:"id,omitempty"`
	Platform string `json:"platform,omitempty"`
}

type CmdArgs struct {
	Type string `json:"type,omitempty"`
	Text string `json:"text,omitempty"`
}
