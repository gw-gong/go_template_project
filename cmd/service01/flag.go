package main

import (
	"context"
	"errors"
	"flag"
	"fmt"

	utils_comm "github.com/gw-gong/gwkit-go/utils/common"
)

var (
	cfgFilePath = flag.String("cfg_path", "../../config/local_config/service01", "config file path")
	cfgFileName = flag.String("cfg_name", "test.yaml", "config file name")
)

func validateFlags(ctx context.Context) {
	if *cfgFilePath == "" {
		flag.Usage()
		utils_comm.ExitOnErr(ctx, errors.New("cfg_path is required"))
	}
	if *cfgFileName == "" {
		flag.Usage()
		utils_comm.ExitOnErr(ctx, fmt.Errorf("cfg_name is required"))
	}
}
