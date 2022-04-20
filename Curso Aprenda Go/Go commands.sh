go version  - Mostra versão do go
go env      - Mostra as variáveis de ambiente
go help     [nome] - Mostra opções de ajuda
go run      (*.go // Executa todos os arquivos go)
            (nome do arquivo // Executa arquivo específico)
go build    - Gera um executável (bin)
            Informa erros caso existam
            Para pacotes, gera um arquivo e descarta o executável
go install  - Para um executável, realize o build e salva em $GOPATH/bin
            Para um pacote, realize o build e salva em $GOPATH/pkg
            Cria archive files (arquivo.a), os arquivos pré-compilados utilizados pelo imports
- flags
            - "-race" para verificação de Race Conditions

// =========
Compilando para Windows => GOOS=windows GOARCH=amd64 go build main.go
- GOOS
- GOARCH