# Multi-package project Makefile script.
.POSIX:
help:

#MAKE = make # (GNU make variants: make (Linux) gmake (FreeBSD)

#export PATH := $(PATH):$(shell go env GOPATH)/bin
export GOPATH := $(shell go env GOPATH)

parent = introgo
version = 0.1.0
SUBDIRS = introutil introforeignc intropractice intromain

.PHONY: configure build testCompile help clean test uninstall install
configure: $(SUBDIRS) ## configure [OPTS=""]
	-for dirX in $^ ; do (cd $$dirX ; sh ./configure.sh $(OPTS)) ; done
help: $(SUBDIRS)
	-for dirX in $^ ; do $(MAKE) -C $$dirX $@ ; done
	@echo "##### Top-level multiproject: $(parent) #####"
	@echo "       $(MAKE) [SUBDIRS="$(SUBDIRS)"] configure [OPTS=??]"
	@echo "Usage: $(MAKE) [SUBDIRS="$(SUBDIRS)"] [target]"
build: $(SUBDIRS)
	-for dirX in $^ ; do (cd $$dirX ; go build $(OPTS) .../$$dirX) ; done
uninstall: $(SUBDIRS)
	-for dirX in $^ ; do (cd $$dirX ; go clean -i $(OPTS) .../$$dirX) ; done
install: $(SUBDIRS)
	-for dirX in $^ ; do (cd $$dirX ; go install $(OPTS) .../$$dirX) ; done
testCompile: $(SUBDIRS)
	-for dirX in $^ ; do \
		if [ "1" = "$(DEBUG)" ] ; then \
			(cd $$dirX ; go test -c -cover $(OPTS)) ; \
		else \
			(cd $$dirX ; go test -c $(OPTS)) ; \
		fi ; \
	done
test: $(SUBDIRS)
	-for dirX in $^ ; do \
		if [ "1" = "$(DEBUG)" ] ; then \
			(cd $$dirX ; ./$$dirX.test -test.coverprofile=build/cover_$$dirX.out $(TOPTS)) ; \
		else \
			(cd $$dirX ; ./$$dirX.test $(TOPTS)) ; \
		fi ; \
	done
clean: $(SUBDIRS)
	-for dirX in $^ ; do $(MAKE) -C $$dirX $@ ; done
	-rm -fr core* *~ .*~ build/* *.log */*.log

#----------------------------------------
FMTS ?= tar.gz,zip
distdir = $(parent)-$(version)

build/$(distdir) : 
	-@mkdir -p build/$(distdir) ; cp -f exclude.lst build/
#	#-zip -9 -q --exclude @exclude.lst -r - . | unzip -od build/$(distdir) -
	-tar --format=posix --dereference --exclude-from=exclude.lst -cf - . | tar -xpf - -C build/$(distdir)

.PHONY: dist doc lint report run
dist | build/$(distdir): $(SUBDIRS)
	-@for fmt in `echo $(FMTS) | tr ',' ' '` ; do \
		case $$fmt in \
			7z) echo "### build/$(distdir).7z ###" ; \
				rm -f build/$(distdir).7z ; \
				(cd build ; 7za a -t7z -mx=9 $(distdir).7z $(distdir)) ;; \
			zip) echo "### build/$(distdir).zip ###" ; \
				rm -f build/$(distdir).zip ; \
				(cd build ; zip -9 -q -r $(distdir).zip $(distdir)) ;; \
			*) tarext=`echo $$fmt | grep -e '^tar$$' -e '^tar.xz$$' -e '^tar.zst$$' -e '^tar.bz2$$' || echo tar.gz` ; \
				echo "### build/$(distdir).$$tarext ###" ; \
				rm -f build/$(distdir).$$tarext ; \
				(cd build ; tar --posix -L -caf $(distdir).$$tarext $(distdir)) ;; \
		esac \
	done
	-@rm -r build/$(distdir)
	-for dirX in $^ ; do $(MAKE) -C $$dirX $@ ; done
doc lint report: $(SUBDIRS)
	-for dirX in $^ ; do \
		if [ "doc" = "$@" ] ; then \
			rm $$dirX/build/doc_$$dirX.txt ; \
			(cd $$dirX ; go doc -all $(OPTS) $$dirX >> build/doc_$$dirX.txt) ; \
		elif [ "lint" = "$@" ] ; then \
			rm $$dirX/build/lint_$$dirX.txt ; \
			(cd $$dirX ; $(GOPATH)/bin/golint $(OPTS) .../$$dirX >> build/lint_$$dirX.txt) ; \
		elif [ "report" = "$@" ] ; then \
			(cd $$dirX ; go tool cover -html=build/cover_$$dirX.out -o build/cover_$$dirX.html ; go tool cover -func=build/cover_$$dirX.out) ; \
		fi ; \
	done
run: $(SUBDIRS)
	-for dirX in $^ ; do \
		if [ -e $(GOPATH)/bin/$$dirX ] ; then \
			(cd $$dirX ; $(GOPATH)/bin/$$dirX $(ARGS)); \
		fi ; \
	done
