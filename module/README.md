# Go BLS Signatures Example w/Go modules

## Set env vars

You will need to set flags to point to wherever you installed the Chia BLS
Signatures C++ library. The header files are needed by the C++ compiler and the
location of the library is needed by the linker.

If you compiled the library by followed the instructions in the README one
level up, you should be able to just use the Makefile to build:

```sh
make build
./module
```
