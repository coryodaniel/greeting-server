# greeting-server

A greeting server for demoing [Bonny](https://github.com/coryodaniel/bonny), a Kubernetes Operator SDK written in Elixir.

[Docker Image](https://quay.io/bonny/greeting-server)

## Usage

*Building:* 

```shell
go test && go build
```

Using the default greeting:

```shell
./greeting-server
```

* http://localhost:5000/greeting/ - Hello, World
* http://localhost:5000/greeting/Chauncy - Hello, Chauncy

Change the greeting:

```shell
GREETING="Hola" ./greeting-server
```

* http://localhost:5000/greeting/ - Hola, World
* http://localhost:5000/greeting/Chauncy - Hola, Chauncy
