# The Go Ledger System
This ledger system is a prototype implementation of the real world accounting ledger system. This project was inspired after taking on similar project in my current day-job.
The aim of this project is to both learn about how the real world ledger accounting system works while also learning and improving on my go-lang skills.

# Database Tools
## The folowing go database tools needs to be installed
1. Gorm (for ORM)
2. Gen tool (from Gorm for generating models and data access objects from sql)
3. Goose (for managing migrations)


## How it works
1. ### Use the Goose tool to create a migration
The goose tool is used to generate versioned migration written in either go or sql. Here is a sample command to generate sql migration file named init:

```sh
~/go/bin/goose -dir pkg/db/migrations postgres "host=localhost user=postgres password=password dbname=accounting_ledger port=5432 sslmode=disable TimeZone=Asia/Shanghai" create init sql
```

When you run the above command, goose will generate a migration file that may look like this:
```sql
-- +goose Up
-- +goose StatementBegin
SELECT * WHERE 1 = 1
-- +goose StatementEnd
```

2. ### Edit the migration file with your desired SQL (or go) command.
notice that goose generated an empty sql file. This means that you have to manually write the sql schema yourself. You can edit the sql file generated and put in your sql statements.
Here is a sample SQL command to create a user table:

```sql
-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users(
    id UUID DEFAULT uuid_generate_v4 (),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NULL DEFAULT NULL,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    email_address VARCHAR(255) NOT NULL UNIQUE,
    phone_number VARCHAR(20) NOT NULL UNIQUE,
    PRIMARY KEY (id)
);
-- +goose StatementEnd
```


3. ### Run the migration 
After generating and modifying the migration file(s), the next step is to execute (or run) the migration file(s). 
This step is to ensure that the sql schema is effected in the database. Here is a sample command for running migration:

```sh
~/go/bin/goose -dir pkg/db/migrations postgres "host=localhost user=postgres password=password dbname=accounting_ledger port=5432 sslmode=disable TimeZone=Asia/Shanghai" up
```

### Check Migration Status (optional)
You can optionally check migration status:
```sh
~/go/bin/goose postgres "host=localhost user=postgres password=password dbname=accounting_ledger port=5432 sslmode=disable TimeZone=Asia/Shanghai" status
```

4. ### Generate models
The next step is to generate models and data access objects from the database using the Gorm gen tool. This step will automatically generate these models and data objects from the migrated database schema. The following is a sample command to generate the models:
```sh
~/go/bin/gentool -dsn "postgresql://postgres:postgres@localhost:5432/accounting_ledger?connect_timeout=10&sslmode=disable" -db postgres -outPath "pkg/db/dao"
```

