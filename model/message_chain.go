package model

import (
	"encoding/json"
	"fmt"
	"github.com/kgysf/go-mirai-http/internal/pkg"
	"strings"
)

// MessageChainType 同步消息类型
type MessageChainType string

const (
	MsgChainTypeSource         MessageChainType = "Source"         //
	MsgChainTypeQuote          MessageChainType = "Quote"          //
	MsgChainTypeAt             MessageChainType = "At"             //
	MsgChainTypeAtAll          MessageChainType = "AtAll"          //
	MsgChainTypeFace           MessageChainType = "Face"           //
	MsgChainTypePlain          MessageChainType = "Plain"          //
	MsgChainTypeImage          MessageChainType = "Image"          //
	MsgChainTypeFlashImage     MessageChainType = "FlashImage"     //
	MsgChainTypeVoice          MessageChainType = "Voice"          //
	MsgChainTypeXml            MessageChainType = "Xml"            //
	MsgChainTypeJson           MessageChainType = "Json"           //
	MsgChainTypeApp            MessageChainType = "App"            //
	MsgChainTypePoke           MessageChainType = "Poke"           //
	MsgChainTypeDice           MessageChainType = "Dice"           //
	MsgChainTypeMarketFace     MessageChainType = "MarketFace"     //
	MsgChainTypeMusicShare     MessageChainType = "MusicShare"     //
	MsgChainTypeForwardMessage MessageChainType = "ForwardMessage" //
	MsgChainTypeFile           MessageChainType = "File"           //
	MsgChainTypeMiraiCode      MessageChainType = "MiraiCode"      //
)

var MessageChainTypeValue = map[MessageChainType]any{
	"source":         SourceMsgChain{},
	"quote":          QuoteMsgChain{},
	"at":             AtMsgChain{},
	"atAll":          AtAllMsgChain{},
	"face":           FaceMsgChain{},
	"plain":          PlainMsgChain{},
	"image":          ImageMsgChain{},
	"flashImage":     FlashImageMsgChain{},
	"voice":          VoiceMsgChain{},
	"xml":            XmlMsgChain{},
	"json":           JsonMsgChain{},
	"app":            AppMsgChain{},
	"poke":           PokeMsgChain{},
	"dice":           DiceMsgChain{},
	"marketFace":     MarketFaceMsgChain{},
	"musicShare":     MusicShareMsgChain{},
	"forwardMessage": ForwardMessageMsgChain{},
	"file":           FileMsgChain{},
	"miraiCode":      MiraiCodeMsgChain{},
}

