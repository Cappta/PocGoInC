#include "golang.h"

__declspec(dllexport) void PrintInt(int x) {
    GoPrintInt(x);
}