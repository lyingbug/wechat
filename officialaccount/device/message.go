package device

import (
	"encoding/json"
	"fmt"
	"github.com/lyingbug/wechat/v2/util"
)

type WxMsgMusic struct {
	Title       string `json:"title"`
	Artist      string `json:"artist"`
	Url         string `json:"url"`
	DataUrl     string `json:"data_url"`
	LowUrl      string `json:"low_url"`
	LowDataUrl  string `json:"low_data_url"`
	FromAppName string `json:"from_appname"`
}
type WxMsgImage struct {
	Type        string `json:"type"`
	Name        string `json:"name"`
	Size        int    `json:"size"`
	Md5         string `json:"md5"`
	DownloadUrl string `json:"download_url"`
	EncKey      string `json:"enckey"`
}

type Services struct {
	WxMsgMusic *WxMsgMusic `json:"wxmsg_music"`
	WxMsgImage *WxMsgImage `json:"wxmsg_image"`
}

//MsgDevice 设备消息响应
type MsgDevice struct {
	DeviceType string   `json:"device_type"`
	DeviceID   string   `json:"device_id"`
	SessionID  string   `json:"session_id"`
	OpenID     string   `json:"user"`
	Services   Services `json:"services"`
}

// ReqTransMsg 主动发消息给设备
type ReqTransMsg struct {
	DeviceID   string `json:"device_id"`
	DeviceType string `json:"device_type"`
	OpenID     string `json:"open_id"`
	Content    string `json:"content"`
}

// TransMsg 给设备发消息
func (d *Device) TransMsg(req ReqTransMsg) (err error) {
	var accessToken string
	if accessToken, err = d.GetAccessToken(); err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", uriTransMsg, accessToken)
	var response []byte
	if response, err = util.PostJSON(uri, req); err != nil {
		return
	}
	var result resBind
	if err = json.Unmarshal(response, &result); err != nil {
		return
	}
	if result.BaseResp.ErrCode != 0 {
		err = fmt.Errorf("DeviceBind Error , errcode=%d , errmsg=%s", result.BaseResp.ErrCode, result.BaseResp.ErrMsg)
		return
	}
	return
}
