#include <stdio.h>
#include "GoDLL.h"

// force gcc to link in go runtime (may be a better solution than this)
void dummy() {
    PrintBye();

    PrintHello((GoString){"", 0});
    int i = GetIntFromDLL();
}

int main() {

}