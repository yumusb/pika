# SQLite 版本部署指南

## 环境要求

- Docker 20.10+
- Docker Compose 1.29+

## 快速开始

### 1. 下载配置文件

```bash
# 下载 docker-compose.yml 配置文件
curl -O https://raw.githubusercontent.com/dushixiang/pika/main/docker-compose.sqlite.yml
# 下载配置文件示例
curl -o config.yaml https://raw.githubusercontent.com/dushixiang/pika/main/config.sqlite.yaml

# 或使用 wget
wget https://raw.githubusercontent.com/dushixiang/pika/main/docker-compose.sqlite.yml
wget -O config.yaml https://raw.githubusercontent.com/dushixiang/pika/main/config.sqlite.yaml
```

### 2. 修改配置

数据库配置（SQLite）：

```yaml
database:
  enabled: true
  type: sqlite
  sqlite:
    path: "./data/pika.db"
```

其他配置项请参考 [通用配置说明](common-config.md)。

### 3. 启动服务

```bash
# 启动所有服务
docker-compose -f docker-compose.sqlite.yml up -d

# 查看服务状态
docker-compose -f docker-compose.sqlite.yml ps

# 查看日志
docker-compose -f docker-compose.sqlite.yml logs -f pika
```

### 4. 访问服务

服务启动后，访问 http://localhost:8080

默认账户：
- 用户名: `admin`
- 密码: `admin123`

### 5. 停止服务

```bash
# 停止服务
docker-compose -f docker-compose.sqlite.yml stop

# 停止并删除容器
docker-compose -f docker-compose.sqlite.yml down

# 停止并删除容器及数据卷
docker-compose -f docker-compose.sqlite.yml down -v
```

## 数据持久化

系统数据分别存储在以下目录：
- **SQLite 数据库**：`./data/pika.db` - 存储配置和审计数据
- **VictoriaMetrics 数据**：`./data/vmdata` - 存储时序指标数据（默认保留 7 天）

### 备份 SQLite 数据库

```bash
# 备份数据库
cp ./data/pika.db ./data/pika.db.backup-$(date +%Y%m%d)

# 或使用 SQLite 的备份命令
sqlite3 ./data/pika.db ".backup './data/pika.db.backup-$(date +%Y%m%d)'"
```

### 备份 VictoriaMetrics 数据

```bash
# 备份 VictoriaMetrics 数据
tar -czf vmdata-backup-$(date +%Y%m%d).tar.gz ./data/vmdata

# 恢复时停止服务，解压到 ./data/vmdata 目录即可
```

## 故障排查

### 数据库文件权限问题

- 确认 `./data` 目录存在且有写入权限
- 检查 SQLite 数据库文件权限：`ls -l ./data/pika.db`
- 如果遇到权限问题，可以尝试：
  ```bash
  chmod 755 ./data
  chmod 644 ./data/pika.db
  ```

其他故障排查请参考 [通用配置说明](common-config.md)。
