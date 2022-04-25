package model

type FileParams struct {
	Id               string `json:"id"`                         // 文件夹id
	Path             string `json:"path,omitempty"`             // 文件夹路径, 文件夹允许重名, 不保证准确, 准确定位使用 id
	Target           int    `json:"target,omitempty"`           // 群号或好友QQ号
	Group            int    `json:"group,omitempty"`            // 群号
	Qq               int    `json:"qq,omitempty"`               // 好友QQ号
	WithDownloadInfo bool   `json:"withDownloadInfo,omitempty"` // 是否携带下载信息，额外请求，无必要不要携带
}

type File struct {
	Name         string        `json:"name"`
	Id           string        `json:"id"`
	Path         string        `json:"path"`
	Parent       *File         `json:"parent"`
	Contact      GroupOrFriend `json:"contact"`
	IsFile       bool          `json:"isFile"`
	IsDictionary bool          `json:"isDictionary"`
	IsDirectory  bool          `json:"isDirectory"`
	DownloadInfo DownloadInfo  `json:"downloadInfo"`
}

type DownloadInfo struct {
	Sha1           string `json:"sha1"`
	Md5            string `json:"md5"`
	DownloadTimes  int    `json:"downloadTimes"`
	UploaderId     int    `json:"uploaderId"`
	UploadTime     int    `json:"uploadTime"`
	LastModifyTime int    `json:"lastModifyTime"`
	Url            string `json:"url"`
}
