package main

import "fmt"

const (
	USER = "test_user"
	WRONG_USER = "wrong_user"
	PASSWORD = "password"
	GET = "GET"
	URL = "/users"
)

func main() {
	body := map[string]string{
		"password": PASSWORD,
	}

	proxy := NewProxy()

	// 不正ユーザー
	body["user"] = WRONG_USER
	code, res := proxy.HandleRequest(URL, GET, body)
	fmt.Printf("Status Code: %d, Response: %s\n\n", code, res)

	// 正規ユーザー
	body["user"] = USER
	code, res = proxy.HandleRequest(URL, GET, body)
	fmt.Printf("Status Code: %d, Response: %s\n\n", code, res)

	// キャッシュ
	code, res = proxy.HandleRequest(URL, GET, body)
	fmt.Printf("Status Code: %d, Response: %s\n\n", code, res)

	// 不正URL
	code, res = proxy.HandleRequest("", GET, body)
	fmt.Printf("Status Code: %d, Response: %s\n\n", code, res)
}