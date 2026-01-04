package api

import (
	"fmt"
	"os"
	"perturabo/utils"
)

func Init() {

	settings := utils.LoadSettings()
	migrationsFolder := settings.Migrations
	templateFile := settings.TemplateGerardMigrationsTable
	data := map[string]string{}
	os.MkdirAll(migrationsFolder, 0755)
	outFile := fmt.Sprintf("%s/%s.go", migrationsFolder, "0000_create_gerard_migrations_table")
	if _, err := os.Stat(outFile); !os.IsNotExist(err) {
		return
	}
	path, err := utils.ParseTemplate(templateFile, outFile, data, []string{})
	if err != nil {
		utils.ErrorPrintln("Error creating "+path+":", err.Error())
		return
	}
}
