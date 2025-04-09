package main

import (
	"log"

	"github.com/AnhducNA/go-ecommerce/global"
	"github.com/AnhducNA/go-ecommerce/internal/initialize"
	"gorm.io/gen"
)

func main() {
	// Initialize logger
	initialize.InitLogger()
	// Initialize DB connection
	initialize.InitMysql()

	// Verify DB connection
	if global.Mdb == nil {
		log.Fatal("Database connection not initialized")
	}

	// Create generator
	g := gen.NewGenerator(gen.Config{
		OutPath:      "./internal/models",
		Mode:         gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		ModelPkgPath: "github.com/AnhducNA/go-ecommerce/internal/models",
	})

	g.UseDB(global.Mdb)

	// Generate model for specific table
	g.GenerateModel("go_crm_user",
		gen.FieldType("created_at", "time.Time"),
		gen.FieldType("updated_at", "time.Time"),
	)

	// Execute and generate files
	g.Execute()
}
