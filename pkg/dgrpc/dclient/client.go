package main

import (
	"demo/dgrpc/proto"
	"flag"
	"fmt"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"strings"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
)

var (
	ServiceName = flag.String("ServiceName", "greet_service", "service name")                                                     //服务名称
	EtcdAddr    = flag.String("EtcdAddr", "172.21.219.70:20000;172.21.219.70:20002;172.21.219.70:20004", "register etcd address") //etcd的地址
)

var cli *clientv3.Client

// etcd解析器
type etcdResolver struct {
	etcdAddr   string
	clientConn resolver.ClientConn
}

// 初始化一个etcd解析器
func newResolver(etcdAddr string) resolver.Builder {
	return &etcdResolver{etcdAddr: etcdAddr}
}

func (r *etcdResolver) Scheme() string {
	return "etcd"
}

// ResolveNow watch有变化以后会调用
func (r *etcdResolver) ResolveNow(rn resolver.ResolveNowOptions) {
	log.Println("ResolveNow")
	fmt.Println(rn)
}

// Close 解析器关闭时调用
func (r *etcdResolver) Close() {
	log.Println("Close")
}

// Build 构建解析器 grpc.Dial()同步调用
func (r *etcdResolver) Build(target resolver.Target, clientConn resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	var err error
	fmt.Println("call build...")
	//构建etcd client
	if cli == nil {
		cli, err = clientv3.New(clientv3.Config{
			Endpoints:   strings.Split(r.etcdAddr, ";"),
			DialTimeout: 15 * time.Second,
		})
		if err != nil {
			fmt.Printf("连接etcd失败：%s\n", err)
			return nil, err
		}
	}

	r.clientConn = clientConn

	go r.watch(fmt.Sprintf("/%s/%s/", target.Scheme, target.Endpoint()))

	return r, nil
}

// 监听etcd中某个key前缀的服务地址列表的变化
func (r *etcdResolver) watch(keyPrefix string) {
	//初始化服务地址列表
	var addrList []resolver.Address

	resp, err := cli.Get(context.Background(), keyPrefix, clientv3.WithPrefix())
	if err != nil {
		fmt.Println("获取服务地址列表失败：", err)
	} else {
		for i := range resp.Kvs {
			addrList = append(addrList, resolver.Address{Addr: strings.TrimPrefix(string(resp.Kvs[i].Key), keyPrefix)})
		}
	}

	r.clientConn.NewAddress(addrList)

	//监听服务地址列表的变化
	rch := cli.Watch(context.Background(), keyPrefix, clientv3.WithPrefix())
	for n := range rch {
		for _, ev := range n.Events {
			addr := strings.TrimPrefix(string(ev.Kv.Key), keyPrefix)
			switch ev.Type {
			case mvccpb.PUT:
				if !exists(addrList, addr) {
					addrList = append(addrList, resolver.Address{Addr: addr})
					r.clientConn.NewAddress(addrList)
				}
				fmt.Println("有新的服务注册：", addr)
			case mvccpb.DELETE:
				if s, ok := remove(addrList, addr); ok {
					addrList = s
					r.clientConn.NewAddress(addrList)
				}
				fmt.Println("服务注销：", addr)
			}
		}
	}
}

func exists(l []resolver.Address, addr string) bool {
	for i := range l {
		if l[i].Addr == addr {
			return true
		}
	}
	return false
}

func remove(s []resolver.Address, addr string) ([]resolver.Address, bool) {
	for i := range s {
		if s[i].Addr == addr {
			s[i] = s[len(s)-1]
			return s[:len(s)-1], true
		}
	}
	return nil, false
}

func main() {
	flag.Parse()

	//注册etcd解析器
	r := newResolver(*EtcdAddr)
	resolver.Register(r)

	//客户端连接服务器(负载均衡：轮询) 会同步调用r.Build()
	conn, err := grpc.Dial(r.Scheme()+":///"+*ServiceName, grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
		grpc.WithInsecure()) // 负载均衡
	//conn, err := grpc.Dial("192.168.1.45:3000", grpc.WithInsecure()) //直连
	if err != nil {
		fmt.Println("连接服务器失败：", err)
		return
	}
	defer conn.Close()

	//获得grpc句柄
	c := proto.NewGreetClient(conn)
	ticker := time.NewTicker(5 * time.Second)
	i := 1
	for range ticker.C {
		resp1, err := c.Hello(
			context.Background(),
			&proto.GreetRequest{Name: fmt.Sprintf("张三%d", i)},
		)
		if err != nil {
			fmt.Println("Hello调用失败：", err)
			return
		}
		fmt.Printf("Hello 响应：%s，来自：%s\n", resp1.Message, resp1.From)

		i++
	}
}
