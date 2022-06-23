Run bdd (postgresql) using docker-compose

```yaml
docker-compose -f docker-compopse.yml up
```

Use adminer (or console as you please) and then create sysdig_padel scheme

After scheme creation, you can run the application
```bash
go run cmd/main/main.go
```

Enpoint: http://localhost:8000
Swagger: http://localhost8000/swagger/index.html