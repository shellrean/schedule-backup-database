# Shellrean CronJob Backup Database

Shellrean CronJob backup database is a tool to backup database.

## Configuration
Create and change configuration on config.json.
```
{
  "port": "3306",
  "host": "localhost",
  "user": "root",
  "password": "",
  "db_name": "",
  "path": "./backup",
  "every": "@hourly"
}
```
For every property.\
|Value | Description|
|------|------------|
|@hourly| Every hour|
|daily | Every day midnight|
|@every 0h0m1s|Custom <hour>h<minutes>m<second>s|
|------|---------|
## Installation

Make sure go was installed on your computer.

```bash
go run app/main.go
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)