package main

import (
	"context"
	"fmt"
	"runtime"

	"github.com/ebitengine/purego"
)

/*
*
* 插件接口
*
 */
type DynamicPlugin interface {
	Init(map[string]interface{}) error
	Start(context.Context, context.CancelFunc)
	Stop() error
}

/*
*
* 插件结构体
*
 */
type CC1Plugin struct {
	Type   string
	ctx    context.Context
	cancel context.CancelFunc
}

func getSystemLibrary() string {
	switch runtime.GOOS {
	case "linux":
		return "./libtest.so"
	default:
		panic(fmt.Errorf("GOOS=%s is not supported", runtime.GOOS))
	}
}

func main() {
	libc, err := purego.Dlopen(getSystemLibrary(),
		purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	var fun1 func()
	purego.RegisterLibFunc(&fun1, libc, "fun1")
	go fun1()
	println("test ok")
}
