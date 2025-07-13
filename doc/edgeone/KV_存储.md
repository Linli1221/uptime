---
title: KV 存储
url: https://edgeone.cloud.tencent.com/pages/document/162936897742577664
publishedTime: null
---

### 概述

EdgeOne Pages KV 是一种多边缘节点部署的 KV 持久化数据存储，允许用户在全球范围内读写数据。遵循最终一致性，并确保 60s 内全球同步访问。您可以结合 EdgeOne Pages 中的函数能力，构建需要存储小型数据的 API、网站等。

### 使用案例

#### 计数器

您可以根据业务需求，将某个按钮、页面的点击数，存储到 EdgeOne Pages KV，点击时更新 KV 中的键值记录，以便统计和分析点击或访问行为。

#### 密钥库

对于一些敏感性的信息，不方便放置在代码仓库中时，您可以将其存储到 Edgeone Pages KV 中，动态获取以便更好的保护您的数据。

#### 购物车

Edgeone Pages KV 可以跨多个页面、多个终端地保留用户数据，可以助您便捷实现用户购物车、用户订单等简单的业务需求。

### 基础概念

#### 账户

Edgeone Pages KV 账户是 Pages 业务中 KV 用量统计和计费的最小单位。一个 EdgeOne Pages 账户对应 一个 KV 账户，用户在控制台页面开通后创建。账户存储容量 100MB。

#### 命名空间

命名空间（namespace）是 Edgeone Pages KV 中数据隔离的基本单位。每个命名空间可以看作是一个独立的 database，各个命名空间之间的数据互不干扰。一个账户可以创建 10 个命名空间。

#### 键值对

键值对是一种用户存储数据的结构，其中每个数据项由两个部分组成：键（Key）和值（Value）。键是唯一的标识符，用户标识值，值是与键相关联的数据。

#### 变量名

变量名是绑定项目跟命名空间时定义的运行时环境变量名称。 在使用命名空间数据之前，需要先将命名空间跟 Edgeone Pages 项目进行绑定。命名空间跟项目的绑定关系为多对多，在绑定时通过不同的运行时环境变量名称来区分使用。

### 快速开始

#### 初始化账户

1\. **进入控制台**：切换到“KV存储”页面。

2\. **申请开通账户**：点击“立即申请”按钮。

**说明：**

当前 KV 存储开放名额有限，请申请开通，我们会尽快处理，申请成功后即可使用。

