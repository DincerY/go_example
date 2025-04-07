package main

import (
	"fmt"
)

type Abser interface {
	Abs() int32
}

type IPAddr [4]byte

/*func (addr IPAddr) String() string {
	res := ""
	for i := 0; i < len(addr); i++ {
		res += strconv.Itoa(int(addr[i])) + "."
	}
	res = strings.TrimSuffix(res, ".")
	return res
}*/

func (a IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", a[0], a[1], a[2], a[3])
}

func main() {
	var a Abser
	f := MyInt(-10)
	v := Vertex{3, 4}
	a = f
	a = &v

	fmt.Println(a.Abs())

	var nilInterface interface{}
	nilInterface = 1
	nilInterface = " "
	nilInterface = Vertex{3, 5}
	fmt.Println(nilInterface)

	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

type MyInt int32

type Vertex struct {
	x int
	y int
}

func (i MyInt) Abs() int32 {
	if i < 0 {
		return int32(-i)
	} else {
		return int32(i)
	}
}

func (v *Vertex) Abs() int32 {
	res := v.x * v.y
	if res < 0 {
		return int32(-res)
	} else {
		return int32(res)
	}
}
