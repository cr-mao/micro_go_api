VERSION = `git rev-parse --short HEAD`
BUILDTIME = `date +%FT%T`

user_serve:
	cd user_web && go run  main.go
goods_serve:
	cd goods_web && go run main.go
order_serve:
	cd order_web && go run main.go

codegen:
	sh scripts/codegen.sh