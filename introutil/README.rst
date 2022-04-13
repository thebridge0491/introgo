Introgo/Introutil
===========================================
.. .rst to .html: rst2html5 foo.rst > foo.html
..                pandoc -s -f rst -t html5 -o foo.html foo.rst

Utilities sub-package for Golang Intro examples project.

Installation
------------
source code tarball download:
    
        # [aria2c --check-certificate=false | wget --no-check-certificate | curl -kOL]
        
        FETCHCMD='aria2c --check-certificate=false'
        
        $FETCHCMD https://bitbucket.org/thebridge0491/introgo/[get | archive]/master.zip

version control repository clone:
        
        git clone https://bitbucket.org/thebridge0491/introgo.git

build example with rake:
cd $GOPATH/src/<path> ; [sh] ./configure.sh [--prefix=$PREFIX] [--help]

[PKG_CONFIG_PATH=$PREFIX/lib/pkgconfig] [sudo] rake install [test]

build example with make:
cd $GOPATH/src/<path> ; [sh] ./configure.sh [--prefix=$PREFIX] [--help]

[PKG_CONFIG_PATH=$PREFIX/lib/pkgconfig] [sudo] make install [test]

Usage
-----
        // PKG_CONFIG='pkg-config --with-path=$PREFIX/lib/pkgconfig'
        
        // export GOPATH := $(GOPATH):<path>
        
        // export CGO_CFLAGS += $PKG_CONFIG --cflags <ffi-lib>
        
        // export CGO_LDFLAGS += $PKG_CONFIG --libs <ffi-lib>
        
        import ( "bitbucket.org/thebridge0491/introgo/introutil" )
        
        var (arr1 []int = []int{0, 1, 2} ; arr2 []int = []int{10, 20, 30})
        
        var arrProd = introutil.CartesianProd(arr1, arr2)

Author/Copyright
----------------
Copyright (c) 2015 by thebridge0491 <thebridge0491-codelab@yahoo.com>

License
-------
Licensed under the Apache-2.0 License. See LICENSE for details.
