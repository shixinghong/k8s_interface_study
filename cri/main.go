package main

import (
	"fmt"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/remote"
)

func main() {

	ref, err := name.ParseReference("hub.fastonetech.com/infra/nginx")
	if err != nil {
		panic(err)
	}

	desc, err := remote.Get(ref)
	if err != nil {
		panic(err)
	}

	if desc.MediaType.IsIndex() { // 如果是index模式 绝大多是都是index模式
		idx, err := desc.ImageIndex()
		if err != nil {
			panic(err)
		}

		mf, err := idx.IndexManifest()
		if err != nil {
			panic(err)
		}
		for _, v := range mf.Manifests {
			img, err := idx.Image(v.Digest)
			if err != nil {
				panic(err)
			}
			cf, err := img.ConfigFile()
			if err != nil {
				panic(err)
			}
			fmt.Println(cf.Config.Entrypoint, cf.Config.Cmd)
		}
	}

	if desc.MediaType.IsImage() {
		image, err := desc.Image()
		if err != nil {
			panic(err)
		}
		cf, err := image.ConfigFile()
		if err != nil {
			panic(err)
		}
		fmt.Println(cf.Config.Entrypoint, cf.Config.Cmd)
	}

}
