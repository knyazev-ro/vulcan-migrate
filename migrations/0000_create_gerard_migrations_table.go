package migrations

import (
	"github.com/knyazev-ro/vulcan-migrate/create"
	"github.com/knyazev-ro/vulcan-migrate/facades"
	"github.com/knyazev-ro/vulcan-migrate/types"
)

var createGerardMigrationsTable_0000 = facades.NewMigration("0000_create_gerard_migrations_table", func(m *facades.Migration) {
	m.Up(
		func() any {
			return &create.Table{
				Name: "gerard_migrations",
				Body: []*create.Column{
					create.NewId(),
					create.NewString("name", 255),
					create.NewBigInteger("wave_id"),

					create.NewTimestamp("updated_at").Default(types.Now()),
					create.NewTimestamp("created_at").Default(types.Now()),
				},
			}
		},
	)

	m.Down(
		func() any {
			return &create.Table{
				Name: "gerard_migrations",
				Drop: true,
			}
		},
	)
})
