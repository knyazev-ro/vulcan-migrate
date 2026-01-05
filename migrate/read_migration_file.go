package migrate

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/knyazev-ro/perturabo/alter"
	"github.com/knyazev-ro/perturabo/create"
	"github.com/knyazev-ro/perturabo/registry"
	"github.com/knyazev-ro/perturabo/utils"
)

func LoadMigrationFiles() ([]string, error) {
	settings := utils.LoadSettings()
	migrationArr, err := os.ReadDir(settings.Migrations)
	if err != nil {
		for _, readEntry := range migrationArr {
			fmt.Println(readEntry.Name())
		}
		fmt.Println("Error during reading migrations folder. ", err.Error())
		return []string{}, err
	}

	files := []string{}
	for _, readEntry := range migrationArr {
		filename := readEntry.Name()
		norm, err := ValidateFileName(filename)
		if err != nil {

			if norm == "head" {
				continue
			}

			fmt.Println("Not valid. Skip ", filename)
			continue
		}
		files = append(files, norm)
	}
	sort.Strings(files)

	return files, nil
}

func ValidateFileName(filename string) (string, error) {
	split := strings.Split(filename, ".")
	ext := split[len(split)-1]
	if ext != "go" {
		return "", os.ErrInvalid
	}
	splitName := strings.Split(split[0], "_")

	isCreate := utils.Contains(splitName, func(s string) bool {
		return s == "create"
	}) >= 0

	isAlter := utils.Contains(splitName, func(s string) bool {
		return s == "alter"
	}) >= 0

	isAlterAndCreate := isCreate && isAlter

	if (isCreate || isAlter) && !isAlterAndCreate {
		return split[0], nil
	}
	return splitName[0], os.ErrInvalid
}

func Get(action string) ([]string, error) {
	loadMigrationFormattedFiles, err := LoadMigrationFiles()
	if err != nil {
		fmt.Println("Error", err.Error())
	}
	migrationsList := []string{}
	for _, readEntry := range loadMigrationFormattedFiles {

		isCreate := strings.Split(readEntry, "_")[1] == "create"

		if action == registry.Action.Up {
			if f, ok := registry.Up[readEntry]; ok {
				if isCreate {
					table := f().(*create.Table) // приведение типа
					migrationsList = append(migrationsList, GenerateCreateTableSQL(table))
				} else {
					table := f().(*alter.Table) // приведение типа
					migrationsList = append(migrationsList, GenerateAlterTableSQL(table))
				}
			}
		} else {
			if f, ok := registry.Down[readEntry]; ok {
				if isCreate {
					table := f().(*create.Table) // приведение типа
					migrationsList = append(migrationsList, GenerateCreateTableSQL(table))
				} else {
					table := f().(*alter.Table) // приведение типа
					migrationsList = append(migrationsList, GenerateAlterTableSQL(table))
				}
			}
		}

	}

	return migrationsList, nil
}
