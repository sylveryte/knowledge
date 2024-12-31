# Postgres for database SQL

## [[Start a local docker quick]]

## Docker compose

{/ ./docker-compose.yml}

```sh
sudo docker compose  up db
```

## Connect

connection string

```sh
pgcli "postgresql://localhost:5000?user=sylveryte&password=study"
```

## Text vs varchar

- use text [stackoverflow](https://stackoverflow.com/questions/4848964/difference-between-text-and-varchar-character-varying)
- also use varchar (should limit it)
  -- name should never exceed 100char

```markdown
To sum it all up:
char(n) – takes too much space when dealing with values shorter than n (pads them to n), and can lead to subtle errors because of adding trailing spaces, plus it is problematic to change the limit
varchar(n) – it's problematic to change the limit in live environment (requires exclusive lock while altering table)
varchar – just like text
text – for me a winner – over (n) data types because it lacks their problems, and over varchar – because it has distinct name
```

## Triggers

{https://stackoverflow.com/questions/2679333/database-design-for-invoices-invoice-lines-revisions}[For invoice stop changing]

## Cool domains phone and email

{https://github.com/rin-nas/postgresql-patterns-library/blob/master/domains/email.sql}
