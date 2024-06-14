package version

import (
	"fmt"

	mlog "github.com/jinmukeji/go-pkg/v2/log"
)

var (
	log = mlog.StandardLogger()
)

// VersionInfo 版本信息
type VersionInfo interface {
	GetProductName() string
	GetProductVersion() string
	GetGitCommit() string
	GetBuildTime() string
	GetGoVersion() string
}

// PrintFullVersionInfo 打印版本信息
func PrintFullVersionInfo(v VersionInfo) {
	fmt.Println("Product:", v.GetProductName(), v.GetProductVersion())
	fmt.Println("GitCommit:", v.GetGitCommit())
	fmt.Println("BuildTime:", v.GetBuildTime())
	fmt.Println("Go:", v.GetGoVersion())
}

// LogVersionInfo 日志输出版本信息
func LogVersionInfo(v VersionInfo) {
	log.Infof("Product: %s %s", v.GetProductName(), v.GetProductVersion())
	log.Infof("Git Commit: %s", v.GetGitCommit())
	log.Infof("Build Time: %s", v.GetBuildTime())
	log.Infof("Go Version: %s", v.GetGoVersion())
}
