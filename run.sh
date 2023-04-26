rm  -f ./main
go build -o main main.go
# ./main --execute="$1"
# ./main --execute="lua/test.lua"
./main --execute="lua/test_bind.lua"
