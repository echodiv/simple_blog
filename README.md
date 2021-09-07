
### setup postgres
```bash
$ sudo docker-compose exec db sh
# createdb -U dev_user dev_user
```
### Using golang migrate
```bash
migrate -path migrations -database "postgres://localhost/dev_db?sslmode=disable&user=dev_user&password=secret12345" up 1
```