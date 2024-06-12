# Ant-plugin

- [Synopsis](#Synopsis)
- [Plugin Image](#Plugin-Image)
- [Building](#building)
- [Examples](#Examples)


## Synopsis

This plugin is used to build Java application using [Apache Ant](https://ant.apache.org). 


## Plugin Image

The plugin `Rakshitharness/drone-ant-plugin` is available for the following architectures:

| OS            | Tag                                |
| ------------- | ---------------------------------- |
| linux/amd64   | `linux-amd64`                      |
| linux/arm64   | `linux-arm64`                      |
| windows/amd64 | `windows-amd64`                    |


## Building

Build the plugin image:

```text
./scripts/build.sh
```

## Examples

```
# Plugin YAML
- step:
    type: Plugin
    name: ant-plugin-arm64
    identifier: ant-plugin-arm64
    spec:
        connectorRef: harness-docker-connector
        image: rakshitagar/drone-ant-plugin:linux-arm64
       

- step:
    type: Plugin
    name: ant-plugin-amd64
    identifier: ant-plugin-arm64
    spec:
        connectorRef: harness-docker-connector
        image: rakshitagar/drone-ant-plugin:linux-amd64
        