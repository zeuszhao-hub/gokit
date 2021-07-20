package env

type Env string

var Type = struct {
	Production  Env
	Development Env
}{
	"production",
	"development",
}
