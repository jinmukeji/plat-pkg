module github.com/jinmukeji/plat-pkg

go 1.13

// TODO: fix go mod tidy
replace github.com/gogo/protobuf => github.com/gogo/protobuf v1.3.0

require (
	github.com/chzyer/logex v1.1.10 // indirect
	github.com/chzyer/test v0.0.0-20180213035817-a1ea475d72b1 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-telegram-bot-api/telegram-bot-api v4.6.4+incompatible // indirect
	github.com/gobwas/glob v0.2.3
	github.com/hashicorp/consul/api v1.2.0
	github.com/jinmukeji/go-pkg v0.0.0-20191004044456-0aded5c0032f
	github.com/jinmukeji/proto v0.0.0-20191006061359-d74fb967d82b
	github.com/jinzhu/gorm v1.9.11
	github.com/lusis/go-slackbot v0.0.0-20180109053408-401027ccfef5 // indirect
	github.com/lusis/slack-test v0.0.0-20190426140909-c40012f20018 // indirect
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.11.0
	github.com/micro/go-plugins v1.3.0
	github.com/micro/micro v1.11.0
	github.com/sirupsen/logrus v1.4.2
	github.com/stretchr/testify v1.4.0
	gopkg.in/yaml.v2 v2.2.4
)
