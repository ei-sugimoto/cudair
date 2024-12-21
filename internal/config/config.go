package config

import "github.com/BurntSushi/toml"

type CudairConfig struct {
	Root   string            `toml:"root"`
	TmpDir string            `toml:"tmp_dir"`
	Build  CudairBuildConfig `toml:"build"`
}

type CudairBuildConfig struct {
	Bin        string   `toml:"bin"`
	Cmd        string   `toml:"cmd"`
	Log        string   `toml:"log"`
	ExcludeDir []string `toml:"exclude_dir"`
}

func NewCudairConfig(path string) (*CudairConfig, error) {
	var config CudairConfig
	if _, err := toml.DecodeFile(path, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
