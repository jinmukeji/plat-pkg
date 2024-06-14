package whitelist

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Presets map[string][]string

func (p Presets) GetRulePatterns(name string) []string {
	return p[name]
}

func LoadPreset(data []byte) (Presets, error) {
	p := Presets{}
	err := yaml.Unmarshal(data, &p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// LoadPresetFromYamlFile 从单个 yaml 配置文件加载
func LoadPresetFromYamlFile(file string) (Presets, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return LoadPreset(content)
}

// LoadPresetFromYamlDir 从文件目录中加载所有 yaml 文件定义的配置，重复定义将返回错误
func LoadPresetFromYamlDir(dir string) (Presets, error) {
	r := make(Presets)

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
			p, err := LoadPresetFromYamlFile(path)
			if err != nil {
				return nil, err
			}
			r, err = mergePreset(r, p)
			if err != nil {
				return nil, err
			}
		}
	}

	return r, nil
}

func mergePreset(a, b Presets) (Presets, error) {
	if a == nil || b == nil {
		return nil, errors.New("empty preset can't be merged")
	}

	r := make(Presets)

	for k, v := range a {
		r[k] = v
	}

	for k, v := range b {
		if _, exist := r[k]; exist {
			return nil, fmt.Errorf("duplicate preset: %q", k)
		}

		r[k] = v
	}

	return r, nil
}
