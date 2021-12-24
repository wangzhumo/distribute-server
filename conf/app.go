package conf

import "fmt"

const Ip = ""
const Port = 8990

func Addr() string {
	return fmt.Sprintf("%s:%d", Ip, Port)
}
