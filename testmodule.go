package testmodule

import _ "github.com/valyala/fasthttp" //I do realy need an empty import

//ShowDep returns info about mymodule version and it's dependencies
func ShowDep() string {
	return "testmodule v1.0.1, fasthttp v1.15.1"
}
