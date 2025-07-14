package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func setup() *fiber.App {
	app := fiber.New()
	// TODO:
	// cfg := config.NewConfig()
	// db := bootstrap.NewDB(&cfg)
	// ctx := ctx.ProvideContext()

	app.Use(logger.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000", // Allow requests from the frontend
		AllowMethods: "GET,POST,PUT,DELETE",   // Allow specific HTTP methods
		AllowHeaders: "Content-Type,Authorization",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("decauth is running 🚀")
	})

	// TODO: auth
	// oauthRepoFactory := repository.NewOAuthRepositoryFactory()
	// oauthRepoFactory.Register("google", oauth.NewGoogleOAuthRepository(&cfg))
	//
	// userRepo := gorm.NewUserGormRepository(db)
	// userUsecase := usecase.NewUserUsecase(userRepo)
	// _userHandler := rest.NewUserHandler(userUsecase)
	//
	// authUsecase := usecase.NewAuthUsecase(userRepo, &oauthRepoFactory, &cfg, minioRepo, ctx)
	// authHandler := rest.NewAuthHandler(authUsecase)
	//
	//
	// authRoute := app.Group("/auth")
	// authRoute.Post("/login", authHandler.LoginWithCredentials)
	// authRoute.Post("/login/oauth", authHandler.LoginWithOAuth)
	// authRoute.Post("/register", authHandler.Register)

	return app
}

func main() {
	app := setup()
	app.Listen(":9090")
}
