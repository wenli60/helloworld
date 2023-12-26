module git.woa.com/test/helloworld

go 1.12

replace git.woa.com/trpcprotocol/test/helloworld => ./stub/git.woa.com/trpcprotocol/test/helloworld

require (
	git.code.oa.com/trpc-go/trpc-database/gorm v0.2.5
	git.code.oa.com/trpc-go/trpc-filter/debuglog v0.1.7
	git.code.oa.com/trpc-go/trpc-filter/recovery v0.1.4
	git.code.oa.com/trpc-go/trpc-go v0.16.0
	git.woa.com/trpcprotocol/test/helloworld v0.0.0-00010101000000-000000000000
	github.com/golang/mock v1.6.0
	gorm.io/gorm v1.25.5
)