type MessageChain struct {
	Type MessageChainType `json:"type"`
	//SourceMsgChain
	//QuoteMsgChain
	//AtMsgChain
	//AtAllMsgChain
	//FaceMsgChain
	//PlainMsgChain
	//ImageMsgChain
	//FlashImageMsgChain
	//VoiceMsgChain
	//XmlMsgChain
	//JsonMsgChain
	//AppMsgChain
	//PokeMsgChain
	//DiceMsgChain
	//MarketFaceMsgChain
	//MusicShareMsgChain
	//ForwardMessageMsgChain
	//FileMsgChain
	//MiraiCodeMsgChain

	// 消息的识别号，根据type不同分以下几种情况
	//  - Source: 消息的识别号，用于引用回复（Source类型永远为chain的第一个元素）
	//  - Quote: 被引用回复的原消息的messageId
	//  - MarketFace: 商城表情唯一标识
	//  - File: 文件识别id
	Id int `json:"id"`

	// Source
	Time int `json:"time"` // 时间戳

	// Quote 引用他人消息
	GroupId  int            `json:"groupId"`  // 被引用回复的原消息所接收的群号，当为好友消息时为0
	SenderId int            `json:"senderId"` // 被引用回复的原消息的发送者的QQ号
	TargetId int            `json:"targetId"` // 被引用回复的原消息的接收者者的QQ号（或群号）
	Origin   []MessageChain `json:"origin"`   // 被引用回复的原消息的消息链对象 Todo: 有待测试各种不同类型消息引用的情况

	// At
	Target  int    `json:"target"`  // 群员QQ号
	Display string `json:"display"` // At时显示的文字，发送消息时无效，自动使用群名片

	// Face QQ表情
	FaceId int `json:"faceId"` // QQ表情编号，可选，优先高于name

	// Poke
	//  - "Poke": 戳一戳
	//  - "ShowLove": 比心
	//  - "Like": 点赞
	//  - "Heartbroken": 心碎
	//  - "SixSixSix": 666
	//  - "FangDaZhao": 放大招
	// MarketFace 商城表情
	//  - 表情显示名称
	// Face QQ表情
	//  - 表情显示拼音
	// File
	//  - 文件名
	Name string `json:"name"`

	// Plain 文字消息
	Text string `json:"text"`

	// Image | FlashImage  照片 | 闪照
	// type为FlashImage（闪照）时，三个参数任选其一，出现多个参数时，按照imageId > url > path > base64的优先级
	ImageId string `json:"imageId"` // 图片的imageId，群图片与好友图片格式不同。不为空时将忽略url属性

	// Voice
	VoiceId string `json:"voiceId"` // 语音的voiceId，不为空时将忽略url属性
	Length  int    `json:"length"`  // 返回的语音长度, 发送消息时可以不传

	// Image | FlashImage | Voice
	//  - Image | FlashImage: 图片的URL，发送时可作网络图片的链接；接收时为腾讯图片服务器的链接，可用于图片下载
	//  - Voice: 语音的URL，发送时可作网络语音的链接；接收时为腾讯语音服务器的链接，可用于语音下载
	Url string `json:"url"`
	//  - Image | FlashImage: 图片的路径，发送本地图片，路径相对于 JVM 工作路径（默认是当前路径，可通过 -Duser.dir=...指定），也可传入绝对路径。
	//  - Voice: 语音的路径，发送本地语音，路径相对于 JVM 工作路径（默认是当前路径，可通过 -Duser.dir=...指定），也可传入绝对路径。
	Path string `json:"path"`
	//  - Image | FlashImage: 图片的 Base64 编码
	//  - Voice: 语音的 Base64 编码
	Base64 string `json:"base64"`

	// XML
	Xml string `json:"xml"`

	// JSON
	Json string `json:"json"`

	// App App分享
	Content string `json:"content"`

	// Dice 骰子
	Value int `json:"value"` // 点数

	// MusicShare 音乐分享
	Kind       string `json:"kind"`       // 类型
	Title      string `json:"title"`      // 标题
	Summary    string `json:"summary"`    // 概括
	JumpUrl    string `json:"jumpUrl"`    // 跳转路径
	PictureUrl string `json:"pictureUrl"` // 封面路径
	MusicUrl   string `json:"musicUrl"`   // 音源路径
	Brief      string `json:"brief"`      // 简介

	// ForwardMessage 转发消息
	NodeList []ForwardMessage `json:"nodeList"`

	// File
	Size int `json:"size"`

	// MiraiCode
	Code string `json:"code"`
}

func (m MessageChain) MarshalJSON() (b []byte, e error) {
	t, _err := m.ConvertType()
	if _err != nil {
		e = _err
	} else {
		b, e = json.Marshal(t)
	}
	return
}

func (m MessageChain) ConvertType() (t any, err error) {
	if m.Type == "" {
		err = fmt.Errorf("type 不能为空！")
		return
	}
	_type := MessageChainTypeValue[MessageChainType(strings.ToLower(string(m.Type)))]
	t, err = pkg.ConvertStruct(m, &_type)
	return
}

// ConvertMsgChain
// 将MessageChain转换为具体对应的数据类型对象
func ConvertMsgChain[T any](source MessageChain) (t T, err error) {
	if source.Type == "" {
		err = fmt.Errorf("type 不能为空！")
		return
	}
	_type := MessageChainTypeValue[MessageChainType(strings.ToLower(string(source.Type)))]
	_t, _err := pkg.ConvertStruct(source, &_type)
	if _err != nil {
		err = _err
		return
	}
	if typ, ok := _t.(T); ok {
		t = typ
	} else {
		err = fmt.Errorf("转换失败！")
	}
	return
}

