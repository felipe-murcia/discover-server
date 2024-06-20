package main

import (
	"net/http"
	"time"

	"discoverco.co/server/configs"
	"discoverco.co/server/routes"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

func init() {
	configs.ConnectToDB()
}

func main() {

	r := gin.Default()

	// Apply the middleware to the router (works with groups too)
	r.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     false,
		ValidateHeaders: false,
	}))

	routes.PersonRouter(r)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world from server Go.",
		})
	})
	r.Run()
}

// func main() {
// 	// Información de conexión
// 	//dsn := "user=tu_usuario password=tu_contraseña dbname=tu_base_de_datos port=5432 sslmode=disable TimeZone=Asia/Shanghai"
// 	//dsn := "host=ep-holy-flower-a4ctw7pj-pooler.us-east-1.aws.neon.tech user=default password=jhdp04RJtUCn dbname=verceldb port=5432 sslmode=require"
// 	// Obtener variables de entorno

// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}

// 	host := os.Getenv("PGHOST")
// 	port := os.Getenv("PGPORT")
// 	user := os.Getenv("PGUSER")
// 	password := os.Getenv("PGPASSWORD")
// 	dbname := os.Getenv("PGDATABASE")
// 	sslmode := os.Getenv("PGSSLMODE")

// 	fmt.Println("Host de DB :", host)

// 	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
// 		host, port, user, password, dbname, sslmode)

// 	fmt.Println("Conexion postgresql:", dsn)

// 	// Abre la conexión a la base de datos
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		fmt.Println("Error al abrir la conexión:", err)
// 		return
// 	}

// 	// Verifica la conexión
// 	sqlDB, err := db.DB()
// 	if err != nil {
// 		fmt.Println("Error al obtener la instancia de la base de datos:", err)
// 		return
// 	}

// 	err = sqlDB.Ping()
// 	if err != nil {
// 		fmt.Println("No se pudo conectar a la base de datos:", err)
// 		return
// 	}

// 	fmt.Println("Conexión exitosa a la base de datos PostgreSQL")
// }
