#include "golang.h"
#include "maindll.h"

DLL_EXPORT void PrintInt(int x) {
    GoPrintInt(x);
}