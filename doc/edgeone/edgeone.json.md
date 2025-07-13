---
title: edgeone.json
url: https://edgeone.cloud.tencent.com/pages/document/162936771610066944
publishedTime: null
---

除了通过控制台进行项目设置，您还可以在项目根目录中创建 `edgeone.json` 文件。该文件允许您定义和覆盖项目的默认行为，以便更灵活地配置项目。

配置文件包括以下设置项：

#### buildCommand

可用于覆盖 控制台 - 项目设置- 构建部署配置 中的编译命令。

```json
{ "buildCommand": "next build" }
```

#### installCommand

可用于覆盖 控制台 - 项目设置- 构建部署配置 中的安装命令。此配置项可自定义构建过程使用的[包管理器](https://edgeone.cloud.tencent.com/pages/document/162936788693114880#f777c5cb-1483-47fb-85bc-7c610686105d)。

```json
{ "installCommand": "npm install" }
```

#### outputDirectory

可用于覆盖 控制台 - 项目设置- 构建部署配置 中的输出目录。

```json
{ "outputDirectory": "./build" }
```

#### nodeVersion

可用于指定构建环境的 node 版本。建议使用系统预装的 14.21.3、16.20.2、18.20.4、20.18.0、22.11.0 几个 node 版本，若填写其他版本号，可能导致部署失败。

```json
{ "nodeVersion": "22.11.0" }
```

#### redirects

可用于将请求从一个 URL 重定向到另一个 URL，以下是几个重定向的示例。

此示例使用 301 永久重定向将请求从 `/articles/+ ID` 的 URL（例如 /articles/123），重定向到 `/news-articles/ + ID` 的 URL（例如 /news-articles/123）：

```json
{
  "redirects": [
    {
      "source": "/articles/:id",
      "destination": "/news-articles/:id",
      "statusCode": 301
    }
  ]
}
```

此示例使用 302 临时重定向将请求从 `/old-path` ，重定向到 `/new-path` ：

```json
{
  "redirects": [
    { "source": "/old-path", "destination": "/new-path", "statusCode": 302 }
  ]
}
```

此示例使用 301 永久重定向将请求从 `/template-source` ，重定向到外部站点的绝对路径 `https://github.com/TencentEdgeOne/pages-templates/tree/main/examples/chrome-ai` ：

```json
{
  "redirects": [
    {
      "source": "/template-source",
      "destination": "https://github.com/TencentEdgeOne/pages-templates/tree/main/examples/chrome-ai",
      "statusCode": 301
    }
  ]
}
```

此示例使用 301 永久重定向将非 www 请求重定向到 www，同时也支持反向重定向**（仅对自定义域名生效）**：

```json
{
  "redirects": [
    { "source": "$host", "destination": "$wwwhost", "statusCode": 301 }
  ]
}
```

**注意：**

redirect 最大数量限制为 30 个

source 与 destination 不超过 500 个字符

#### rewrites

此示例将所有以 /assets/ 开头的请求重写到 /assets-new/ 目录下，同时保留原始请求的路径部分。

```json
{ "rewrites": [{ "source": "/assets/*", "destination": "/assets-new/:splat" }] }
```

我们还可以进一步细化重写规则，专门针对 PNG 格式的图像文件进行处理。以下示例将确保所有以 .png 结尾的请求都被重写到新的路径，同时保留文件名。

```json
{
  "rewrites": [
    { "source": "/assets/*.png", "destination": "/assets-new/:splat.png" }
  ]
}
```

**注意：**

rewrite 最大数量限制为 30 个

source 与 destination 不超过 500 个字符

该配置仅适用于静态资源访问

不支持 SPA (单页应用) 的前端路由重写

源路径必须以 \`/\` 开头

**SPA 应用重写建议**

如需在 SPA 中实现 URL 重写，建议采用以下方案：

前端路由重定向

使用框架自带的路由系统进行路径重定向

在路由配置中定义重写规则

#### headers

可用于自定义和管理 HTTP 响应头，以改善网站的性能和安全性，同时提升用户体验。

此示例通过为所有请求设置 X-Frame-Options 头部来增强网站的安全性，防止点击劫持攻击；同时，通过 Cache-Control 指定响应可以被缓存 2 小时，以提高性能和减少服务器负担。

```json
{
  "headers": [
    {
      "source": "/*",
      "headers": [
        { "key": "X-Frame-Options", "value": "DENY" },
        { "key": "Cache-Control", "value": "max-age=7200" }
      ]
    }
  ]
}
```

我们还可以进一步优化特定资源的缓存策略，特别是针对 /assets/ 目录下的静态资源，此示例将为该目录下的所有文件设置更长的缓存时间。

```json
{
  "headers": [
    {
      "source": "/assets/*",
      "headers": [{ "key": "Cache-Control", "value": "max-age=31536000" }]
    }
  ]
}
```

**注意：**

header 最大数量限制为 30 个

每个 header 的 key 限制为 1 - 100 个字符，允许使用数字、字母及特殊符号 '-'

每个 header 的 value 限制为 1 - 1000 个字符，不支持中文

#### caches

可用于针对不同资源配置边缘缓存时间，优化不同资源的边缘缓存策略，提升请求资源的加载速度。

此示例将 images 目录下所有文件资源设置为缓存 1 天。

```json
 "caches": [    {      "source": "/images/*",      "cacheTtl": 86400    }  ]
```

我们还可以针对特定文件设置缓存，此示例将 sitemap.xml 文件设置为不缓存， 将 images 下所有 jpg 文件设置为缓存 3600 秒。

```json
"caches": [    {      "source": "/sitemap.xml",      "cacheTtl": 0    },    {      "source": "/images/*.jpg",      "cacheTtl": 3600    }  ]
```

**注意：**

cacheTtl 以秒为单位，不允许小数且不能小于 0，设置为 0 时不缓存

#### source 匹配规则说明

在配置 redirects、rewrites、 headers 和 caches 时，source 字段用于定义请求路径的匹配规则。以下是主要的匹配特性：

1.路径匹配

source 支持使用特定的模式来匹配请求路径。匹配规则会根据请求的 URL 进行解析。

2.通配符

使用星号（`*`）作为通配符，可以匹配路径中的任意字符。请注意，source 中只能包含一个通配符。

3.占位符

占位符以冒号（`:`）开头，后跟占位符名称。每个占位符只能在 source 中使用一次，并且会匹配除分隔符外的所有字符。

#### edgeone.json 文件示例

此示例展示了如何将多个设置组合在一个配置文件中，但并不涵盖所有可用选项。请注意，文件中的每个设置项并非必需。

```json
{
  "name": "example-app",
  "buildCommand": "next build",
  "installCommand": "npm install",
  "outputDirectory": "./build",
  "nodeVersion": "22.11.0",
  "redirects": [
    {
      "source": "/articles/:id",
      "destination": "/news-articles/:id",
      "statusCode": 301
    },
    { "source": "/old-path", "destination": "/new-path", "statusCode": 302 }
  ],
  "rewrites": [{ "source": "/assets/*", "destination": "/assets-new/:splat" }],
  "headers": [
    {
      "source": "/*",
      "headers": [
        { "key": "X-Frame-Options", "value": "DENY" },
        { "key": "Cache-Control", "value": "max-age=7200" }
      ]
    },
    {
      "source": "/assets/*",
      "headers": [{ "key": "Cache-Control", "value": "max-age=31536000" }]
    }
  ]
}
```

