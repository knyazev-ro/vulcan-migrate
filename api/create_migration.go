package api

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/knyazev-ro/perturabo/migrate"
	"github.com/knyazev-ro/perturabo/utils"
)

func CreateMigration(args []string) {

	settings := utils.LoadSettings()
	migrationsFolder := settings.Migrations
	templateFile := settings.TemplateCreate

	existingFIles, err := migrate.LoadMigrationFiles()
	if err != nil {
		utils.ErrorPrintln("Error during reading migration file. Needs at least 1 migration")
		return
	}

	var orderNumberOfLastOne int64
	var newOrderNumber int64
	if len(existingFIles) <= 0 {
		newOrderNumber = 0
	} else {
		lastOne := strings.Split(existingFIles[len(existingFIles)-1], "_")[0]
		orderNumberOfLastOne, err = strconv.ParseInt(lastOne, 10, 64)
		if err != nil {
			utils.ErrorPrintln(err.Error())
		}
		newOrderNumber = orderNumberOfLastOne + 1
	}

	formatNumber := fmt.Sprintf("%04d", newOrderNumber)

	name := formatNumber + "_create_" + args[0]

	tableName := strings.ToLower(args[1])
	MigrationNameCamelCase, MigrationNameVar, MigrationName, err := utils.Normalize(name)
	if err != nil {
		utils.ErrorPrintln("Error! ", err.Error())
		return
	}

	data := map[string]string{
		"MigrationNameCamelCase": MigrationNameCamelCase,
		"MigrationName":          MigrationName,
		"MigrationNameVar":       MigrationNameVar,
		"TableName":              tableName,
	}

	os.MkdirAll(migrationsFolder, 0755)
	outFile := fmt.Sprintf("%s/%s.go", migrationsFolder, strings.ToLower(name))
	path, err := utils.ParseTemplate(templateFile, outFile, data, args)
	if err != nil {
		utils.ErrorPrintln("Error creating "+path+":", err.Error())
		return
	}

	utils.SuccessPrintln("Created migration on CREATE: ", outFile)
}
