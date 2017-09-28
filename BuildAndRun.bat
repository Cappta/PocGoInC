go build -buildmode=c-archive main.go
gcc -pthread main.c main.a -o main -lWinMM -lntdll -lWS2_32
main.exe