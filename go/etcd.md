# etcd log show code line number
```
etcd --debug --log-output=stdout
```
Source code (embed/config.go):
```
func (cfg *Config) SetupLogging() {
  ...
  switch cfg.LogOutput {
  case "stdout":
    capnslog.SetFormatter(capnslog.NewPrettyFormatter(os.Stdout, cfg.Debug))
  case "stderr":
    capnslog.SetFormatter(capnslog.NewPrettyFormatter(os.Stderr, cfg.Debug))
  case DefaultLogOutput:
  default:
    plog.Panicf(`unknown log-output %q (only supports %q, "stdout", "stderr")`, cfg.LogOutput, DefaultLogOutput)
  }
```
In file: capnslog/formatters.go
```
func (c *PrettyFormatter) Format(pkg string, l LogLevel, depth int, entries ...interface{}) {
  ...
  if c.debug {
    _, file, line, ok := runtime.Caller(depth) // It's always the same number of frames to the user's call.
    if !ok {
      file = "???"
      line = 1
    } else {
      slash := strings.LastIndex(file, "/")
      if slash >= 0 {
        file = file[slash+1:]
      }
    }
    if line < 0 {
      line = 0 // not a real line number
    }
    c.w.WriteString(fmt.Sprintf(" [%s:%d]", file, line))
  }
  ...
}
```

# find a bug in etcd: the log code:line is not the place to call logger
```
2018-03-26 01:56:12.756412 [pkg_logger.go:124] N | embed: serving insecure client requests on 172.20.1.14:12379, this is strongly discouraged!
2018-03-26 01:56:59.898028 [pkg_logger.go:104] E | rafthttp: request cluster ID mismatch (got 9fc6e94b3ca3d354 want 931d299ef5b15a11)
```
Where is the code that log error?

The reason why logger know the call site, is: it uses runtime.Call(depth), see capnslog/formatters.go:
```
func (c *PrettyFormatter) Format(pkg string, l LogLevel, depth int, entries ...interface{}) {
  ...
  if c.debug {
    _, file, line, ok := runtime.Caller(depth) // It's always the same number of frames to the user's call.
    ...
```

To fix: trace how to pass the `depth` parameter, finally we can find in file capnslog/pkg_logger.go:27
```
const calldepth = 2
```
Change the value to 3, and build:
```
$ go install github.com/coreos/etcd/vendor/github.com/coreos/pkg/capnslog
$ go install github.com/coreos/etcd
```
Now the log can show the correct place that call the log:
```
2018-03-26 18:36:09.633165 [serve.go:124] N | embed: serving insecure client requests on 127.0.0.1:12379, this is strongly discouraged!
2018-03-26 18:44:51.307961 [http.go:331] E | rafthttp: request cluster ID mismatch (got 9fc6e94b3ca3d354 want 931d299ef5b15a11)
```

# Start etcd cluster with 3 nodes
Assume we want to start etcd cluster in 3 servers: 172.20.1.[2,4,8], with client port 12379 and peer-to-peer port 12380.

vi ~/.bashrc
```
# etcd setting
port_cli=12379  # client port
port_pp=12380   # peer to peer port
http_404=http://172.20.1.14
http_408=http://172.20.1.18
http_402=http://172.20.1.12
etcd_debug="" # --debug --log-output=stdout
export ETCDCTL_API=3
export ETCD_INITIAL_CLUSTER="fnode404=$http_404:$port_pp,fnode408=$http_408:$port_pp,fnode402=$http_402:$port_pp"
export ETCD_INITIAL_CLUSTER_STATE=new

alias etcd_fnode404="etcd $etcd_debug --name fnode404 --initial-advertise-peer-urls $http_404:$port_pp --listen-peer-urls $http_404:$port_pp --listen-client-urls $http_404:$port_cli --advertise-client-urls $http_404:$port_cli --initial-cluster-token etcd-cluster-1"
alias etcd_fnode408="etcd $etcd_debug --name fnode408 --initial-advertise-peer-urls $http_408:$port_pp --listen-peer-urls $http_408:$port_pp --listen-client-urls $http_408:$port_cli --advertise-client-urls $http_408:$port_cli --initial-cluster-token etcd-cluster-1"
alias etcd_fnode402="etcd $etcd_debug --name fnode402 --initial-advertise-peer-urls $http_402:$port_pp --listen-peer-urls $http_402:$port_pp --listen-client-urls $http_402:$port_cli --advertise-client-urls $http_402:$port_cli --initial-cluster-token etcd-cluster-1"
```
then `source ~/.bashrc` in each server. And then run:
```
$ etcd_fnode404 (or 408, 402 in the related server)
```

## Bug: etcd --name xxxx-xxx
The name cannot have '-', or else fail.

## Pure command line
```
ETCDCTL_API=3 etcd  --name=fnode404 --initial-advertise-peer-urls=http://172.20.1.14:12380 --listen-peer-urls=http://172.20.1.14:12380 --listen-client-urls=http://172.20.1.14:12379 --advertise-client-urls=http://172.20.1.14:12379 --initial-cluster-token=etcd-cluster-1 --initial-cluster=http://172.20.1.14:12380,fnode408=http://172.20.1.18:12380,fnode402=http://172.20.1.12:12380 --initial-cluster-state=new
```

## issue when start cluster: "rafthttp: request cluster ID mismatch"
Reason: the old data dir mismatcher

Solution: remove the old data dir (--data-dir, default is: ${name}.etcd/)

## etcdctl show cluster member list:
```
$ etcdctl --endpoints=172.20.1.14:12379 member list
b5af032ec7bb1a1d, started, fnode404, http://172.20.1.14:12380, http://172.20.1.14:12379
d4bff22f96299b61, started, fnode408, http://172.20.1.18:12380, http://172.20.1.18:12379
```

# Code Study
## Server handling request
src/github.com/coreos/etcd/etcdserver/v3_server.go
```
func (s *EtcdServer) Put(ctx context.Context, r *pb.PutRequest) (*pb.PutResponse, error) {
  resp, err := s.raftRequest(ctx, pb.InternalRaftRequest{Put: r})
  if err != nil {
    return nil, err
  }
  return resp.(*pb.PutResponse), nil
}
```

# 集群大小与容错

集群的大小指集群节点的个数。根据 etcd 的分布式数据冗余策略，集群节点越多，容错能力(Failure Tolerance)越强，同时写性能也会越差。
所以关于集群大小的优化，其实就是容错和写性能的一个平衡。

另外， etcd 推荐使用 奇数 作为集群节点个数。因为奇数个节点与和其配对的偶数个节点相比(比如 3节点和4节点对比)，
容错能力相同，却可以少一个节点。

所以综合考虑性能和容错能力，etcd 官方文档推荐的 etcd 集群大小是 3, 5, 7。至于到底选择 3,5 还是 7，根据需要的容错能力而定。

关于节点数和容错能力对应关系，如下表所示：
```
集群大小  最大容错
1 0
3 1
4 1
5 2
6 2
7 3
8 3
9 4
```
所以，如果启动 2 个 etcd，挂了一个 etcd之后就不能用了，这是 by design 的。

Reference: [搭建 etcd 集群 - 暴走漫画容器实践系列 Part3](https://segmentfault.com/a/1190000003852735)