![](https://write-document-release-1258344699.cos.ap-guangzhou.myqcloud.com/100026466949%2Fae13c9f5b6ed11efb5b052540055f650.png)

Loading…

#### 创建命名空间

1\. **进入命名空间列表**：开通 KV 账户后，在“KV存储”页面，点击“创建命名空间”按钮。

![](https://write-document-release-1258344699.cos.ap-guangzhou.myqcloud.com/100026466949%2F1c232126b3ab11ef8c01525400fdb830.png)

Loading…

2\. **创建命名空间：**输入需要创建的命名空间名称后，点击“创建”按钮，等待创建完成。

![](https://write-document-release-1258344699.cos.ap-guangzhou.myqcloud.com/100026466949%2F3b7c0c6fb12711ef8c01525400fdb830.png)

Loading…

3\. **创建完成：**创建完成后，可以在命名空间列表页，查看创建好的命名空间。

![](https://write-document-release-1258344699.cos.ap-guangzhou.myqcloud.com/100026466949%2F36accaf4b3a711ef96e55254002693fd.png)

Loading…

#### 添加记录

1\. **进入命名空间详情：**在命名空间列表中，点击对应命名空间名称，进入命名空间详情。

![](https://write-document-release-1258344699.cos.ap-guangzhou.myqcloud.com/100026466949%2F4a2e5865b3ab11ef96e55254002693fd.png)

Loading…

2\. **添加记录：**点击“新建记录”按钮。

![](https://write-document-release-1258344699.cos.ap-guangzhou.myqcloud.com/100026466949%2F3b7b7288b12711efbfb3525400bdab9d.png)

Loading…

3\. **创建完成：**创建完成后，可以在记录列表中查看添加好的记录。

![](https://write-document-release-1258344699.cos.ap-guangzhou.myqcloud.com/100026466949%2Fafbdb25ab3ab11ef8c01525400fdb830.png)

Loading…

#### 绑定命名空间

命名空间与项目间的绑定，在“项目”和“KV存储”中，都可以进行绑定。

##### KV 储存中绑定

1\. **进入关联项目：**在命名空间详情页，点击“关联项目”选项卡。

![](https://write-document-release-1258344699.cos.ap-guangzhou.myqcloud.com/100026466949%2F1159c2dab3ac11ef9dc0525400329841.png)

Loading…

2\. **绑定项目：**点击“绑定项目”按钮。

![](https://write-document-release-1258344699.cos.ap-guangzhou.myqcloud.com/100026466949%2F3b7d9781b12711ef9d2952540055f650.png)

Loading…

3\. **绑定成功：**绑定成功后，在绑定项目列表可以查看绑定好的项目。

![](https://write-document-release-1258344699.cos.ap-guangzhou.myqcloud.com/100026466949%2F5ab240f6b3ac11ef96e55254002693fd.png)

Loading…

##### 项目中绑定

1\. **进入项目-KV 存储：**进入项目详情后，点击“KV 存储”菜单。

![](https://write-document-release-1258344699.cos.ap-guangzhou.myqcloud.com/100026466949%2F02f0e57db61b11ef9dc0525400329841.png)

Loading…

2\. **绑定命名空间：**点击“绑定命名空间”按钮。

![](https://write-document-release-1258344699.cos.ap-guangzhou.myqcloud.com/100026466949%2Fce67cd00b3ad11ef96e55254002693fd.png)

Loading…

3\. **绑定成功：**绑定成功后，在绑定命名空间列表可以查看绑定好的命名空间。

![](https://write-document-release-1258344699.cos.ap-guangzhou.myqcloud.com/100026466949%2F3b82736fb12711ef9dc0525400329841.png)

Loading…

### 使用 KV

**说明：**

以下示例的 `my_kv` 是您在项目中绑定命名空间时的变量名，如上图的 `varname1`。

目前提供了一组 KV 存储操作方法供 Edgeone Pages 函数使用，具体方法如下：

#### put

写入 KV 数据，用户创建新的键值对后者更新已有键值对的值。

```jsx
put(key: string,value: string | ArrayBuffer | ArrayBufferView | ReadableStream): Promise<void>;
```

**参数**

\- **key**：需要创建或更新的键，长度小于等于 512 B。

\- **value**：需要写入的数据，数据长度小于等于 25 MB。

**返回值**

返回一个 `Promise`，需要 `await` 该 `Promise` 以验证写入是否成功。

**使用示例**

```jsx
await my_kv.put(key, value);
```

#### get

根据特定的 key 读取 KV 中的数据，并根据指定类型返回数据。

```jsx
get(key:string, object?:{type:string}): Promise<value:string|object|ArrayBuffer|ReadableStream>
```

**参数**

\- **key**：指定获取数据的 key。

\- **type**：用于指定返回的 value 类型，默认为 text。

\- **text**：`String`，将 value 转成 string 的形式返回。

\- **json**：`Object`，用户将 json 反序列化为 object 形式返回。

\- **arrayBuffer**：`ArrayBuffer`，用户将二进制的 value 转换为 ArrayBuffer 返回。

\- **stream**：`ReadableStream`，通常用于 value 较大的场景。

**返回值**

返回一个 `Promise`，需要 `await` 该 `Promise` 以获取 value。如果 key 不存在或 value 为空时，则返回 null。

**使用示例**

```jsx
let value = await my_kv.get(key);
let value = await my_kv.get(key, "json");
let value = await my_kv.get(key, { type: "json" });
```

#### delete

从 KV 中删除指定的 key。

```jsx
delete(key: string): Promise<void>;
```

**参数**

\- **key** : 需要删除的键值对的 key。

**返回值**

返回一个 `Promise`，需要 `await` 该 `Promise` 以验证删除是否成功。

**使用示例**

#### list

用于遍历 KV 中的所有 key。

```jsx
list({prefix?: string, limit?: number, cursor?: string}): Promise<ListResult>;
```

**参数**

\- **prefix**：用于过滤指定前缀的键。默认为空时，按照字典序排序返回。

\- **limit**：返回的 key 的最大数量，用于分页。默认值为 256，上限为 256。

\- **cursor**：游标，从指定键开始遍历，用于分页。默认为空。

**返回值**

返回一个 `Promise`，需要 `await` 该 `Promise` 以获取 ListResult：

```jsx
class ListResult {    complete: Boolean;    cursor: String;    keys: Array<ListKey>;}class ListKey {    key: String;}
```

\- **complete**：标记 list 操作是否完成，true 完成，false 未完成。

\- **cursor**：游标，为下一页的第一个 key，当 list 遍历完成时为 null。

\- **keys**：描述每个 key 的 object 数组。

**使用示例**

```jsx
// 通过 list 获取部分数据let result = await kv.list({ "prefix": "a", "limit": 10, "cursor": "abc" });// 通过 list 遍历所有数据let result;let cursor;do {    result = await kv.list(options);    cursor = result.cursor;} while (result && result.complete);
```

### KV 示例

```jsx
export async function onRequest({ request, params, env }) {    // 获取变量名为 my_kv 的命名空间 key    let count = await my_kv.get('count');    count = Number(count) + 1;    // 重新写入 visitCount 键值    await my_kv.put('count', String(count));    return new Response("ok", {});}
```

### KV 模板

﻿[计数器模板](https://functions-kv.edgeone.site/): 可以让用户记录和显示项目业务需求发生次数，可以用于统计“点击次数”、“博客访问量” 等业务场景。

