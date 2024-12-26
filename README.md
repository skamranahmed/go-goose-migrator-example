## Go Goose Migrator Example

This project is a simple example of how to use [Goose](https://github.com/pressly/goose) for database migrations in Go. It demonstrates the creation and management of database migrations using a standalone binary.

### Prerequisites

- Go 1.23.3 or later
- Docker (for PostgreSQL setup)

### Getting Started

#### Clone the Repository
```bash
git clone https://github.com/skamranahmed/go-goose-migrator-example.git
cd go-goose-migrator-example
```


#### Setup PostgreSQL
To set up a PostgreSQL database using Docker, run:
```bash
make setup-postgres
```

This command will start a PostgreSQL container with the default credentials.

#### Build the Migrator
To build the migrator binary, run:

```bash
make build-migrator
```

#### Create a Migration
To create a new migration, use the following command:

```bash
make migrate-create name=<migration_name>
```

Replace `<migration_name>` with the desired name for your migration.

#### Run Migrations
To apply all pending migrations, run:

```bash
make migrate-up
```


#### Rollback Migrations
To rollback the last applied migration, use:

```bash
make migrate-down
```

### Environment Variables
The following environment variables are required for the application to connect to the PostgreSQL database:

- `DATABASE_USER`: The database user (default: `root`)
- `DATABASE_PASSWORD`: The database password (default: `password`)
- `DATABASE_NAME`: The name of the database (default: `postgres`)
- `DATABASE_HOST`: The database host (default: `localhost`)
- `DATABASE_PORT`: The database port (default: `5432`)

You can set these variables in your shell or modify the `Makefile` to change the defaults.

### Migration Files
Migration files are located in the `migrations` directory. Each migration file should follow the naming convention `YYYYMMDDHHMMSS_<migration_name>.go`.

### 📝 License
This project is licensed under the [MIT License](https://choosealicense.com/licenses/mit/)