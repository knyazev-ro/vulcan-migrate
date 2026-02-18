package facades

import "github.com/knyazev-ro/vulcan-migrate/registry"

type Migration struct {
	MigrationName string
}

func NewMigration(name string, f func(*Migration)) *Migration {
	m := &Migration{
		MigrationName: name,
	}
	f(m)
	return m

}

func (m *Migration) Up(f func() any) {
	registry.Register(registry.Action.Up, m.MigrationName, f)
}

func (m *Migration) Down(f func() any) {
	registry.Register(registry.Action.Down, m.MigrationName, f)
}
