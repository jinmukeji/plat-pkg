package microconfig_test

import (
	"testing"
	"time"

	m "github.com/jinmukeji/plat-pkg/jwt/keystore/micro-config"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/encoder/yaml"
	"github.com/micro/go-micro/config/source"
	"github.com/micro/go-micro/config/source/consul"
	"github.com/stretchr/testify/suite"

	"github.com/hashicorp/consul/api"
)

const (
	cfgConsulPrefix = "/micro/config/jm"
	cfgConsulAddr   = "localhost:8500"
)

type MicroConfigTestSuite struct {
	suite.Suite
}

// SetupSuite 设置测试环境
func (suite *MicroConfigTestSuite) SetupSuite() {
	// 连接 Consul 并读取配置信息

	encoder := yaml.NewEncoder()

	consulSource := consul.NewSource(
		// optionally specify consul address;
		consul.WithAddress(cfgConsulAddr),
		// optionally specify prefix;
		consul.WithPrefix(cfgConsulPrefix),
		// optionally strip the provided prefix from the keys
		consul.StripPrefix(true),
		source.WithEncoder(encoder),
	)

	if err := config.Load(consulSource); err != nil {
		suite.FailNow(err.Error())
	}
}

func (suite *MicroConfigTestSuite) TestMicroConfigStore_Get() {
	baseKeyPath := []string{"platform", "app-key"}
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
	baseKeyPath := []string{"platform", "app-key"}
	store := m.NewMicroConfigStore(baseKeyPath...)
	key := store.Get("app-test-not-exists")
	suite.Assert().Nil(key)
}

func (suite *MicroConfigTestSuite) TestMicroConfigStore_Get_Disabled() {
	baseKeyPath := []string{"platform", "app-key"}
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

	// Get a new client
	consulCfg := api.DefaultConfig()
	consulCfg.Address = cfgConsulAddr
	client, err := api.NewClient(consulCfg)
	if err != nil {
		suite.FailNow(err.Error())
	}
	kv := client.KV()

	consulKey := "micro/config/jm/platform/app-key/app-test2"
	p := &api.KVPair{Key: consulKey, Value: []byte(pValue)}
	_, err = kv.Put(p, nil)
	if err != nil {
		suite.FailNow(err.Error())
	}
	time.Sleep(500 * time.Millisecond)

	baseKeyPath := []string{"platform", "app-key"}
	store := m.NewMicroConfigStore(baseKeyPath...)
	key := store.Get("app-test2")
	suite.Assert().NotNil(key)
	suite.Assert().Equal("app-test2", key.ID())

	// 变更配置内容
	newPValue := pValue + "testing: true"
	newP := &api.KVPair{Key: consulKey, Value: []byte(newPValue)}
	_, err = kv.Put(newP, nil)
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