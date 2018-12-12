#include <windows.h>

typedef void (*pfn)(void);

int main(int argc, char* argv[]) {
    pfn fn;

    HMODULE h = LoadLibrary("hello.so");
    fn = (pfn) GetProcAddress(h, "Hello");

    fn();

    return 0;
}
