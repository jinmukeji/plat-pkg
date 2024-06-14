package microconfig_test

import (
	"context"
	"flag"
	"os"
	"testing"
	"time"

	m "github.com/jinmukeji/plat-pkg/v4/auth/jwt/keystore/micro-config"

	"github.com/go-micro/plugins/v4/config/encoder/yaml"
	"github.com/go-micro/plugins/v4/config/source/etcd"
	"github.com/stretchr/testify/suite"
	"go-micro.dev/v4/config"
	"go-micro.dev/v4/config/source"
	clientv3 "go.etcd.io/etcd/client/v3"
)

const (
	cfgEtcdPrefix = "/micro/config/jm/"
	cfgEtcdAddr   = "localhost:2379"
)

var (
	etcdAddress = flag.String("etcd.server", cfgEtcdAddr, "etcd address")
)

func init() {
	if len(os.Getenv("ETCD_ADDR")) > 0 {
		*etcdAddress = os.Getenv("ETCD_ADDR")
	}
}

type MicroConfigTestSuite struct {
	suite.Suite
}

// SetupSuite 设置测试环境
func (suite *MicroConfigTestSuite) SetupSuite() {
	// 连接 Etcd 并读取配置信息

	encoder := yaml.NewEncoder()

	etcdSource := etcd.NewSource(
		// optionally specify etcd address;
		etcd.WithAddress(*etcdAddress),
		// optionally specify prefix;
		etcd.WithPrefix(cfgEtcdPrefix),
		// optionally strip the provided prefix from the keys
		// etcd.StripPrefix(true),
		source.WithEncoder(encoder),
	)

	if err := config.Load(etcdSource); err != nil {
		suite.FailNow(err.Error())
	}
}

func (suite *MicroConfigTestSuite) TestMicroConfigStore_Get() {
	baseKeyPath := []string{"micro", "config", "jm", "platform", "app-key"}
	store := m.NewMicroConfigStore(baseKeyPath...)
	key := store.Get("app-test1")
	suite.Assert().NotNil(key)
	suite.Assert().Equal("app-test1", key.ID())
	suite.Assert().Equal("5c:32:dd:4c:b7:21:1a:7f:c2:31:ca:d2:f5:51:77:bc:78:dc:65:ac", key.Fingerprint())
	suite.Assert().NotNil(key.PublicKey())
	bigN := "24814136704917335065189361067306516922621475879872786673955385327384947446008875021460827146828269021172599285405419359895519758108949983553592776845039010176619175729934726402266149332866163773671594470478708307141054279031430051537214125282548066162870539103108156537067221184418732571173692408966204260362965560985263458004213164004538677190199882194355079546598050580521940976499375532156605970889530663798302058170003182927293077456851560546617407577135682107175710033340202967014299509043524427941607441448580871654610158802050632106062514136382116419078777272923001832767886340149931903103741512844582902745827"
	suite.Assert().Equal(bigN, key.PublicKey().N.String())

	// 第二次读取来自 cache
	key = store.Get("app-test1")
	suite.Assert().NotNil(key)
}

func (suite *MicroConfigTestSuite) TestMicroConfigStore_Get_NotExists() {
	baseKeyPath := []string{"micro", "config", "jm", "platform", "app-key"}
	store := m.NewMicroConfigStore(baseKeyPath...)
	key := store.Get("app-test-not-exists")
	suite.Assert().Nil(key)
}

func (suite *MicroConfigTestSuite) TestMicroConfigStore_Get_Disabled() {
	baseKeyPath := []string{"micro", "config", "jm", "platform", "app-key"}
	store := m.NewMicroConfigStore(baseKeyPath...)
	key := store.Get("app-test4")
	suite.Assert().Nil(key)
}

func (suite *MicroConfigTestSuite) TestMicroConfigStore_Get_ConfigChanged() {
	const pValue = `id: "app-test2"
disabled: false
fingerprint: "15:e7:c6:d3:b5:fe:30:12:d4:cb:65:e5:73:09:f5:e2:ac:68:e2:c2"
public_key: |
    -----BEGIN PUBLIC KEY-----
    MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA0gALJf5Nh/GexqVTm9sX
    /Xr6kki6FAtbawvr8ZV2E6xP0DJN5RIPdXkAGlGI0Ob5iKUh7tbYw6c/6QzR1PqO
    MkKHLGiRh8VweclP/LUSWQ8uTNbBvJ8KvmEt0KkeGSVNiaOcdKOPtoVxgYzMa53t
    8w+J/5BE1ufDppUCuCMdqKjkqLCv6HelkT3E2dE5JVmKiayGMYQYAaotiP/aWpLR
    drhVQ3QRckTM7rVVBTZoCual8ggRE5UtCHP6qOq1TBg/oBa6vg/u6EpkgqsQSjPh
    +YrOKvx5N/qBike+q62musbepwjmQAfrwBACsUTo05nwX5m/b3wACG++A3ASPov6
    bQIDAQAB
    -----END PUBLIC KEY-----
`

	// FIXME: ETCD

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{*etcdAddress},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		suite.FailNow(err.Error())
	}
	defer cli.Close()

	etcdKey := "/micro/config/jm/platform/app-key/app-test2"
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	_, err = cli.Put(ctx, etcdKey, pValue)
	cancel()
	if err != nil {
		suite.FailNow(err.Error())
	}

	time.Sleep(500 * time.Millisecond)

	baseKeyPath := []string{"micro", "config", "jm", "platform", "app-key"}
	store := m.NewMicroConfigStore(baseKeyPath...)
	key := store.Get("app-test2")
	suite.Assert().NotNil(key)
	suite.Assert().Equal("app-test2", key.ID())

	// 变更配置内容
	newPValue := pValue + "testing: true"
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	_, err = cli.Put(ctx, etcdKey, newPValue)
	cancel()
	if err != nil {
		suite.FailNow(err.Error())
	}

	time.Sleep(500 * time.Millisecond)

	key = store.Get("app-test2")
	suite.Assert().NotNil(key)
	suite.Assert().Equal("app-test2", key.ID())
}

func TestMicroConfigTestSuite(t *testing.T) {
	suite.Run(t, new(MicroConfigTestSuite))
}
