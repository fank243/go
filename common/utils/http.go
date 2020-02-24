package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

/**
发送GET请求
@headers 请求头参数
@content 响应参数
@statusCode 响应状态码
*/
func Get(url string, headers map[string]string) (content string, statusCode int) {
	return sendHttp("GET", url, headers, "")
}

/**
发送POST请求
@headers 请求头参数
@body 请求体参数
@content 响应参数
@statusCode 响应状态码
*/
func Post(url string, headers map[string]string, body string) (content string, statusCode int) {
	return sendHttp("POST", url, headers, body)
}

/**
发送　HTTP请求
@method 请求方式
@url 请求地址
@headers 请求头参数
@req 请求体
@content 响应参数
@statusCode 响应状态码
*/
func sendHttp(method string, url string, headers map[string]string, body string) (content string, statusCode int) {
	client := http.Client{}

	var req *http.Request
	if method == "POST" {
		req, _ = http.NewRequest("POST", url, strings.NewReader(body))
	} else {
		req, _ = http.NewRequest("GET", url, nil)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	if headers != nil && len(headers) > 0 {
		for header := range headers {
			req.Header.Add(header, headers[header])
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		statusCode = 500
		fmt.Println(err.Error())
		return
	}

	defer resp.Body.Close()

	result, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		statusCode = resp.StatusCode
		fmt.Println(err1.Error())
		return
	}

	content = string(result)
	statusCode = resp.StatusCode
	return
}
