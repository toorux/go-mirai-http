package model

type GroupConfig struct {
	Name              string `json:"name"`              // 群名
	Announcement      string `json:"announcement"`      // 群公告
	ConfessTalk       bool   `json:"confessTalk"`       // 是否开启坦白说
	AllowMemberInvite bool   `json:"allowMemberInvite"` // 是否允许群员邀请
	AutoApprove       bool   `json:"autoApprove"`       // 是否开启自动审批入群
	AnonymousChat     bool   `json:"anonymousChat"`     // 是否允许匿名聊天
}
