# POC enriquecimento log Go

Esse projeto tem como objetivo avaliar as soluções de log com JSON para verificar as características desejáveis em um contexto de padronização.

## Solução existentes identificadas

Soluções existentes para o padrão JSON que é uma das características desejáveis iniciais:

- [Logrus](https://github.com/sirupsen/logrus)
- [Zap](https://github.com/uber-go/zap)
- [Zerolog](https://github.com/rs/zerolog)

A primeira rodada será realizada com o `Zerolog` que é de acordo com os benchmarks das próprias dependências, o mais performático entre os 3.

## Características desejáveis

- Padrão de log no formato JSON
- Padronização de valores e nome de atributos comuns
- Enriquecimento de informação no log em um contexto global
- Enriquecimento de informação no momento do log
- Máscara de dados sensíveis
- Externalização de configurações

### Execução do projeto com o Zerolog

Acessar o projeto `app-test-zerolog` e executar no terminal o comando: `go run main.go`
Em outro terminal, pode ser realizado o `curl` para chamar um endpoint de `Hello World` que adiciona um log: `curl localhost:8888`

### Resultado Zerolog

- Padrão de log no formato JSON
    - Já é dado por padrão, sem a necessidade de configurações adicionais
- Padronização de valores e nome de atributos comuns
    - É possível. Foi realizando alterando o nome do campo timestamp para `timestamp_custom`. O formato do valor também pôde ser alterado.
- Enriquecimento de informação no log em um contexto global
    - É possível, foi adicionado o campo `application` na POC.
- Enriquecimento de informação no momento do log
    - É possível. No teste foi adicionado um campo `cpf` na chamada da API.
- Máscara de dados sensíveis
    - De forma transparente, não é possível interceptar os valores enriquecidos e substituir por um valor mascarado. O mecanismo de Hook do `Zerolog` não fornece esse tipo de acesso. Poderiam ser criadas funcionalidades específicas de log da dependência, que já validem esse mecanismo através de configurações externas, para logar já adicionando os campos com valores ofuscasdos. Serão validadas outras soluções antes de avançar em um mecanismo desses, que pode ser problemático em termos de manutenção e evolução, além de manter disponível os acessos aos métodos originais que não teriam a lógica de ofuscação.
- Externalização de configurações
    - É possível com a ajuda do [Viper](https://github.com/spf13/viper). O campo `application` foi adicionado com base em uma configuração de um arquivo `applicationProperties.yaml` na aplicação final.