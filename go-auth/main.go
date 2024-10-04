package main

// 使用的框架:fiber(web框架) gorm(数据库) air(热重载) driver(mysql) jwt(令牌)
// 自己那个数据库表的password可能要换longtext类型
import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/zccccc01/go-auth/database"
	"github.com/zccccc01/go-auth/routes"
)

func main() {
	database.Connect()

	app := fiber.New()

	// 设置 CORS 中间件 允许前端3000端口
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",                       // 允许来自前端的请求
		AllowMethods:     "GET,POST,PUT,DELETE",                         // 允许的 HTTP 方法
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization", // 允许的头部
		AllowCredentials: true,                                          // 允许发送凭证
	}))

	routes.Setup(app)

	app.Listen(":8000")
}
