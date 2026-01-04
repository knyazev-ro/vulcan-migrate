package migrations

import (
	"perturabo/create"
	"perturabo/facades"
	"perturabo/types"
)

var createUserTable_0002 = facades.NewMigration("0002_create_user_table", func(m *facades.Migration) { 
	// Name of migration MUST be the same as the filename without .go extenstion
	m.Up(
		func() any {
			return &create.Table{
				Name: "users",
				Body: []*create.Column{
					create.NewId(),

					create.NewTimestamp("updated_at").Default(types.Now()),
					create.NewTimestamp("created_at").Default(types.Now()),
				},
			}
		},
	)

	m.Down(
		func() any {
			return &create.Table{
				Name: "users",
				Drop: true,
			}
		},
	)
})
