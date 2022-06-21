Run bdd (postgresql) using docker-compose

```yaml
docker-compose -f postgre.yml up
```

use adminer (or console as you please) and then create sysdig_padel scheme

after that run from root:
```bash
go run cmd/main/main.go
```


## TODO

- [x] Form for adding new players
- [x] New model match
- [ ] List Matches
- [ ] Confirmation message (CREATED)
- [ ] Delete User page
- [ ] Delete match page
- [ ] Set results
- [ ] Tournament case

