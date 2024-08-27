package whitelist

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type WhitelistPolicy struct {
	ID      string   `json:"id" yaml:"id"`
	Presets []string `json:"presets" yaml:"presets"`
}

type PolicySet struct {
	set     map[string]*WhitelistPolicy
	presets Presets
}

func NewPolicySet(policies []*WhitelistPolicy, presets Presets) *PolicySet {
	set := make(map[string]*WhitelistPolicy)
	for _, p := range policies {
		set[p.ID] = p
	}

	return &PolicySet{
		set:     set,
		presets: presets,
	}
}

func (ps PolicySet) GetWhitelistPolicy(id string) *WhitelistPolicy {
	return ps.set[id]
}

func (ps PolicySet) GetAllPatterns() []string {
	patternSet := make(map[string]bool)

	for _, p := range ps.set {
		if p != nil {
			for _, presetName := range p.Presets {
				for _, pattern := range ps.presets.GetRulePatterns(presetName) {
					patternSet[pattern] = true
				}
			}
		}
	}

	keys := make([]string, 0, len(patternSet))
	for k := range patternSet {
		keys = append(keys, k)
	}

	return keys
}

func LoadPolicy(data []byte) (*WhitelistPolicy, error) {
	w := WhitelistPolicy{}
	err := yaml.Unmarshal(data, &w)
	if err != nil {
		return nil, err
	}
	return &w, nil
}

// LoadPolicyFromYamlFile 从单个 yaml 配置文件加载
func LoadPolicyFromYamlFile(file string) (*WhitelistPolicy, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return LoadPolicy(content)
}

// LoadPolicyFromYamlDir 从文件目录中加载所有 yaml 文件定义的配置，重复定义将返回错误
func LoadPolicyFromYamlDir(dir string) ([]*WhitelistPolicy, error) {
	r := make([]*WhitelistPolicy, 0)

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		path := filepath.Join(dir, file.Name())
		ext := filepath.Ext(path)
		if ext == ".yml" || ext == ".yaml" {
			p, err := LoadPolicyFromYamlFile(path)
			if err != nil {
				return nil, err
			}
			r = append(r, p)
		}
	}

	return r, nil
}
