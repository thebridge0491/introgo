# FFI auxiliary makefile script

ffi_libdir = $(shell PKG_CONFIG_PATH=$(PREFIX)/lib/pkgconfig pkg-config --variable=libdir intro_c-practice || echo .)
ffi_incdir = $(shell PKG_CONFIG_PATH=$(PREFIX)/lib/pkgconfig pkg-config --variable=includedir intro_c-practice || echo .)
LD_LIBRARY_PATH := $(LD_LIBRARY_PATH):$(ffi_libdir)
export LD_LIBRARY_PATH

export CGO_CFLAGS += `PKG_CONFIG_PATH=$(PREFIX)/lib/pkgconfig pkg-config --cflags intro_c-practice`
export CGO_LDFLAGS += `PKG_CONFIG_PATH=$(PREFIX)/lib/pkgconfig pkg-config --libs intro_c-practice`

.PHONY: prep_swig
pkg/classicc/classicc_wrap.c: classicc.i
	-swig -go -cgo -intgosize 32 -package classicc -v -I$(ffi_incdir) -outdir pkg/classicc -o $@ $<

prep_swig: pkg/classicc/classicc_wrap.c ## prepare Swig files
