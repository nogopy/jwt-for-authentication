package exception

import "errors"

var UsernameExists = errors.New("username exists")
var WrongPassword = errors.New("wrong password")
