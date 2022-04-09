package modules

import (
	"github.com/quantumsheep/plouf"
	modules_users "github.com/quantumsheep/plouf/example/modules/users"
)

type MainModule struct {
	plouf.Module

	UsersModule *modules_users.UsersModule
}
