package alter

import (
	"fmt"
	"strings"

	"github.com/knyazev-ro/vulcan-migrate/create"
	"github.com/knyazev-ro/vulcan-migrate/types"
)

func (c *Column) Type(dbtype *types.DatabaseType) *Column {
	c.typeAlter = fmt.Sprintf("%s TYPE %s", c.field, dbtype.Field)
	return c
}

func (c *Column) Default(t *types.ConvertType) *Column {
	c.defaultSet = fmt.Sprintf("%s SET DEFAULT %s", c.field, t.Field)
	return c
}

func (c *Column) DropDefault() *Column {
	c.defaultDrop = fmt.Sprintf("%s DROP DEFAULT", c.field)
	return c
}

func (c *Column) NotNull() *Column {
	c.notNullSet = fmt.Sprintf("%s SET NOT NULL", c.field)
	return c
}

func (c *Column) DropNotNull() *Column {
	c.notNullDrop = fmt.Sprintf("%s DROP NOT NULL", c.field)
	return c
}

func (c *Column) Rename(newName string) *Column {
	c.columnRename = fmt.Sprintf("RENAME COLUMN %s TO %s", c.field, newName)
	return c
}

func (c *Column) RenameTypeUsing(dbtype *types.DatabaseType) *Column {
	c.columnRenameUsing = fmt.Sprintf("%s TYPE %s USING %s::%s", c.field, strings.ToUpper(dbtype.Field), c.field, dbtype.Field)
	return c
}

func (c *Column) Add(newcol *create.Column) *Column {
	c.columnAdd = fmt.Sprintf("ADD COLUMN %s", newcol.ToSQL())
	return c
}

func (c *Column) Drop() *Column {
	c.columnDrop = fmt.Sprintf("DROP COLUMN %s", c.field)
	return c
}

func (c *Column) Statistics(value int) *Column {
	c.statistics = fmt.Sprintf("%s SET STATISTICS %d", c.field, value)
	return c
}
