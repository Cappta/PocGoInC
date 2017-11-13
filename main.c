#include <windows.h>

typedef void (*PrintIntPtr)(int);

int main(int argc, char **argv)
{
    HINSTANCE mainDLL = LoadLibrary("main.dll");
    FARPROC PrintIntAddr = GetProcAddress((HMODULE)mainDLL, "PrintInt");

    PrintIntPtr PrintInt;
    PrintInt = (PrintIntPtr)PrintIntAddr;
    PrintInt(5);

    FreeLibrary(mainDLL);
    return 0;
}