type ForwardMessage struct {
	SenderId     int            `json:"senderId"`     // 发送人QQ号
	Time         int            `json:"time"`         // 发送时间
	SenderName   string         `json:"senderName"`   // 显示名称
	MessageChain []MessageChain `json:"messageChain"` // 消息数组
	MessageId    int            `json:"messageId"`    // 可以只使用消息messageId，从缓存中读取一条消息作为节点
}

type SourceMsgChain struct {
	Type MessageChainType `json:"type"`
	Id   int              `json:"id"`
	Time int              `json:"time"`
}

type QuoteMsgChain struct {
	Type     MessageChainType `json:"type"`
	Id       int              `json:"id"`
	GroupId  int              `json:"groupId"`
	SenderId int              `json:"senderId"`
	TargetId int              `json:"targetId"`
	Origin   []MessageChain   `json:"origin"`
}

type AtMsgChain struct {
	Type    MessageChainType `json:"type"`
	Target  int              `json:"target"`
	Display string           `json:"display"`
}

type AtAllMsgChain struct {
	Type MessageChainType `json:"type"`
}

type FaceMsgChain struct {
	Type   MessageChainType `json:"type"`
	FaceId int              `json:"faceId"`
	Name   string           `json:"name"`
}

type PlainMsgChain struct {
	Type MessageChainType `json:"type"`
	Text string           `json:"text"`
}

type ImageMsgChain struct {
	Type    MessageChainType `json:"type"`
	ImageId string           `json:"imageId"`
	Url     string           `json:"url"`
	Path    interface{}      `json:"path"`
	Base64  interface{}      `json:"base64"`
}

type FlashImageMsgChain struct {
	Type    MessageChainType `json:"type"`
	ImageId string           `json:"imageId"`
	Url     string           `json:"url"`
	Path    interface{}      `json:"path"`
	Base64  interface{}      `json:"base64"`
}

type VoiceMsgChain struct {
	Type    MessageChainType `json:"type"`
	VoiceId string           `json:"voiceId"`
	Url     string           `json:"url"`
	Path    interface{}      `json:"path"`
	Base64  interface{}      `json:"base64"`
	Length  int              `json:"length"`
}

type XmlMsgChain struct {
	Type MessageChainType `json:"type"`
	Xml  string           `json:"xml"`
}

type JsonMsgChain struct {
	Type MessageChainType `json:"type"`
	Json string           `json:"json"`
}

type AppMsgChain struct {
	Type    MessageChainType `json:"type"`
	Content string           `json:"content"`
}

type PokeMsgChain struct {
	Type MessageChainType `json:"type"`
	Name string           `json:"name"`
}

type DiceMsgChain struct {
	Type  MessageChainType `json:"type"`
	Value int              `json:"value"`
}

type MarketFaceMsgChain struct {
	Type MessageChainType `json:"type"`
	Id   int              `json:"id"`
	Name string           `json:"name"`
}

type MusicShareMsgChain struct {
	Type       MessageChainType `json:"type"`
	Kind       string           `json:"kind"`       // 类型
	Title      string           `json:"title"`      // 标题
	Summary    string           `json:"summary"`    // 概括
	JumpUrl    string           `json:"jumpUrl"`    // 跳转路径
	PictureUrl string           `json:"pictureUrl"` // 封面路径
	MusicUrl   string           `json:"musicUrl"`   // 音源路径
	Brief      string           `json:"brief"`      // 简介
}

type ForwardMessageMsgChain struct {
	Type     MessageChainType `json:"type"`
	NodeList []MessageChain   `json:"nodeList"`
}

type FileMsgChain struct {
	Type MessageChainType `json:"type"`
	Id   string           `json:"id"`
	Name string           `json:"name"`
	Size int              `json:"size"`
}

type MiraiCodeMsgChain struct {
	Type MessageChainType `json:"type"`
	Code string           `json:"code"`
}
