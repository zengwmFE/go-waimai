# Sky-Take-Out Go (苍穹外卖 Go版)

Go backend for the 苍穹外卖 food delivery system, replacing the original Java Spring Boot server.

## Tech Stack

- **Framework**: Gin
- **ORM**: GORM
- **Auth**: JWT (dual admin/user token)
- **Config**: Viper (YAML)
- **Logging**: Zap
- **Database**: MySQL
- **Cache**: Redis

## Project Structure

```
server/
├── main.go              # Entry point
├── config.yaml           # Configuration
├── config/               # Config loading
├── controller/           # HTTP handlers
│   ├── admin/            # Admin-side controllers
│   └── user/             # User/mini-program controllers
├── service/              # Business logic
├── dao/                  # Data access layer (GORM)
├── model/
│   ├── entity/           # DB table models
│   ├── dto/              # Request DTOs
│   └── vo/               # Response VOs
├── middleware/            # Auth, CORS, logging
├── router/               # Route definitions
├── utils/                # JWT, response helpers
└── websocket/            # Real-time order alerts
admin/                    # Vue.js admin dashboard
miniprogram/              # WeChat mini-program
database/                 # SQL scripts
```

## Quick Start

1. Start MySQL and create the database:
```bash
mysql -u root -p < database/init.sql
```

2. Edit `config.yaml` with your database settings.

3. Run the server:
```bash
cd server
go mod tidy
go run main.go
```

4. Server starts at `http://localhost:8080`

## API Documentation

API routes mirror the original project's endpoints:

- **Admin**: `/admin/employee/*`, `/admin/category/*`, `/admin/dish/*`, `/admin/setmeal/*`, `/admin/order/*`, `/admin/shop/*`, `/admin/common/*`, `/admin/report/*`, `/admin/workspace/*`
- **User**: `/user/user/*`, `/user/category/*`, `/user/dish/*`, `/user/setmeal/*`, `/user/shop/*`, `/user/shoppingCart/*`, `/user/order/*`, `/user/addressBook/*`
- **WebSocket**: `/ws/:sid`

## Authentication

- Admin requests: JWT token in `token` header
- User requests: JWT token in `authentication` header
