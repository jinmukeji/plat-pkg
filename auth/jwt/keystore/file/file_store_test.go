package file_test

import (
	"testing"

	"github.com/jinmukeji/plat-pkg/v4/auth/jwt/keystore/file"
	"github.com/stretchr/testify/suite"
)

// FileStoreTestSuite 是 Echo rpc 的单元测试的 Test Suite
type FileStoreTestSuite struct {
	suite.Suite
}

const (
	testdataDir = "../../tools/testdata"
)

// SetupSuite 设置测试环境
func (suite *FileStoreTestSuite) SetupSuite() {
}

// TestLoad 测试 Load 方法
func (suite *FileStoreTestSuite) TestLoad() {
	store := file.NewFileStore()

	err := store.Load(testdataDir, "app-test1")
	suite.Assert().NoError(err)

	err = store.Load(testdataDir, "app-test2")
	suite.Assert().NoError(err)

	err = store.Load(testdataDir, "app-test-not-exist")
	suite.Assert().Error(err)
}

func (suite *FileStoreTestSuite) TestFileStore_Get() {
	store := file.NewFileStore()
	_ = store.Load(testdataDir, "app-test1")

	key := store.Get("app-test1")
	suite.Assert().NotNil(key)
	suite.Assert().Equal("app-test1", key.ID())
	suite.Assert().Equal("5c:32:dd:4c:b7:21:1a:7f:c2:31:ca:d2:f5:51:77:bc:78:dc:65:ac", key.Fingerprint())
	suite.Assert().NotNil(key.PublicKey())
	bigN := "24814136704917335065189361067306516922621475879872786673955385327384947446008875021460827146828269021172599285405419359895519758108949983553592776845039010176619175729934726402266149332866163773671594470478708307141054279031430051537214125282548066162870539103108156537067221184418732571173692408966204260362965560985263458004213164004538677190199882194355079546598050580521940976499375532156605970889530663798302058170003182927293077456851560546617407577135682107175710033340202967014299509043524427941607441448580871654610158802050632106062514136382116419078777272923001832767886340149931903103741512844582902745827"
	suite.Assert().Equal(bigN, key.PublicKey().N.String())

	key = store.Get("app-test-not-loaded")
	suite.Assert().Nil(key)
}

func TestFileStoreTestSuite(t *testing.T) {
	suite.Run(t, new(FileStoreTestSuite))
}
