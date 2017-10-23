// gcc -o teste main.c GoDLL.dll

#include <stdio.h>
#include "windows.h"
#include "GoDLL.h"

HINSTANCE libHandle;

void LoadDLL(void)
{
    libHandle = NULL;

    libHandle = LoadLibrary("GoDLL");

    if(libHandle == NULL) {
        printf("Error loading GoDLL.dll\r\n");
    } else {
		printf("Load successful: GoDLL.dll\r\n");
    }
}

int main(int argc, char *argv[])
{
    LoadDLL();

	PrintBye();

	const char *str = "Jorge";

    PrintHello((GoString){str, strlen(str)});

    printf("From DLL: %d\r\n", GetIntFromDLL());

    return 0;
}