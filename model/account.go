package model

type Profile struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	Level    int    `json:"level"`
	Sign     string `json:"sign"`
	Sex      string `json:"sex"`
}

// Group 消息发送目标来自的群信息
type Group struct {
	Id         int    `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Permission string `json:"permission,omitempty"`
}

type Friend struct {
	Id       int    `json:"id,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Remark   string `json:"remark,omitempty"`
}

type GroupOrFriend struct {
	Group
	Friend
}

type Member struct {
	Id                 int    `json:"id,omitempty"`
	MemberName         string `json:"memberName,omitempty"`
	Permission         string `json:"permission,omitempty"`
	SpecialTitle       string `json:"specialTitle,omitempty"`
	JoinTimestamp      int    `json:"joinTimestamp,omitempty"`
	LastSpeakTimestamp int    `json:"lastSpeakTimestamp,omitempty"`
	MuteTimeRemaining  int    `json:"muteTimeRemaining,omitempty"`
	Group              Group  `json:"group,omitempty"`
}
