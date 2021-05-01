# Raspberry-Pi

## CPU information

processor	: 3
model name	: ARMv7 Processor rev 4 (v7l)
BogoMIPS	: 38.40
Features	: half thumb fastmult vfp edsp neon vfpv3 tls vfpv4 idiva idivt vfpd32 lpae evtstrm crc32
CPU implementer	: 0x41
CPU architecture: 7
CPU variant	: 0x0
CPU part	: 0xd03
CPU revision	: 4

Hardware	: BCM2835
Revision	: a22082
Serial		: 000000002de95cb6
Model		: Raspberry Pi 3 Model B Rev 1.2

--- 

```bash
env GOOS=linux GOARCH=arm GOARM=5 go build -o hello-world main.go 
```

hello-world: ELF 32-bit LSB executable, ARM, EABI5 version 1 (SYSV), statically linked, Go BuildID=b5zm4ZKh9BPrEgwbiVpf/gyf6xAJcY1q-I0hK5xbN/ZFDcF2a5z6aIT11PjFxF/whbPzB160c6Cw7u1f_do, not stripped

```bash
GOARCH=arm64 GOOS=linux go build -o hello-world main.go 
```

hello-world: ELF 64-bit LSB executable, ARM aarch64, version 1 (SYSV), statically linked, not stripped

```C
#include <stdio.h>
int main()
{
   printf("Hello, World!");
   return 0;
}
```

hello: ELF 32-bit LSB executable, ARM, EABI5 version 1 (SYSV), dynamically linked, interpreter /lib/ld-linux-armhf.so.3, for GNU/Linux 3.2.0, BuildID[sha1]=630c0bb5ef2b9bc051fe0b8d0c1034a4e166d7b3, not stripped




[Tutorial](https://medium.com/@farissyariati/go-raspberry-pi-hello-world-tutorial-7e830d08b3ae)












