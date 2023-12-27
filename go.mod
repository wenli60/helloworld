module git.woa.com/test/helloworld

go 1.12

replace git.woa.com/trpcprotocol/test/helloworld => ./stub/git.woa.com/trpcprotocol/test/helloworld

require (
	git.code.oa.com/trpc-go/trpc-database/gorm v0.2.5
	git.code.oa.com/trpc-go/trpc-filter/debuglog v0.1.7
	git.code.oa.com/trpc-go/trpc-filter/recovery v0.1.4
	git.code.oa.com/trpc-go/trpc-go v0.16.0
	git.woa.com/trpcprotocol/test/helloworld v0.0.0-00010101000000-000000000000
	github.com/astaxie/beego v1.12.3
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.9.1
	github.com/go-gormigrate/gormigrate/v2 v2.1.1
	github.com/golang/mock v1.6.0
	github.com/gopherjs/gopherjs v0.0.0-20181103185306-d547d1d9531e // indirect
	github.com/smartystreets/assertions v0.0.0-20190116191733-b6c0e53d7304 // indirect
	github.com/ugorji/go v1.1.5-pre // indirect
	gorm.io/gorm v1.25.5
)
