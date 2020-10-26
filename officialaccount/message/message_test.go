package message

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMsgDecode(t *testing.T) {
	m := `{"device_id":"AAYAAGTTmWLAEzWhAdkg-h-5t1m10S-b8TKWV5WS3_g@ilink.im.sdk","device_type":"gh_f867aa7db510","msg_id":870460595,"msg_type":"set","services":{"wxmsg_file":{"type":"","name":"","size":0,"md5":"","download_url":"http://wxapp.tc.qq.com/202/20303/stodownload?filekey=30350201010421301f020200ca0402534804107006f733e3cb9e733a6b60549d712cdc020304118c040d00000004627466730000000131&hy=SH&storeid=32303230313032363137303530323030306235313831323435333562363465653231396430393030303030306361&bizid=1023"},"wxmsg_image":{"type":"","name":"","size":0,"md5":"","download_url":"http://wxapp.tc.qq.com/202/20303/stodownload?filekey=30350201010421301f020200ca0402534804107006f733e3cb9e733a6b60549d712cdc020304118c040d00000004627466730000000131&hy=SH&storeid=32303230313032363137303530323030306235313831323435333562363465653231396430393030303030306361&bizid=1023","enckey":"547384705817645346"}},"user":"ovuIruOOq2YyJrWYwhHPhL8Zz5Nk"}`
	var msg MixMessage
	err := json.Unmarshal([]byte(m), &msg)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", msg)
}
