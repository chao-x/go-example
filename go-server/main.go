package main

import (
	_ "go_server/common/conf"
	_ "go_server/common/mysql"
)

func main() {
	InitHandler()
}
