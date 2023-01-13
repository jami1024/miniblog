// Copyright 2023 jami1024 &lt;996013797@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/jami1024/miniblog.

package main

import (
	"os"

	"github.com/jami1024/miniblog/internal/miniblog"
	// _ "go.uber.org/automaxprocs"
)

// Go 程序的默认入口函数(主函数).
func main() {
	command := miniblog.NewMiniBlogCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
