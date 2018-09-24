package helps

import (
	"fmt"
	"testing"
)

func TestMail(t *testing.T) {
	userMail := "ymagain@qq.com"
	body := fmt.Sprintf(`
	<p>If your need change WorkOS.top password, please click this button Within 48 hours. </p>
	<a href="http://workos.top/change-password/%s/%s/" >Change My Password</a>
`, userMail, "token123")
	SendMail("xxx@163.com", "passwordxxx", "smtp.163.com:25", userMail, "WorkOS Change Password", body, "html")
}
