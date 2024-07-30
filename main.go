package main

import (
	"goframe-shop-v2/internal/cmd"
	_ "goframe-shop-v2/internal/logic"
	_ "goframe-shop-v2/internal/mq"
	_ "goframe-shop-v2/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.New())
}
