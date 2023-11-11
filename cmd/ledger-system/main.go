package main

import (
	"github.com/gin-gonic/gin"
	config "github.com/israel-duff/ledger-system/pkg/config"
	"github.com/israel-duff/ledger-system/pkg/routes"

	"github.com/israel-duff/ledger-system/pkg/db"
)

func main() {
	databaseObject, err := config.NewDBConnection()

	if err != nil {
		panic("Could not connect to database!!!")
	}

	db.RunMigrationUp(databaseObject.GetDBConnection())

	r := gin.Default()

	// TODO: use globle Middlewhare here
	// r.Use(cors.New(cors.Config{
	// 	// AllowOrigins:     []string{"http://localhost", "http://127.0.0.1"},
	// 	AllowMethods:     []string{"PUT", "PATCH", "POST", "GET"},
	// 	AllowHeaders:     []string{"Origin", ""},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	AllowOriginFunc: func(origin string) bool {
	// 		return origin != "https://github.com"
	// 	},
	// 	// MaxAge: 12 * time.Hour,
	// }))

	r.Use(corsMiddleware())

	// Register Routes here
	routes.RegisterAuthRoutes(&r.RouterGroup)
	routes.RegisterCOARoutes(&r.RouterGroup)
	routes.RegisterWalletRoutes(&r.RouterGroup)

	r.Run(":5050")

}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
