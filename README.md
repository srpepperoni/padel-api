Run bdd (postgresql) using docker-compose

```yaml
docker-compose -f deploy/docker-compose.yml up
```

Use adminer (or console as you please) and then create api_padel scheme

After scheme creation, you can run the application
```bash
go run cmd/main/main.go
```

Enpoint: http://localhost:8000
Swagger: http://localhost8000/swagger/index.html


TODO

Actualizar set de resultado
Update match setea los valores de pareja en el body deberia coger solo referencia de matchid