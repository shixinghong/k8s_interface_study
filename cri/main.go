package main

import (
	"fmt"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/remote"
)

func main() {

	ref, err := name.ParseReference("xxxxxx/infra/alpine")
	if err != nil {
		panic(err)
	}

	desc, err := remote.Get(ref)
	if err != nil {
		panic(err)
	}

	if desc.MediaType.IsImage() {
		fmt.Println("image")
	}
	image, err := desc.Image()
	if err != nil {
		panic(err)
	}

	cfg, err := image.ConfigFile() // 获取image中的config信息
	if err != nil {
		panic(err)
	}
	fmt.Println(cfg.OS, cfg.Architecture, cfg.Config.Cmd, cfg.Config.Entrypoint)
}
