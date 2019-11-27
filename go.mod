module github.com/jinmukeji/plat-pkg

go 1.13

// TODO: fix go mod tidy
replace (
	github.com/gogo/protobuf => github.com/gogo/protobuf v1.3.1
	github.com/nicksnyder/go-i18n => github.com/nicksnyder/go-i18n v1.10.1
)

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gobwas/glob v0.2.3
	github.com/jinmukeji/go-pkg v0.0.0-20191105023801-49bf3fa2962f
	github.com/jinmukeji/proto v0.0.0-20191106063829-0f6255fab313
	github.com/jinzhu/gorm v1.9.11
	github.com/manifoldco/promptui v0.3.2 // indirect
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.17.1
	github.com/micro/go-plugins v1.5.1
	github.com/micro/micro v1.16.0
	github.com/nicksnyder/go-i18n v0.0.0-00010101000000-000000000000 // indirect
	github.com/samfoo/ansi v0.0.0-20160124022901-b6bd2ded7189 // indirect
	github.com/sirupsen/logrus v1.4.2
	github.com/smallstep/assert v0.0.0-20180720014142-de77670473b5 // indirect
	github.com/smallstep/cli v0.13.3
	github.com/stretchr/testify v1.4.0
	go.etcd.io/etcd v3.3.17+incompatible
	gopkg.in/alecthomas/kingpin.v3-unstable v3.0.0-20191105091915-95d230a53780 // indirect
	gopkg.in/yaml.v2 v2.2.7
)
