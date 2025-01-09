package static

import "errors"

var (
	ErrUninitializedDatabase = errors.New("error cannot initialize database")
	ErrNoAppConfig           = errors.New("error no app config")
)
