package service

import mlog "github.com/jinmukeji/go-pkg/v2/log"

var (
	log = mlog.StandardLogger()
)

func Logger() *mlog.Logger {
	return mlog.StandardLogger()
}
