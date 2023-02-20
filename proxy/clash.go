package proxy

import (
	"context"
	"fmt"
	"github.com/Dreamacro/clash/adapter/outbound"
	"github.com/Dreamacro/clash/listener/mixed"
	"io"
	"net"

	"github.com/Dreamacro/clash/constant"
)

func Run() {
	in := make(chan constant.ConnContext, 100)
	defer close(in)

	l, err := mixed.New("0.0.0.0:20001", in)
	if err != nil {
		panic(err)
	}
	defer l.Close()
	println("listen at:", l.Address())
	socksOption := &outbound.Socks5Option{
		Name:   "test",
		Server: "170.106.110.127",
		Port:   10130,
		//UserName: "uat_team-zone-custom",
		//Password: "314f71aaa600567a315b405626102ed4",
	}
	direct := outbound.NewSocks5(*socksOption)
	//direct := outbound.NewDirect()
	//auth.SetAuthenticator(auth2.NewAuthenticator([]auth2.AuthUser{{User: "testAdmin", Pass: "904028"}}))
	//auth.SetAuthenticator(simpleAuth{})
	for c := range in {
		conn := c
		metadata := conn.Metadata()
		fmt.Printf("request incoming from %s to %s\n", metadata.SourceAddress(), metadata.RemoteAddress())
		go func() {
			remote, err := direct.DialContext(context.Background(), metadata)
			if err != nil {
				fmt.Printf("dial error: %s\n", err.Error())
				return
			}
			relay(remote, conn.Conn())
		}()
	}
}

type simpleAuth struct {
}

func (s simpleAuth) Verify(user string, pass string) bool {
	return true
}

func (s simpleAuth) Users() []string {
	return []string{}
}

func relay(l, r net.Conn) {
	go io.Copy(l, r)
	io.Copy(r, l)
}
