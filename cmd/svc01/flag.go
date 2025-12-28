package main

import (
	"errors"
	"flag"
	"path/filepath"
)

const (
	RootPath           = "../../"
	defaultCfgFilePath = "config/svc01/localcfg"
	defaultCfgFileName = "test.yaml"
)

func initFlags() (string, string, error) {
	flagCfgFilePath := flag.String("cfg_path", defaultCfgFilePath, "config file path")
	flagCfgFileName := flag.String("cfg_name", defaultCfgFileName, "config file name")

	flag.Parse()

	if *flagCfgFilePath == "" {
		return "", "", errors.New("cfg_path is required")
	}
	if *flagCfgFileName == "" {
		return "", "", errors.New("cfg_name is required")
	}
	return filepath.Join(RootPath, *flagCfgFilePath), *flagCfgFileName, nil
}
