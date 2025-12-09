<div align="center">

# 🍀 GoFrame 实验三 Web Demo

[![Go](https://img.shields.io/badge/Go-1.20+-00ADD8?logo=go)](https://go.dev/)
[![Framework](https://img.shields.io/badge/GoFrame-v2-brightgreen)](https://goframe.org)
[![DB](https://img.shields.io/badge/DB-MySQL-blue?logo=mysql)](https://www.mysql.com/)
[![Auth](https://img.shields.io/badge/Auth-JWT-orange)](https://jwt.io/)

</div>

---

> 📧 邮箱验证码注册/登录 & 🔑 JWT 鉴权 & 📦 物品 RESTful CRUD，附带前端演示页与接口调试示例。

---

## 📚 目录
- [🍀 GoFrame 实验三 Web Demo](#-goframe-实验三-web-demo)
  - [📚 目录](#-目录)
  - [🧭 业务概览](#-业务概览)
  - [🧰 技术栈](#-技术栈)
  - [⚙️ 环境准备](#️-环境准备)
  - [📂 目录速览](#-目录速览)
  - [🚀 运行](#-运行)
  - [🛠️ 核心接口](#️-核心接口)
  - [🖥️ 前端体验](#️-前端体验)
  - [🧪 调试与测试](#-调试与测试)
  - [⚠️ 注意事项](#️-注意事项)

## 🧭 业务概览
- 邮箱验证码注册/登录、密码登录、密码重置；验证码记录在 `email_code` 表并校验后标记已用。
- 登录颁发 JWT（`Authorization: Bearer <token>`），中间件兼容 Session。
- 物品资源 `item` 表提供标准 RESTful CRUD，用于受保护业务示例。
- 保留 customer / merchant / order 示例模块，可继续扩展。
- 前端 `web/index.html`：发送验证码、注册/登录、重置密码、物品 CRUD，自动缓存 token。

## 🧰 技术栈
- GoFrame v2、MySQL
- JWT：github.com/golang-jwt/jwt/v5
- 邮件发送：github.com/jordan-wright/email

## ⚙️ 环境准备
1. Go 1.20+，`go mod tidy` 已同步依赖。
2. MySQL：导入 `manifest/sql/create.sql`（含 user、email_code、item、customer、merchant、order 等表）。
3. 修改 `manifest/config/config.yaml`：数据库、SMTP、JWT。
   ```yaml
   email:
     from: "Demo <noreply@example.com>"
     smtp:
       host: "smtp.example.com"
       port: 587
       username: "noreply@example.com"
       password: "your_smtp_password"
     codeExpireMinutes: 10
   jwt:
     secret: "change-me-to-a-long-random-secret"
     expireHours: 24
   ```

## 📂 目录速览
- `api/` 接口请求/响应定义（user、item 等）
- `internal/controller/` 路由实现（用户、物品、商户、订单）
- `internal/logic/` 业务逻辑（用户、JWT、会话、中间件）
- `internal/dao/`、`internal/model/` 数据访问、实体/DO
- `manifest/sql/` 建表脚本；`manifest/config/` 配置模板
- `web/` 前端演示；`docs/api.http` 接口调试示例

## 🚀 运行
```bash
go run main.go
# http://localhost:8000
# 前端演示页：http://localhost:8000/index.html
```

## 🛠️ 核心接口
- 邮箱验证码：`POST /user/email/send-code`（purpose: register|login|reset）
- 邮箱注册：`POST /user/email/sign-up` → JWT
- 邮箱登录：`POST /user/email/sign-in` → JWT
- 重置密码：`POST /user/email/reset-password` → JWT
- 密码登录：`POST /user/sign-in`（passport 可用邮箱）→ JWT
- 个人资料：`GET /user/profile`（Bearer）
- 物品：`GET /items` | `POST /items` | `GET /items/{id}` | `PUT /items/{id}` | `DELETE /items/{id}`

## 🖥️ 前端体验
- 打开 `http://localhost:8000/index.html`
- 功能：发送验证码、邮箱注册/登录、密码登录、重置密码、物品 CRUD，响应在右侧输出并自动保存 token。

## 🧪 调试与测试
- 使用 `docs/api.http`（VSCode REST Client / Apipost / Postman）快速调试。
- 验证路径：发送验证码 → 注册获取 JWT → 携带 Bearer 调 `GET /user/profile` → 体验 `/items` CRUD。

## ⚠️ 注意事项
- SMTP 必须可用，否则验证码无法发送；验证码默认 10 分钟失效。
- 已开启 CORS；受保护接口需 Bearer token 或已有 Session。
