package main

import (
	"fmt"
	"github.com/fank243/go/common/utils"
	"net/url"
	"strconv"
	"time"
)

//测试POST
func testGet() {
	content, statusCode := utils.Get("https://www.baidu.com", nil)
	fmt.Println(content, statusCode)
}

//测试POST
func testPost() {
	var headers map[string]string
	headers = make(map[string]string)
	headers["AccessToken"] = "d98d81c9980727a11fbf3cc1b7726e5d"

	body := url.Values{}
	var timestamp = time.Now().UnixNano() / 1e6

	body.Set("apiVer", "10")
	body.Set("version", "1.1.1")
	body.Set("timestamp", strconv.FormatInt(timestamp, 10))
	body.Set("deviceType", "1")
	body.Set("deviceNumber", "asdfsafsaf")
	body.Set("payload", "")

	var plainText = body.Get("apiVer") + body.Get("timestamp") + body.Get("deviceNumber")
	body.Set("signature", utils.Md5Encode(plainText))

	content, statusCode := utils.Post("http://api.test.class2dance.com/app/common/upgrade", headers, body.Encode())

	fmt.Println(content, statusCode)
}

func main() {
	testGet()

}
