package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex

type Proxy struct {
	app *App
	cache map[string](map[int]string) // ex: {url_method: {statusCode: response}}
}

func NewProxy() *Proxy {
	return &Proxy{
		cache: map[string](map[int]string){},
	}
}

func (p *Proxy) HandleRequest(url, method string, body map[string]string) (int, string) {
	// 認証 (Access Proxy: アクセスプロキシ)
	user := body["user"]
	password := body["password"]
	if ok := p.checkAuth(user, password); !ok {
		return 401, "Unauthorized"
	}

	// 遅延生成 (Virtual Proxy: 仮想プロキシ)
	p.createApp()

	// キャッシュ (Cache Proxy: キャッシュプロキシ)
	key := fmt.Sprintf("%s_%s", url, method)
	if res, ok := p.cache[key]; ok {
		for code, response := range res {
			fmt.Println("response from proxy cache.")
			return code, response
		}
	}

	statusCode, response := p.app.HandleRequest(url, method, body)
	p.cache[key] = map[int]string{statusCode: response}
	return statusCode, response
}

func (p *Proxy) checkAuth(user, password string) bool {
	if user == "test_user" && password == "password" {
		return true
	}
	return false
}

func (p *Proxy) createApp() {
	mu.Lock()
	if p.app == nil {
		fmt.Println("create App instance.")
		p.app = NewApp()
	}
	mu.Unlock()
}