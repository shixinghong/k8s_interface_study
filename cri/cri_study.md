## CRI (container runtime interface) 容器运行时接口
### 接口规范
#### image spec(镜像标准)  [github链接](https://github.com/opencontainers/image-spec)
###### 规范的组成部分包括:
  - Image Manifest: 定义image的清单 描述组成容器image的组件 [详细链接](https://github.com/opencontainers/image-spec/blob/main/manifest.md)
    ```shell
    cat alpine/manifest.json
    [
      {
        "Config": "c059bfaa849c4d8e4aecaeb3a10c2d9b3d85f5165c66ad3a4d937758128c4d18.json",
        "RepoTags": [
          "alpine:3.15",
          "alpine:latest"
        ],
        "Layers": [
          "25fe74d3f1f6ccd36452f82043bd02b6b0ce82b6efaece7ee3b8fee9ef1acdc6/layer.tar"
        ]
      },
      {
        "Config": "0a97eee8041e2b6c0e65abb2700b0705d0da5525ca69060b9e0bde8a3d17afdb.json",
        "RepoTags": null,
        "Layers": [
          "a9c4799b6f93649010503ff965e25d1613b50098fc7bbdefeadf8dcded7db226/layer.tar"
        ]   
      }
    ]
    ```

  - Image Index  带注释的image清单索引 [详细链接](https://github.com/opencontainers/image-spec/blob/main/image-index.md) 
    ```shell
    
    ```
  - Image Layout 
  - Filesystem Layer 
  - Image Configuration 
  - Conversion
  - Descriptor 