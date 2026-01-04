package migrate

import (
	"fmt"

	"github.com/knyazev-ro/perturabo/alter"
	"github.com/knyazev-ro/perturabo/create"
)

func GenerateCreateTableSQL(table *create.Table) string {
	columns := table.Body
	body := ""
	name := table.Name
	drop := table.Drop

	if drop {
		return DropTableIfExists(name)
	}

	for i, col := range columns {
		body += "    " + col.ToSQL()
		if i < len(columns)-1 {
			body += ",\n"
		} else {
			body += "\n"
		}
	}
	return fmt.Sprintf("CREATE TABLE %s (\n%s);", name, body)
}

func GenerateAlterTableSQL(table *alter.Table) string {
	name := table.Name
	body := ""
	columns := table.Body
	for _, col := range columns {
		body += fmt.Sprintf("ALTER TABLE %s %s;\n", name, col.ToSQL())
	}
	return body
}

func DropTableIfExists(name string) string {
	return fmt.Sprintf("DROP TABLE IF EXISTS %s;", name)
}
