package main

import (
	"fmt"
	"github.com/go-jet/jet/v2/generator/metadata"
	"github.com/go-jet/jet/v2/generator/postgres"
	"github.com/go-jet/jet/v2/generator/template"
	postgres2 "github.com/go-jet/jet/v2/postgres"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dsn := os.Getenv("DB_CONNECTION_STRING")
	log.Println("DSN", dsn)
	err = postgres.GenerateDSN(dsn, "public", "./.gen", template.Default(postgres2.Dialect).
		UseSchema(func(schemaMetaData metadata.Schema) template.Schema {
			return template.DefaultSchema(schemaMetaData).
				UseModel(template.DefaultModel().
					UseTable(func(table metadata.Table) template.TableModel {
						return template.DefaultTableModel(table).
							UseField(func(columnMetaData metadata.Column) template.TableModelField {
								defaultTableModelField := template.DefaultTableModelField(columnMetaData)
								return defaultTableModelField.UseTags(
									fmt.Sprintf(`json:"%s"`, toSnakeCase(columnMetaData.Name)),
									fmt.Sprintf(`xml:"%s"`, columnMetaData.Name),
								)
							})
					}),
				)
		}))

	if err != nil {
		log.Fatal("Sheesh!", err.Error())
	}

	log.Println("Done generating...")
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
