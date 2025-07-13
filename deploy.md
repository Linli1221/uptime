# 部署指南

## 后端部署（Docker容器平台）

### 方式一：使用 Docker Compose（推荐）

1. 确保安装了 Docker 和 Docker Compose
2. 克隆项目代码
3. 运行部署命令：

```bash
# 构建并启动所有服务
docker-compose up -d

# 查看服务状态
docker-compose ps

# 查看日志
docker-compose logs -f backend
```

### 方式二：单独部署后端

1. 构建 Docker 镜像：
```bash
cd backend
docker build -t uptime-backend .
```

2. 运行容器：
```bash
docker run -d \
  --name uptime-backend \
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
  uptime-backend
```

## 前端部署（EdgeOne Pages）

### 步骤 1：构建静态文件

```bash
cd frontend

# 设置生产环境 API 地址
echo "NEXT_PUBLIC_API_BASE_URL=https://your-backend-domain.com" > .env.local

# 安装依赖
npm install

# 构建
npm run build
```

### 步骤 2：上传到 EdgeOne Pages

1. 登录腾讯云 EdgeOne 控制台
2. 进入 Pages 服务
3. 创建新项目
4. 上传 `frontend/out` 目录中的所有文件

### 步骤 3：配置域名（可选）

1. 在 EdgeOne Pages 中配置自定义域名
2. 设置 SSL 证书
3. 配置 CDN 加速

## 环境变量配置

### 后端环境变量

| 变量名 | 描述 | 示例值 |
|--------|------|--------|
| DB_DIALECT | 数据库类型 | mysql |
| DB_HOST | 数据库主机 | localhost |
| DB_PORT | 数据库端口 | 3306 |
| DB_USER | 数据库用户名 | uptime |
| DB_PASSWORD | 数据库密码 | password123 |
| DB_NAME | 数据库名 | uptime |
| REDIS_ADDR | Redis 地址 | localhost:6379 |
| REDIS_PASSWORD | Redis 密码 | redispass |
| REDIS_DB | Redis 数据库编号 | 0 |

### 前端环境变量

| 变量名 | 描述 | 示例值 |
|--------|------|--------|
| NEXT_PUBLIC_API_BASE_URL | 后端 API 地址 | https://api.example.com |

## 健康检查

### 后端健康检查

```bash
curl http://your-backend-domain:8080/ping
```

预期响应：
```json
{"message": "pong"}
```

### 前端健康检查

直接访问前端域名，应该能看到监控界面。

## 数据库初始化

首次部署时，应用会自动创建必要的数据库表。如果需要手动初始化：

```sql
CREATE DATABASE uptime CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

## 故障排查

### 常见问题

1. **后端无法连接数据库**
   - 检查数据库服务是否运行
   - 验证连接参数是否正确
   - 检查网络连通性

2. **前端无法连接后端**
   - 确认 CORS 配置正确
   - 检查 API 地址是否正确
   - 验证网络连通性

3. **监控检查失败**
   - 检查目标 URL 是否可访问
   - 验证网络出站规则
   - 检查 SSL 证书问题

### 日志查看

```bash
# 查看后端日志
docker-compose logs -f backend

# 查看所有服务日志
docker-compose logs -f
```

## 性能优化

1. **数据库优化**
   - 为查询字段添加索引
   - 定期清理历史数据
   - 配置连接池

2. **Redis 优化**
   - 设置合适的内存限制
   - 配置持久化策略
   - 监控内存使用

3. **前端优化**
   - 启用 CDN 加速
   - 配置缓存策略
   - 压缩静态资源

## 监控和告警

建议设置以下监控指标：

- 后端服务可用性
- 数据库连接状态
- Redis 连接状态
- API 响应时间
- 错误率统计

## 备份策略

1. **数据库备份**
   ```bash
   docker-compose exec mysql mysqldump -u uptime -p uptime > backup.sql
   ```

2. **配置备份**
   - 备份环境变量配置
   - 备份 Docker Compose 文件
   - 备份前端构建配置

## 更新部署

1. **后端更新**
   ```bash
   docker-compose build backend
   docker-compose up -d backend
   ```

2. **前端更新**
   ```bash
   cd frontend
   npm run build
   # 重新上传 out 目录到 EdgeOne Pages
   ```

## 安全建议

1. 使用强密码
2. 定期更新依赖
3. 配置防火墙规则
4. 启用 HTTPS
5. 定期备份数据
6. 监控异常访问