package microconfig

import (
	"crypto/rsa"
	"strings"
	"sync"

	"github.com/jinmukeji/plat-pkg/v4/auth/jwt"
	"github.com/jinmukeji/plat-pkg/v4/auth/jwt/keystore"
	"go-micro.dev/v4/config"
)

type keyItem struct {
	keyPath   []string
	publicKey *rsa.PublicKey

	CfgID          string `json:"id" yaml:"id"`
	CfgFingerprint string `json:"fingerprint" yaml:"fingerprint"`
	CfgPublicKey   string `json:"public_key" yaml:"public_key"`
	Disabled       bool   `json:"disabled" yaml:"disabled"`
}

var _ keystore.KeyItem = (*keyItem)(nil)

func (i keyItem) ID() string {
	return i.CfgID
}

func (i keyItem) PublicKey() *rsa.PublicKey {
	return i.publicKey
}

func (i keyItem) Fingerprint() string {
	return i.CfgFingerprint
}

type MicroConfigStore struct {
	baseCfg    []string
	cachedKeys sync.Map
}

var _ keystore.Store = (*MicroConfigStore)(nil)

func NewMicroConfigStore(baseCfg ...string) *MicroConfigStore {
	s := &MicroConfigStore{
		baseCfg:    baseCfg,
		cachedKeys: sync.Map{},
	}

	// watch changes on an isolated goroutine
	// 监控配置变化，当变化发生时重新创建 cachedKeys
	go watchConfigChanges(s)

	return s
}

func watchConfigChanges(s *MicroConfigStore) {
	watcher, err := config.Watch(s.baseCfg...)
	if err != nil {
		// 不能正常开启监听器，则 Fatal 终止程序
		log.Fatalf("start watching config on error: %s", err)
		return
	}

	cfgPath := configPath(s.baseCfg...)
	log.Infof("start watching config: %s", cfgPath)

	for {
		// 监听配置变更过程中，如果产生错误，不退出程序，但需要输出 ERROR 级别日志

		v, err := watcher.Next()
		if err != nil {
			log.Errorf("watch config error，%s", err)
			continue
		}

		log.Infof("Config is changed: %s", cfgPath)
		log.Debugf("New config: %s=%s", cfgPath, string(v.Bytes()))

		// 清除缓存，全部重新加载
		s.eraseCache()
	}
}

func configPath(cfgKey ...string) string {
	return strings.Join(cfgKey, "/")
}

func (s *MicroConfigStore) Get(id string) keystore.KeyItem {
	// 首先从缓存中找
	if k, ok := s.loadCache(id); ok && !k.Disabled {
		return k
	}

	item, err := s.loadFromConfig(id)
	if err != nil {
		return nil
	}

	// 添加到缓存
	s.storeCache(id, item)

	if item.Disabled {
		return nil
	}

	return item
}

func (s *MicroConfigStore) storeCache(id string, key *keyItem) {
	s.cachedKeys.Store(id, key)
}

func (s *MicroConfigStore) loadCache(id string) (*keyItem, bool) {
	result, ok := s.cachedKeys.Load(id)
	if ok {
		return result.(*keyItem), true
	} else {
		return nil, false
	}
}

func (s *MicroConfigStore) eraseCache() {
	s.cachedKeys.Range(func(key interface{}, value interface{}) bool {
		s.cachedKeys.Delete(key)
		return true
	})
}

func (s *MicroConfigStore) loadFromConfig(id string) (*keyItem, error) {
	item := keyItem{}

	// 从配置中找
	path := s.keyPath(id)
	if err := config.Get(path...).Scan(&item); err != nil {
		return nil, err
	}

	// found
	item.keyPath = path
	publicKey, err := jwt.LoadRSAPublicKey([]byte(item.CfgPublicKey))
	if err != nil {
		return nil, err
	}
	item.publicKey = publicKey

	return &item, nil
}

func (s *MicroConfigStore) keyPath(id string) []string {
	// <base_key_path>/:id
	keyPath := append(s.baseCfg, id)
	return keyPath
}
