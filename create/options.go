package create

import "github.com/knyazev-ro/vulcan-migrate/types"

func (column *Column) Nullable() *Column {
	column.nullable = true
	return column
}

func (column *Column) CascadeOnDelete() *Column {
	column.cascadeOnDelete = true
	return column
}

func (column *Column) CascadeOnUpdate() *Column {
	column.cascadeOnUpdate = true
	return column
}

func (column *Column) NullOnDelete() *Column {
	column.nullOnDelete = true
	return column
}

func (column *Column) NullOnUpdate() *Column {
	column.nullOnUpdate = true
	return column
}

func (column *Column) RestrictOnDelete() *Column {
	column.restrictOnDelete = true
	return column
}

func (column *Column) RestrictOnUpdate() *Column {
	column.restrictOnUpdate = true
	return column
}

func (column *Column) NoActionOnDelete() *Column {
	column.noActionOnDelete = true
	return column
}

func (column *Column) NoActionOnUpdate() *Column {
	column.noActionOnUpdate = true
	return column
}

func (column *Column) DefaultOnDelete() *Column {
	column.defaultOnDelete = true
	return column
}

func (column *Column) DefaultOnUpdate() *Column {
	column.defaultOnUpdate = true
	return column
}

func (column *Column) Unique() *Column {
	column.unique = true
	return column
}

func (column *Column) Default(t *types.ConvertType) *Column {
	column.dDefault = t.Field
	return column
}
