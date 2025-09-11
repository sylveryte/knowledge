# Farm database Postgresql

#db

- sqls config {/ ~/.config/sqls/config.yml}

```yaml
# Set to true to use lowercase keywords instead of uppercase.
lowercaseKeywords: false
connections:
  - alias: farmstring
    driver: postgresql
    dataSourceName: postgresql://localhost:5000/farm?user=sylveryte&password=study
```

- Note:
  -- use snake_case for column and constraint names

## Create table

```sql
-- uuid id gen
CREATE TABLE chickens (
  id uuid NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
  name varchar(40) NOT NULL,
  color varchar(30) DEFAULT 'white',
  age int
);

-- AUTO_INCREMENT int
CREATE TABLE coop (
  id bigserial PRIMARY KEY,
  name varchar(40) NOT NULL,
  color varchar(30) DEFAULT 'brown'
);
```

## Alter table

```sql
ALTER TABLE chickens
ADD COLUMN coopId BIGINT
CONSTRAINT chickenCoopFkC REFERENCES coop(id);
```

## Insert values

```sql
INSERT INTO chickens(name,age,color) VALUES ( 'Cili', 24,'pink');
SELECT * FROM chickens;
```

## Update

```sql
UPDATE chickens SET coopid=2 WHERE color='pink';
```

## Delete

```sql
DELETE FROM chickens WHERE name='Cili';
```

## Join

```sql
SELECT
chickens.name,
coop.name
FROM
chickens
INNER JOIN coop
ON chickens.coopid = coop.id;
```
