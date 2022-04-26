package model

// Anno 群公告
type Anno struct {
	Group                 Group  `json:"group"`
	Content               string `json:"content"`               // 内容
	SenderId              int    `json:"senderId"`              // 发布者账号
	Fid                   string `json:"fid"`                   //公告唯一id
	AllConfirmed          bool   `json:"allConfirmed"`          // 是否所有群成员已确认
	ConfirmedMembersCount int    `json:"confirmedMembersCount"` // 确认群成员人数
	PublicationTime       int    `json:"publicationTime"`       // 发布时间
}
