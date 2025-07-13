---
title: Pages Functions
url: https://edgeone.cloud.tencent.com/pages/document/162936866445025280
publishedTime: null
---

### 概述

Pages 边缘函数提供了 EdgeOne 边缘节点的 Serverless 代码执行环境，您只需编写业务函数代码并设置触发规则，无需配置和管理服务器等基础设施，即可在靠近用户的边缘节点上弹性、安全地运行代码。

![](https://write-document-release-1258344699.cos.ap-guangzhou.myqcloud.com/100026466949%2F2daf511ab12711ef96e55254002693fd.png)

Loading…

### 优势

**分布式部署**

EdgeOne 拥有超过 3200+ 边缘节点，边缘函数以分布式部署的方式运行在边缘节点。

**超低延迟**

客户端请求将自动被调度至靠近您用户最近的边缘节点上，命中触发规则触发边缘函数对请求进行处理并响应结果给客户端，可显著降低客户端的访问时延。

**弹性扩容**

边缘函数可以根据客户端请求数的突增，由近及远的将请求调度至有充足计算资源的边缘节点处理，您无需担忧突峰场景。

**Serverless 架构**

您无需再关心和维护底层服务器的内存、CPU、网络和其他基础设施资源，可以挪出精力更专注业务代码的开发。

### 快速指引

1\. 安装 npm 包：`npm install -g edgeone`，更多命令详见脚手架文档

2\. 本地开发：在 Pages 代码项目下

2.1 函数初始化：`edgeone pages init`，自动初始化 functions 目录，承载 functions 代码

2.2 关联项目：`edgeone pages link`，填入当前项目名称，自动关联项目 KV 配置、环境变量等信息

2.3 本地开发：`edgeone pages dev`，启动本地代理服务，进行函数调试

3\. 函数发布：代码推送到远端仓库，自动构建发布函数

### 路由

Pages Functions 基于 `/functions` 目录结构生成访问路由。您可在项目仓库 /functions 目录下创建任意层级的子目录，参考下述示例。

```bash
...functions├── index.js├── hello-pages.js├── helloworld.js├── api    ├── users      ├── list.js      ├── geo.js      ├── [id].js    ├── visit      ├── index.js    ├── [[default]].js...
```

上述目录文件结构，经 EdgeOne Pages 平台构建后将生成以下路由。这些路由将 Pages URL 映射到 `/functions` 文件，当客户端访问 URL 时将触发对应的文件代码被运行：

<table data-slate-node="element" class="readonly fixed"><colgroup contenteditable="false"><col width="45%"><col width="55%"></colgroup><tbody><tr data-slate-node="element" class=""><td data-slate-node="element" class="selectable-area"><p><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">文件路径</span></span></span></p></td><td data-slate-node="element" class="selectable-area"><p><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">路由</span></span></span></p></td></tr><tr data-slate-node="element" class=""><td data-slate-node="element" class="selectable-area"><p><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">/functions/index.js</span></span></span></p></td><td data-slate-node="element" class="selectable-area"><p><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">example.com/</span></span></span></p></td></tr><tr data-slate-node="element" class=""><td data-slate-node="element" class="selectable-area"><p><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">/functions/hello-pages.js</span></span></span></p></td><td data-slate-node="element" class="selectable-area"><p><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">example.com/hello-pages</span></span></span></p></td></tr><tr data-slate-node="element" class=""><td data-slate-node="element" class="selectable-area"><p><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">/functions/helloworld.js</span></span></span></p></td><td data-slate-node="element" class="selectable-area"><p><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">example.com/helloworld</span></span></span></p></td></tr><tr data-slate-node="element" class=""><td data-slate-node="element" class="selectable-area"><p><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">/functions/api/users/list.js</span></span></span></p></td><td data-slate-node="element" class="selectable-area"><p><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">example.com/api/users/list</span></span></span></p></td></tr><tr data-slate-node="element" class=""><td data-slate-node="element" class="selectable-area"><p><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">/functions/api/users/geo.js</span></span></span></p></td><td data-slate-node="element" class="selectable-area"><p><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">example.com/api/users/geo</span></span></span></p></td></tr><tr data-slate-node="element" class=""><td data-slate-node="element" class="selectable-area"><p><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">/functions/api/users/[id].js</span></span></span></p></td><td data-slate-node="element" class="selectable-area"><p><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">example.com/api/users/1024</span></span></span></p></td></tr><tr data-slate-node="element" class=""><td data-slate-node="element" class="selectable-area"><p><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">/functions/api/visit/index.js</span></span></span></p></td><td data-slate-node="element" class="selectable-area"><p><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">example.com/api/visit</span></span></span></p></td></tr><tr data-slate-node="element" class=""><td data-slate-node="element" class="selectable-area"><p><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">/functions/api/[[default]].js</span></span></span></p></td><td data-slate-node="element" class="selectable-area"><p><span contenteditable="false" class="tse-ul-symbol"></span><span class="tse-ul-content"><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">example.com/api/books/list</span></span></span></span></p><p><span contenteditable="false" class="tse-ul-symbol"></span><span class="tse-ul-content"><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">example.com/api/books/1024</span></span></span></span></p><p><span contenteditable="false" class="tse-ul-symbol"></span><span class="tse-ul-content"><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">example.com/api/...</span></span></span></span></p></td></tr></tbody></table>

**说明：**

路由尾部斜杠 / 是可选。`/hello-pages` 和 `/hello-pages/` 将被路由到 /functions/hello-pages.js。

如果没有匹配到 Pages Functions 路由，客户端请求将被路由到 Pages 对应的静态资源。

路由大小写敏感，/helloworld 将被路由到 /functions/helloworld.js，不能被路由到 /functions/HelloWorld.js

**动态路由**

Pages Functions 支持动态路由，上述示例中一级动态路径 /functions/api/users/\[id\].js，多级动态路径 /functions/api/\[\[default\]\].js。参考下述用法：

<table data-slate-node="element" class="readonly fixed"><colgroup contenteditable="false"><col width="33.33%"><col width="33.33%"><col width="33.34%"></colgroup><tbody><tr data-slate-node="element" class=""><td data-slate-node="element" class="selectable-area"><p><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">文件路径</span></span></span></p></td><td data-slate-node="element" class="selectable-area"><p><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">路由</span></span></span></p></td><td data-slate-node="element" class="selectable-area"><p><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">匹配</span></span></span></p></td></tr><tr data-slate-node="element" class=""><td data-slate-node="element" class="selectable-area" colspan="1" rowspan="3"><p><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">/functions/api/users/[id].js</span></span></span></p></td><td data-slate-node="element" class="selectable-area"><p><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">example.com/api/users/1024</span></span></span></p></td><td data-slate-node="element" class="selectable-area"><p><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">是</span></span></span></p></td></tr><tr data-slate-node="element" class=""><td data-slate-node="element" class="hidden selectable-area" colspan="0" rowspan="0"></td><td data-slate-node="element" class="selectable-area"><p><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">example.com/api/users/vip/1024</span></span></span></p></td><td data-slate-node="element" class="selectable-area"><p><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">否</span></span></span></p></td></tr><tr data-slate-node="element" class=""><td data-slate-node="element" class="hidden selectable-area" colspan="0" rowspan="0"></td><td data-slate-node="element" class="selectable-area"><p><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">example.com/api/vip/1024</span></span></span></p></td><td data-slate-node="element" class="selectable-area"><p><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">否</span></span></span></p></td></tr><tr data-slate-node="element" class=""><td data-slate-node="element" class="selectable-area" colspan="1" rowspan="3"><p><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">/functions/api/[[default]].js</span></span></span></p></td><td data-slate-node="element" class="selectable-area"><p><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">example.com/api/books/list</span></span></span></p></td><td data-slate-node="element" class="selectable-area"><p><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">是</span></span></span></p></td></tr><tr data-slate-node="element" class=""><td data-slate-node="element" class="hidden selectable-area" colspan="0" rowspan="0"></td><td data-slate-node="element" class="selectable-area"><p><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">example.com/api/1024</span></span></span></p></td><td data-slate-node="element" class="selectable-area"><p><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">是</span></span></span></p></td></tr><tr data-slate-node="element" class=""><td data-slate-node="element" class="hidden selectable-area" colspan="0" rowspan="0"></td><td data-slate-node="element" class="selectable-area"><p><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">example.com/v2/vip/1024</span></span></span></p></td><td data-slate-node="element" class="selectable-area"><p><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">否</span></span></span></p></td></tr></tbody></table>

### Function Handlers

使用 Functions Handlers 可为 Pages 创建自定义请求处理程序，以及定义 RESTful API 实现全栈应用。支持下述的 Handlers 方法：

<table data-slate-node="element" class="readonly fixed"><colgroup contenteditable="false"><col width="65%"><col width="35%"></colgroup><tbody><tr data-slate-node="element" class=""><td data-slate-node="element" class="selectable-area"><p><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">Handlers 方法</span></span></span></p></td><td data-slate-node="element" class="selectable-area"><p><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">描述</span></span></span></p></td></tr><tr data-slate-node="element" class=""><td data-slate-node="element" class="selectable-area"><p><span data-slate-node="text"><span data-slate-leaf="true"><code><span data-slate-string="true">onRequest</span></code></span></span><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">(context: EventContext): Response | Promise&lt;Response&gt;</span></span></span></p></td><td data-slate-node="element" class="selectable-area"><p><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">匹配 HTTP Methods</span></span></span></p><p><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">(</span></span></span><span data-slate-node="text"><span data-slate-leaf="true"><code><span data-slate-string="true">GET</span></code></span></span><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">, </span></span></span><span data-slate-node="text"><span data-slate-leaf="true"><code><span data-slate-string="true">POST</span></code></span></span><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">, </span></span></span><span data-slate-node="text"><span data-slate-leaf="true"><code><span data-slate-string="true">PATCH</span></code></span></span><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">, </span></span></span><span data-slate-node="text"><span data-slate-leaf="true"><code><span data-slate-string="true">PUT</span></code></span></span><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">, </span></span></span><span data-slate-node="text"><span data-slate-leaf="true"><code><span data-slate-string="true">DELETE</span></code></span></span><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">, </span></span></span><span data-slate-node="text"><span data-slate-leaf="true"><code><span data-slate-string="true">HEAD</span></code></span></span><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">, </span></span></span><span data-slate-node="text"><span data-slate-leaf="true"><code><span data-slate-string="true">OPTIONS</span></code></span></span><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">)</span></span></span></p></td></tr><tr data-slate-node="element" class=""><td data-slate-node="element" class="selectable-area" colspan="1" rowspan="1"><p><span data-slate-node="text"><span data-slate-leaf="true"><code><span data-slate-string="true">onRequestGet</span></code></span></span><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">(context: EventContext): Response | Promise&lt;Response&gt;</span></span></span></p></td><td data-slate-node="element" class="selectable-area" colspan="1" rowspan="1"><p><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">匹配 HTTP Methods (</span></span></span><span data-slate-node="text"><span data-slate-leaf="true"><code><span data-slate-string="true">GET</span></code></span></span><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">)</span></span></span></p></td></tr><tr data-slate-node="element" class=""><td data-slate-node="element" class="selectable-area" colspan="1" rowspan="1"><p><span data-slate-node="text"><span data-slate-leaf="true"><code><span data-slate-string="true">onRequestPost</span></code></span></span><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">(context: EventContext): Response | Promise&lt;Response&gt;</span></span></span></p></td><td data-slate-node="element" class="selectable-area" colspan="1" rowspan="1"><p><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">匹配 HTTP Methods (</span></span></span><span data-slate-node="text"><span data-slate-leaf="true"><code><span data-slate-string="true">POST</span></code></span></span><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">)</span></span></span></p></td></tr><tr data-slate-node="element" class=""><td data-slate-node="element" class="selectable-area" colspan="1" rowspan="1"><p><span data-slate-node="text"><span data-slate-leaf="true"><code><span data-slate-string="true">onRequestPatch</span></code></span></span><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">(context: EventContext): Response | Promise&lt;Response&gt;</span></span></span></p></td><td data-slate-node="element" class="selectable-area" colspan="1" rowspan="1"><p><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">匹配 HTTP Methods (</span></span></span><span data-slate-node="text"><span data-slate-leaf="true"><code><span data-slate-string="true">PATCH</span></code></span></span><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">)</span></span></span></p></td></tr><tr data-slate-node="element" class=""><td data-slate-node="element" class="selectable-area" colspan="1" rowspan="1"><p><span data-slate-node="text"><span data-slate-leaf="true"><code><span data-slate-string="true">onRequestPut</span></code></span></span><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">(context: EventContext): Response | Promise&lt;Response&gt;</span></span></span></p></td><td data-slate-node="element" class="selectable-area" colspan="1" rowspan="1"><p><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">匹配 HTTP Methods (</span></span></span><span data-slate-node="text"><span data-slate-leaf="true"><code><span data-slate-string="true">PUT</span></code></span></span><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">)</span></span></span></p></td></tr><tr data-slate-node="element" class=""><td data-slate-node="element" class="selectable-area" colspan="1" rowspan="1"><p><span data-slate-node="text"><span data-slate-leaf="true"><code><span data-slate-string="true">onRequestDelete</span></code></span></span><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">(context: EventContext): Response | Promise&lt;Response&gt;</span></span></span></p></td><td data-slate-node="element" class="selectable-area" colspan="1" rowspan="1"><p><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">匹配 HTTP Methods (</span></span></span><span data-slate-node="text"><span data-slate-leaf="true"><code><span data-slate-string="true">DELETE</span></code></span></span><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">)</span></span></span></p></td></tr><tr data-slate-node="element" class=""><td data-slate-node="element" class="selectable-area" colspan="1" rowspan="1"><p><span data-slate-node="text"><span data-slate-leaf="true"><code><span data-slate-string="true">onRequestHead</span></code></span></span><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">(context: EventContext): Response | Promise&lt;Response&gt;</span></span></span></p></td><td data-slate-node="element" class="selectable-area" colspan="1" rowspan="1"><p><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">匹配 HTTP Methods (</span></span></span><span data-slate-node="text"><span data-slate-leaf="true"><code><span data-slate-string="true">HEAD</span></code></span></span><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">)</span></span></span></p></td></tr><tr data-slate-node="element" class=""><td data-slate-node="element" class="selectable-area" colspan="1" rowspan="1"><p><span data-slate-node="text"><span data-slate-leaf="true"><code><span data-slate-string="true">onRequestOptions</span></code></span></span><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">(context: EventContext): Response | Promise&lt;Response&gt;</span></span></span></p></td><td data-slate-node="element" class="selectable-area" colspan="1" rowspan="1"><p><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">匹配 HTTP Methods (</span></span></span><span data-slate-node="text"><span data-slate-leaf="true"><code><span data-slate-string="true">OPTIONS</span></code></span></span><span data-slate-node="text"><span data-slate-leaf="true"><span data-slate-string="true">)</span></span></span></p></td></tr></tbody></table>

