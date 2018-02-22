#include <windows.h>
#include <stdio.h>

typedef void* (*TestPtr)(void);

int main(int argc, char **argv)
{
    HINSTANCE mainDLL = LoadLibrary("main.dll");
    FARPROC TestAddr = GetProcAddress((HMODULE)mainDLL, "TestStringSlice");

    TestPtr Test;
    Test = (TestPtr)TestAddr;
    char** result = Test();
    
    printf("\r\nMain %x\r\n", result);
    for(int i=0; result[i] != NULL; i++){
        printf("%d: %s\r\n", i, result[i]);
    }

    FreeLibrary(mainDLL);
    return 0;
}
