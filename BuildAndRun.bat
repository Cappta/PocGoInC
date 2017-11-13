go build -buildmode=c-archive golang.go
gcc -shared -pthread maindll.c golang.a -o main.dll -lWinMM -lntdll -lWS2_32
del golang.a
gcc -pthread main.c -o main -lWinMM -lntdll -lWS2_32
main.exe