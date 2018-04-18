package main

import (
	"net/http"
	"sync"
)

var (
	varsLock sync.RWMutex
	vars     map[*http.Request]map[string]interface{} // vars は複数のHTTPリクエストが同時にアクセスして変更を試みる
)

func OpenVars(r *http.Request) {
	varsLock.Lock()

	// varsが空のときは、指定されたhttp.Requestへのポインタをキーとして空のマップをvarsに追加し、
	// 最後にロックを開放して他のハンドラからのアクセスを可能にします。
	if vars == nil {
		vars = map[*http.Request]map[string]interface{}{}
	}
	vars[r] = map[string]interface{}{}
	varsLock.Unlock()
}

// CloseVars は指定したリクエストに対応するエントリがマップvarsから安全に削除される
func CloseVars(r *http.Request) {
	varsLock.Lock()
	delete(vars, r)
	varsLock.Unlock()
}

func GetVar(r *http.Request, key string) interface{} {
	// 誰かがLockを行っていると、ロックが解除されるまでは他のコードはLockを
	// 試みると実行がブロックされてしまうが、RLockでは他のコードによるRLockを
	// をブロックすることはない。
	varsLock.RLock()
	value := vars[r][key]
	varsLock.RUnlock()
	return value
}
func SetVar(r *http.Request, key string, value interface{}) {
	varsLock.Lock()
	vars[r][key] = value
	varsLock.Unlock()
}
