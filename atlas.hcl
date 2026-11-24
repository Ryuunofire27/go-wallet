data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "./atlas/loader",
  ]
}

env "gorm" {
  src = data.external_schema.gorm.url
  dev = "postgres://postgres:atlas@127.0.0.1:55432/atlas_dev?sslmode=disable"

  migration {
    dir = "file://migrations"
  }

  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}

env "dev" {
  url = "postgres://${getenv("POSTGRES_USER")}:${getenv("POSTGRES_PASSWORD")}@${getenv("DB_HOST")}:${getenv("DB_PORT")}/${getenv("POSTGRES_DB")}?sslmode=disable"

  src = data.external_schema.gorm.url

  dev = getenv("ATLAS_DEV_URL")

  migration {
    dir = "file://migrations"
  }
}