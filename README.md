# hello-world

A hello world server for demoing the Bonny hello-world-operator.

[Bonny](https://github.com/coryodaniel/bonny) is a Kubernetes Operator SDK written in Elixir.

[Docker Image](https://quay.io/coryodaniel/hello-world)

## Usage

*Building:* 

```shell
go test && go build
```

Using the default greeting:

```shell
./hello-world
```

* http://localhost:5000/greeting/ - Hello, World
* http://localhost:5000/greeting/Chauncy - Hello, Chauncy

Change the greeting:

```shell
GREETING="Hola" ./hello-world
```

* http://localhost:5000/greeting/ - Hola, World
* http://localhost:5000/greeting/Chauncy - Hola, Chauncy
