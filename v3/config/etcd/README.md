# ETCD

```
etcdctl put /key {"""k1""":"""v1""""}
```
```
etcdctl put /key {"""k1""":"""v2""""}
etcdctl put /key {"""k1""":"""v2"""","""k2""":"""v3"""}
```

```
go run .
2021-12-17 11:59:45  file=etcd/main.go:24 level=info {"k1":"v1"}
2021-12-17 11:59:55  file=etcd/main.go:38 level=info [watcher] {"k1":"v2"}
2021-12-17 11:59:58  file=etcd/main.go:38 level=info [watcher] {"k1":"v2","k2":"v3"}
```