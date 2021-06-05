module example.com/hello

go 1.16

require (
	example.com/greetings v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.7.2
	golang.org/x/sys v0.0.0-20200930185726-fdedc70b468f // indirect
	golang.org/x/text v0.3.3 // indirect
	rsc.io/quote v1.5.2
)

replace example.com/greetings => ../greetings
