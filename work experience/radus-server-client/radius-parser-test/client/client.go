package main

import (
	"context"
	"fmt"
	"net"

	"parspooyesh.com/scm/ibsng/go-lib-radius/radius/rfc2866"

	"parspooyesh.com/ibsng-cgw/src/layeh.com/radius/rfc2865"
	"parspooyesh.com/scm/ibsng/go-lib-radius/radius"
)

func main() {
	for i := 0; i < 20; i++ {
		packet := radius.New(radius.CodeAccountingRequest, []byte(`secret`))
		rfc2865.UserName_SetString(packet, "ParsPooyesh.com")
		rfc2865.Class_AddString(packet, "2155591729|92|502b.7326.d04a|520176440|0|0|344680552|[P:55591729|S:ADSL|C:dPLaaR8pnGJJRhVi2MCr3KcofuuCt4MO234tmdoupgQ_0pbaVQhyW-BSEAWeMfpXHWn|23e;24w]")
		rfc2866.AcctStatusType_Add(packet, 1)
		rfc2865.FramedIPAddress_Add(packet, net.IP([]byte{168, 10, 25, byte(i)}))
		radius.Exchange(context.Background(), packet, "localhost:1813")
		fmt.Println("send", i)
	}
}
