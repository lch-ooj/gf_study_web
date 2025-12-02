package main

import (
    _ "github.com/gogf/gf/contrib/drivers/mysql/v2"

    "github.com/gogf/gf/v2/os/gctx"

    "gf-demo-user/internal/cmd"
    _ "gf-demo-user/internal/logic"
)

func main() {
    cmd.Main.Run(gctx.New())
}