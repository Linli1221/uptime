---
title: 导入 Git 仓库
url: https://edgeone.cloud.tencent.com/pages/document/171193092266749952
publishedTime: null
---

在本章节中，您将通过导入 Git 仓库的方式，将项目部署到 Pages 平台。

**Pages 目前支持 Github，Gitee 等 Git 提供商的接入**，下面以 Github 的接入方式为例：

### 连接 Git 仓库

部署 Web 应用程序的第一步是连接您的 Git 仓库。Pages 与代码管理系统无缝集成，使开发工作流与部署过程之间能顺畅同步。

连接您的仓库：

注册/登录：您可以使用多种方式快速注册登录[腾讯云控制台](https://console.cloud.tencent.com/edgeone/pages)。

开始使用：首次访问时，点击页面上的“立即开通”。

绑定 Github：在控制台页面，点击“Github”以链接您的仓库。

![](https://write-document-release-1258344699.cos.ap-guangzhou.myqcloud.com/100026466949%2F5253d601322311f09dc05254005ef0f7.png)

Loading…

授权 Github：授予 EdgeOne 访问您仓库的权限。

![](https://write-document-release-1258344699.cos.ap-guangzhou.myqcloud.com/100027864926%2F05ff6b34103511f0854e525400454e06.png)

Loading…

选择仓库：选择您要部署的仓库或授权所有仓库。

![](https://write-document-release-1258344699.cos.ap-guangzhou.myqcloud.com/100027864926%2F2db22d99103511f0aaa3525400e889b2.png)

Loading…

### 自定义构建

连接仓库后，您需要填写构建配置。此步骤对于项目能否正确编译和顺利部署至关重要。

选择仓库：选择您要部署的仓库。

![](https://write-document-release-1258344699.cos.ap-guangzhou.myqcloud.com/100026466949%2F35e06bd4ff3e11ef8c825254001c06ec.png)

Loading…

输入构建命令：输入您的构建命令。如果不确定，请检查 package.json 中的 scripts 部分的值。

![](https://write-document-release-1258344699.cos.ap-guangzhou.myqcloud.com/100026466949%2Ff7c4159000a311f09f0a52540099c741.png)

Loading…

选择加速区域：不同的区域决定了项目分配的节点资源，以及添加自定义域名的时候是否需要备案，详情可参考 域名管理 - 加速区域。

![](https://write-document-release-1258344699.cos.ap-guangzhou.myqcloud.com/100026466949%2Ffb2b21d000a311f0aa645254007c27c5.png)

Loading…

### 全球部署

连接仓库并填写构建配置后，您就可以准备将应用程序部署至全球。

启动部署：

检查您的配置项。

点击“开始部署”。

![](https://write-document-release-1258344699.cos.ap-guangzhou.myqcloud.com/100026466949%2Ffeb924f500a311f09c4a5254001c06ec.png)

Loading…

EdgeOne Pages 将自动构建您的项目并将其部署到全球边缘网络。

如下图表示部署成功！

![](https://write-document-release-1258344699.cos.ap-guangzhou.myqcloud.com/100026466949%2F919edc95ff3f11ef9f695254007c27c5.png)

Loading…

当新的提交推送到部署分支时，EdgeOne 将自动拉取并部署最新的提交。

![](https://write-document-release-1258344699.cos.ap-guangzhou.myqcloud.com/100026466949%2F1ea7f4fcf8c411ef9f695254007c27c5.png)

Loading…

通过以上步骤，您可以通过其强大的全球边缘网络和简化的工作流程，快速轻松地使用 EdgeOne Pages 部署 Web 应用程序。

### 如何处理意外情况

如果部署的版本与仓库不同，请确保拉取最新版本。

如有其他问题，请扫描右上角的“开发者沟通群”二维码，加群与我们联系。

![](https://write-document-release-1258344699.cos.ap-guangzhou.myqcloud.com/100026466949%2F416a85ef322311f0847a525400454e06.png)

Loading…

