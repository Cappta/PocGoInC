#include "golang.h"
#include <stdio.h>

__declspec(dllexport) char** TestStringSlice() {
    char** result = GoTestStringSlice();

    printf("\r\nMainDll %x\r\n", result);
    for(int i=0; result[i] != NULL; i++){
        printf("%d: %s\r\n", i, result[i]);
    }

    return result;
}