Para realização de teste de cobertura, é necessário

- go test -cover [file] <- Roda teste de cobertura
- go test -coverprofile c.out <- Salva resultado em um arquivo
- go tool cover -html=c.out <- Exibe um resultado visual em HTML (Para windows, use strings no stributo HTML)