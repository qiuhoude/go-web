package json

type (
	Server struct {
		// ID 不会导出到JSON中
		ID int `json:"-"`
		// ServerName 的值会进行二次JSON编码
		ServerName  string `json:"serverName"`
		ServerName2 string `json:"serverName2,string,omitempty"`
		// 如果 ServerIP 为空，则不输出到JSON串中
		ServerIP string `json:"serverIP,omitempty"`
	}
	Serverslice struct {
		Servers []Server
	}
)
