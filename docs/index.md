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

```
query findStations {
  stationByCD(stationCD: 1110104) {
    stationCD
    stationName
  }
}
```

## Links

* [春の入門祭り🌸 #7 作って学ぶGraphQL。gqlgenを用いて鉄道データ検索API開発入門](https://future-architect.github.io/articles/20200609/)

# PostgreSQL

```bash
docker-compose -f docker-compose.db.yml up -d postgresql
docker-compose -f docker-compose.db.yml exec postgresql \
    bash -c "psql -U user -d db"
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
* [Connecting to a PostgreSQL database with Go's database/sql package](https://www.calhoun.io/connecting-to-a-postgresql-database-with-gos-database-sql-package/)

# InfluxDB

**InfluxDB API**

```bash
curl -i \
    -u user:password \
    -XPOST http://localhost:8086/query \
    --data-urlencode "q=CREATE DATABASE db"

curl -i \
    -u user:password \
    -XPOST 'http://localhost:8086/write?db=db' \
    --data-binary @mock_data.txt

curl -i \
    -u user:password \
    -XPOST http://localhost:8086/query \
    --data-urlencode "q=DROP DATABASE db"
```

**influx command**

```bash
docker-compose -f docker-compose.db.yml exec influxdb bash

# interactive shell
influx --username user --password password --database db
use db
select * from cpu_load_short

# one shot command
influx --username user --password password --database db \
    --execute "select * from cpu_load_short" --format csv
```

## Links
* [Write data with the InfluxDB API](https://docs.influxdata.com/influxdb/v1.8/guides/write_data/)
* [Using influx - InfluxDB command line interface](https://docs.influxdata.com/influxdb/v1.8/tools/shell/)

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
docker-compose -f docker-compose.db.yml up -d postgresql

# generate code for a postgres schema
mkdir -p pkg/gql/xo
xo "pgsql://user:password@postgresql/db?sslmode=disable" -o pkg/gql/xo
```

## Links

* [xo/xo](https://github.com/xo/xo)


# gin

## Links

* [gin-gonic/gin](https://github.com/gin-gonic/gin)
