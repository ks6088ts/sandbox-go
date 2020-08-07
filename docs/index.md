# Visual Studio Code

## Remote container

1. `cmd + shift + p` to open command palette
1. select `install extension`
1. install [Remote - Containers](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers), `ms-vscode-remote.remote-containers`

### Links
* [Docker Compose keep container running](https://stackoverflow.com/a/55953120)
* [microsoft/vscode-remote-try-go](https://github.com/microsoft/vscode-remote-try-go)

## Settings

1. `cmd + ,` to open settings

### Links
* [GitHub: vscode-textbook](https://github.com/vscode-textbook)

# Go Modules

```bash
# help
go help mod

# add missing and remove unused modules
go mod tidy
```

# gRPC

## Links
* [Go plugin for the protocol compiler](https://grpc.io/docs/languages/go/quickstart/#prerequisites)
* [Protocol Buffer Compiler Installation](https://grpc.io/docs/protoc-installation/#install-using-a-package-manager)
* [grpc-go/examples/helloworld](https://github.com/grpc/grpc-go/tree/master/examples/helloworld)

# Grafana

## Links

* [Configure with environment variables](https://grafana.com/docs/grafana/latest/administration/configuration/#configure-with-environment-variables)
* [Provisioning Grafana:Data sources](https://grafana.com/docs/grafana/latest/administration/provisioning/#data-sources)

# Prometheus

## Links

* [Prometheus:GETTING STARTED](https://prometheus.io/docs/prometheus/latest/getting_started/#getting-started)

# GraphQL

## Links

* [春の入門祭り🌸 #7 作って学ぶGraphQL。gqlgenを用いて鉄道データ検索API開発入門](https://future-architect.github.io/articles/20200609/)

# PostgreSQL

```bash
docker-compose -f docker-compose.gql.yml up -d
docker-compose -f docker-compose.gql.yml exec postgresql \
    bash -c "psql -U user"
```

```sql
# list tables
\dt;

# show table
\d <table>;

# select data from table
select * from <table>;
```

## Links

* [PostgreSQLの基本的なコマンド](https://qiita.com/H-A-L/items/fe8cb0e0ee0041ff3ceb)
* [駅データ.jp > ダウンロード](https://ekidata.jp/dl/)
# Docker

```bash
docker-compose build cli
docker run --rm sandbox-go_cli ./cli --help
```

## Links
* [Networking in Compose: Specify custom networks](https://docs.docker.com/compose/networking/#specify-custom-networks)
* [Issues with COPY when using multistage Dockerfile builds — no such file or directory](https://stackoverflow.com/a/50070187)
* [Issue with Docker multi-stage builds](https://stackoverflow.com/a/56057877)
* [Using cgo with the go command](https://golang.org/cmd/cgo/#hdr-Using_cgo_with_the_go_command)

# xo

```bash
# start PostgreSQL server
docker-compose -f docker-compose.gql.yml up -d postgresql

# generate code for a postgres schema
mkdir -p models
xo "pgsql://user:password@localhost/db?sslmode=disable" -o models
```

## Links

* [xo/xo](https://github.com/xo/xo)
