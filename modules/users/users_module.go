package modules_users

import "github.com/quantumsheep/plouf"

type UsersModule struct {
	plouf.Module

	UsersController *UsersController
}
