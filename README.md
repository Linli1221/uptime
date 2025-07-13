# Uptime Monitor

一个前后端分离的网站监控服务，用于监控网站的可用性和响应时间。

## 架构

- **后端**: Go + Gin + GORM + SQLite/MySQL + Redis
- **前端**: Next.js + TypeScript + Tailwind CSS + SWR
- **部署**: 
  - 后端: Docker 容器平台
  - 前端: EdgeOne Pages (静态托管)

## 功能特性

- ✅ 网站可用性监控
- ✅ 实时状态显示
- ✅ 添加/删除监控目标
- ✅ 响应式UI设计
- ✅ 自动定时检查
- ✅ 状态历史记录

## 本地开发

### 后端启动

```bash
cd backend
cp .env.example .env
# 编辑 .env 文件配置数据库
go mod download
go run main.go
```

### 前端启动

```bash
cd frontend
npm install
npm run dev
```

## 生产部署

### 后端部署 (Docker)

1. 构建 Docker 镜像：
```bash
cd backend
docker build -t uptime-backend .
```

2. 运行容器：
```bash
docker run -d \
  -p 8080:8080 \
  -e DB_DIALECT=mysql \
  -e DB_HOST=your-mysql-host \
  -e DB_PORT=3306 \
  -e DB_USER=your-username \
  -e DB_PASSWORD=your-password \
  -e DB_NAME=uptime \
  -e REDIS_ADDR=your-redis-host:6379 \
  -e REDIS_PASSWORD=your-redis-password \
  -e REDIS_DB=0 \
  --name uptime-backend \
  uptime-backend
```

### 前端部署 (EdgeOne Pages)

1. 配置环境变量：
```bash
cd frontend
cp .env.example .env.local
# 编辑 .env.local，设置 NEXT_PUBLIC_API_BASE_URL 为后端 API 地址
```

2. 构建静态文件：
```bash
npm run build
```

3. 将 `out` 目录中的文件上传到 EdgeOne Pages

## 环境变量

### 后端环境变量

- `DB_DIALECT`: 数据库类型 (sqlite/mysql)
- `DB_HOST`: MySQL 主机地址
- `DB_PORT`: MySQL 端口
- `DB_USER`: MySQL 用户名
- `DB_PASSWORD`: MySQL 密码
- `DB_NAME`: 数据库名
- `REDIS_ADDR`: Redis 地址
- `REDIS_PASSWORD`: Redis 密码
- `REDIS_DB`: Redis 数据库编号

### 前端环境变量

- `NEXT_PUBLIC_API_BASE_URL`: 后端 API 基础地址

## API 接口

### 获取所有监控

```
GET /api/monitors
```

### 创建监控

```
POST /api/monitors
Content-Type: application/json

{
  "name": "Google",
  "url": "https://google.com",
  "interval": 60,
  "status": "pending"
}
```

### 获取单个监控

```
GET /api/monitors/:id
```

### 更新监控

```
PUT /api/monitors/:id
```

### 删除监控

```
DELETE /api/monitors/:id
```

## 技术栈

### 后端

- **Go 1.21+**: 主要编程语言
- **Gin**: Web 框架
- **GORM**: ORM 框架
- **SQLite/MySQL**: 数据存储
- **Redis**: 缓存和会话存储
- **Docker**: 容器化部署

### 前端

- **Next.js 15**: React 框架
- **TypeScript**: 类型安全
- **Tailwind CSS**: 样式框架
- **SWR**: 数据获取
- **EdgeOne Pages**: 静态托管

## 许可证

MIT License