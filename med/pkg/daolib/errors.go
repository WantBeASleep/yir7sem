package daolib

import "errors"

var NoTransactionError = errors.New("ctx doesn't have tx")
