# Perturabo

**Perturabo** — минималистичный инструмент для управления миграциями баз данных на Go.
Позволяет создавать, изменять, применять и откатывать миграции в декларативном виде через Go-код.

---

## Возможности

* Создание миграций с помощью Go-структур (`CREATE` и `ALTER`)
* Поддержка `Up` и `Down` действий
* Автоматическая генерация SQL
* Rollback последних изменений
* Валидация имён и структуры файлов миграций
* CI/CD готовность
* Возможность писать миграции без ORM (использует `Gerard Vulcan Query Builder` для SQL-конструкций)
* CLI для управления миграциями

---

## Установка

```bash
go get github.com/knyazev-ro/perturabo
```

---

## Структура проекта

```
perturabo/
  cmd/           # CLI-команды и обработка аргументов
  migrate/       # Логика применения миграций
  migrations/    # Миграции (Go-файлы)
  registry/      # Регистрация миграций
  create/        # Описание CREATE миграций
  alter/         # Описание ALTER миграций
  utils/         # Вспомогательные функции
  facades/       # Фасад для декларативного описания миграций
```

---

## Использование

### CLI команды

| Команда               | Описание                               |
| --------------------- | -------------------------------------- |
| `pertdb:create-table` | Создаёт миграцию для новой таблицы     |
| `pertdb:alter-table`  | Создаёт миграцию для изменения таблицы |
| `pertdb:run`          | Применяет миграции                     |
| `pertdb:rollback`     | Откатывает последние миграции          |
| `pertdb:help`         | Выводит справку по командам            |

Пример запуска CLI:

```bash
go run main.go pertdb:create-table users_table users
go run main.go pertdb:alter-table users_table users
go run main.go pertdb:run
go run main.go pertdb:rollback
```

---

### Минимальный пример проекта с Perturabo

```go
package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	_ "myproject/migrations" // подключаем миграции, предварительно создайте в корне проекта папку migrations с файлом migrations.go (он должен быть пустой)

	"github.com/gorilla/mux"
	"github.com/knyazev-ro/perturabo/cmd"
	"myproject/routes"
)

func RegisterCli() {
	cmd.Handle(os.Args)
}

func App() {
	args := os.Args
	if args[1] == "server:up" {
		Server(args[1:])
		return
	}
	RegisterCli()
}

func Server(args []string) {
	r := mux.NewRouter()
	routes.RegisterRoutes(r)

	host := "0.0.0.0"
	port := "8080"
	hp := strings.Join([]string{host, port}, ":")

	fmt.Println(fmt.Sprintf("Gerard server Up at %s. Ready to accept connections!", hp))
	http.ListenAndServe(hp, r)
}

func main() {
	App()
}
```

---

### Пример миграции

```go
var createGerardMigrationsTable_0000 = facades.NewMigration(
	"0000_create_gerard_migrations_table",
	func(m *facades.Migration) {
		m.Up(func() any {
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
		})

		m.Down(func() any {
			return &create.Table{
				Name: "gerard_migrations",
				Drop: true,
			}
		})
	},
)
```

* `Up` описывает создание/изменение структуры
* `Down` описывает откат
* Миграция — чистый объект, её можно зарегистрировать и выполнить позже через runner

---

## Валидация миграций

Perturabo проверяет:

* Имя файла — должно содержать `create` или `alter`
* Расширение файла — `.go`
* Уникальность имён миграций