**EventContext 对象描述**

context 是传递给 Function Handlers 方法的对象，包含下述属性：

params：动态路由 `/functions/api/users/[id].js` 参数值

```js
export function onRequestGet(context) {
  return new Response(`User id is ${context.params.id}`);
}
```

env：Pages 环境变量

### Runtime APIs

Pages Functions 基于[边缘函数](https://cloud.tencent.com/document/product/1552/81344)实现，提供了 EdgeOne 边缘节点的 Serverless 代码执行环境。支持 ES6 语法和标准的 Web Service Worker API。其中大部分 Runtime APIs 可参考[边缘函数](https://cloud.tencent.com/document/product/1552/81344)用法，参考下述描述：

**说明：**

当前 [EdgeOne CLI](https://edgeone.cloud.tencent.com/pages/document/162936923278893056) 调试环境中不支持使用 fetch 访问 EdgeOne 节点缓存或回源。

使用 context.request.eo 可获取客户端 [GEO](https://cloud.tencent.com/document/product/1552/81902#eo) 信息。

### 示例函数

**获取用户访问地理位置：**

```ts
export function onRequest({ request }) {
  const geo = request.eo.geo;
  const res = JSON.stringify({ geo: geo });
  return new Response(res, {
    headers: {
      "content-type": "application/json; charset=UTF-8",
      "Access-Control-Allow-Origin": "*",
    },
  });
}
```

**使用 KV 记录页面访问数：**

有关如何使用 KV 存储的详细信息，请参阅 [KV 存储](https://edgeone.cloud.tencent.com/pages/document/162936897742577664)。

```ts
export async function onRequest({ request, params, env }) {  // my_kv 是您在项目中绑定命名空间时的变量名  const visitCount = await my_kv.get('visitCount');  let visitCountInt = Number(visitCount);  visitCountInt += 1;  await my_kv.put('visitCount', visitCountInt.toString());  const res = JSON.stringify({    visitCount: visitCountInt,  });  return new Response(res, {    headers: {      'content-type': 'application/json; charset=UTF-8',      'Access-Control-Allow-Origin': '*',    },  });}
```

**连接 supabase 三方数据库：**

```ts
import { createClient } from "@supabase/supabase-js";
export async function onRequest({ request, params, env }) {
  const supabase = createClient(env.supabaseUrl, env.supabaseKey);
  let { data } = await supabase.from("users").select("*");
  return new Response(JSON.stringify({ users: data }), {
    headers: {
      "content-type": "application/json; charset=UTF-8",
      "Access-Control-Allow-Origin": "*",
    },
  });
}
```

