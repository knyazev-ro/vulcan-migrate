# Perturabo

**Perturabo** — это минималистичный инструмент для управления миграциями баз данных на Go.
Он позволяет создавать, изменять, применять и откатывать миграции в декларативном виде прямо в Go-коде.

## Возможности

* Создание миграций с помощью Go-структур
* Поддержка `CREATE` и `ALTER` миграций
* Автоматическая генерация SQL
* Разделение на `Up` и `Down` действия
* Rollback последних изменений
* Валидация имён и структуры файлов миграций
* CI/CD готовность
* Возможность писать миграции без ORM (использует `Masterminds/squirrel` для SQL-конструкций)

## Установка

```bash
go get github.com/knyazev-ro/perturabo
```

## Структура проекта

```
perturabo/
  commands/      # CLI-команды
  migrate/       # Логика применения миграций
  migrations/    # миграции (Go-файлы)
  registry/      # Регистрация миграций
  create/        # Описание CREATE миграций
  alter/         # Описание ALTER миграций
  utils/         # Утилиты
```

## Использование

### Создание миграции

Создание миграции для создания таблицы:

```bash
perturabo create:migration user_table users
```

Создание миграции для изменения таблицы:

```bash
perturabo alter:migration user_table users
```

### Применение миграций

```bash
perturabo migrate:run
```

### Откат миграций

```bash
perturabo migrate:rollback
```

## Пример миграции

```go
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
```

## Валидация миграций

Perturabo проверяет:

* имя файла (должно содержать `create` или `alter`)
* корректное расширение `.go`
* уникальность имён миграций
