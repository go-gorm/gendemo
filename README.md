# gendemo
The best practices of using [GEN](https://github.com/go-gorm/gen).

### run demo
1. start a test database server instance

```shell
cd cmd & docker-compose up -d
```
2. prepare test table
```shell
go run main.go 
```
3. generate code
```shell
sh generate.sh
```
