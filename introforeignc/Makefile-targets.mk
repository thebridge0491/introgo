# Targets Makefile script.
#----------------------------------------
# Common automatic variables legend (GNU make: make (Linux) gmake (FreeBSD)):
# $* - basename (cur target)  $^ - name(s) (all depns)  $< - name (1st depn)
# $@ - name (cur target)      $% - archive member name  $? - changed depns

FMTS ?= tar.gz,zip
distdir = $(parent)-$(pkg).$(version)

build/$(distdir) : 
	-@mkdir -p build/$(distdir) ; cp -f exclude.lst build
#	#-zip -9 -q --exclude @exclude.lst -r - . | unzip -od build/$(distdir) -
	-tar --format=posix --dereference --exclude-from=exclude.lst -cf - . | tar -xpf - -C build/$(distdir)

.PHONY: help clean test uninstall install dist doc lint report
help: ## help
	@echo "##### subproject: $(parent)-$(pkg) #####"
	@echo "Usage: $(MAKE) [target] -- some valid targets:"
#	-@for fileX in $(MAKEFILE_LIST) `if [ -z "$(MAKEFILE_LIST)" ] ; then echo Makefile ./Makefile-targets.mk ; fi` ; do \
#		grep -ve '^[A-Z]' $$fileX | awk '/^[^.%][-A-Za-z0-9_]+[ ]*:.*$$/ { print "...", substr($$1, 1, length($$1)) }' | sort ; \
#	done
	-@for fileX in $(MAKEFILE_LIST) `if [ -z "$(MAKEFILE_LIST)" ] ; then echo Makefile ./Makefile-targets.mk ; fi` ; do \
		grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $$fileX | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "%-25s%s\n", $$1, $$2}' ; \
	done
clean: ## clean build artifacts
	-rm -rf build/* build/.??* core* *.log .coverage bin pkg $(pkg).test
test: ./$(pkg).test ## run test [TOPTS=""]
#	export [DY]LD_LIBRARY_PATH=. # ([da|ba|z]sh Linux)
#	setenv [DY]LD_LIBRARY_PATH . # (tcsh FreeBSD)
	-if [ "1" = "$(DEBUG)" ] ; then \
		LD_LIBRARY_PATH=$(LD_LIBRARY_PATH):lib ./$(pkg).test \
			-test.coverprofile=build/cover_$(pkg).out $(TOPTS) ; \
	else \
		LD_LIBRARY_PATH=$(LD_LIBRARY_PATH):lib ./$(pkg).test $(TOPTS) ; \
	fi
uninstall install: ## [un]install artifacts [OPTS=""]
	-for pkgX in `go list -e .../$(pkg)` ; do \
		if [ "uninstall" = "$@" ] ; then \
			go clean -i $(OPTS) $$pkgX ; \
		else \
			PKG_CONFIG_PATH=$(PREFIX)/lib/pkgconfig go install $(OPTS) $$pkgX ; \
			go list -e $$pkgX ; \
		fi \
	done
dist:  | build/$(distdir) ## [FMTS="tar.gz,zip"] archive source code
	-@for fmt in `echo $(FMTS) | tr ',' ' '` ; do \
		case $$fmt in \
			7z) echo "### build/$(distdir).7z ###" ; \
				rm -f build/$(distdir).7z ; \
				(cd build ; 7za a -t7z -mx=9 $(distdir).7z $(distdir)) ;; \
			zip) echo "### build/$(distdir).zip ###" ; \
				rm -f build/$(distdir).zip ; \
				(cd build ; zip -9 -q -r $(distdir).zip $(distdir)) ;; \
			*) tarext=`echo $$fmt | grep -e '^tar$$' -e '^tar.xz$$' -e '^tar.bz2$$' || echo tar.gz` ; \
				echo "### build/$(distdir).$$tarext ###" ; \
				rm -f build/$(distdir).$$tarext ; \
				(cd build ; tar --posix -h -caf $(distdir).$$tarext $(distdir)) ;; \
		esac \
	done
	-@rm -r build/$(distdir)
doc: ## generate documentation [OPTS=""]
	-rm -f build/doc_$(pkg).txt
#	#serve docs at http://localhost:6060/$(pkg)
#	#-go doc -http=:6060
	-for pkgX in `go list -e .../$(pkg)` ; do \
		go doc -all $(OPTS) $$pkgX >> build/doc_$(pkg).txt ; \
	done
lint: ## lint check [OPTS=""]
#	-go vet $(OPTS)
	-rm -fr build/lint_$(pkg).txt
	-for pkgX in `go list -e .../$(pkg)` ; do \
		$(GOPATH)/bin/golint $(OPTS) $$pkgX >> build/lint_$(pkg).txt ; \
	done
report: build/cover_$(pkg).out ## report code coverage
	-go tool cover -html=$< -o build/cover_$(pkg).html
	-go tool cover -func=$<
