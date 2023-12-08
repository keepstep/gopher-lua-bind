# go test -v ./bind --run TestGen
rm -f ./gen_bind
go build -o gen_bind ./tool/gen_bind.go
./gen_bind
# delete no use import
goimports -l  -w ./bind
gofmt -l  -w ./bind
# rm -f ./gen_bind
