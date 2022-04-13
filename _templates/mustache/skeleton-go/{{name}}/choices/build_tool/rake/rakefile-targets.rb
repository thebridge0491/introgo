# Targets rakefile script.
require 'rake/clean'
require 'rake/packagetask'

[CLEAN, CLOBBER, Rake::FileList::DEFAULT_IGNORE_PATTERNS].each{|a| a.clear}
CLEAN.include("**/*.o", "*.log", "**/.coverage", "bin", "pkg", "#{VARS.pkg}.test")
CLOBBER.include('build/*', 'build/.??*')

desc 'Help info'
task :help do
  puts "===== subproject: #{VARS.parent}-#{VARS.pkg} =====\nHelp: #{RAKE} [DEBUG=#{ENV['DEBUG']}] [task]"
  sh "#{RAKE} -T"
end

desc 'Run tests: rake test\[topt1,topt2\]'
task :test, [:topt1] => ["./#{VARS.pkg}.test"] do |t, topts|
  if "1" == "#{DEBUG}"
    sh "LD_LIBRARY_PATH=#{ENV['LD_LIBRARY_PATH']}:lib ./#{VARS.pkg}.test -test.coverprofile=build/cover_#{VARS.pkg}.out #{topts[:topt1]} #{topts.extras.join(' ')} || true"
  else
    sh "LD_LIBRARY_PATH=#{ENV['LD_LIBRARY_PATH']}:lib ./#{VARS.pkg}.test #{topts[:topt1]} #{topts.extras.join(' ')} || true"
  end
end

#----------------------------------------
desc 'Uninstall artifacts: rake uninstall\[opt1,opt2\]'
task :uninstall, [:opt1] do |t, opts|
  `go list .../#{VARS.pkg}`.split().each { |pkgX|
    sh "go clean -i #{opts[:opt1]} #{opts.extras.join(' ')} #{pkgX} || true"
    sh "go list -e #{pkgX} || true"
  }
end

desc 'Install artifacts: rake install\[opt1,opt2\]'
task :install, [:opt1] do |t, opts|
  `go list .../#{VARS.pkg}`.split().each { |pkgX|
    sh "PKG_CONFIG_PATH=#{PREFIX}/lib/pkgconfig go install #{opts[:opt1]} #{opts.extras.join(' ')} #{pkgX} || true"
    sh "go list -e #{pkgX} || true"
  }
end

file "build/#{VARS.name}-#{VARS.version}" do |p|
  mkdir_p(p.name)
  # sh "zip -9 -q -x @exclude.lst -r - . | unzip -od #{p.name} -"
  sh "tar --posix -h -X exclude.lst -cf - . | tar -xpf - -C #{p.name}"
end
if defined? Rake::PackageTask
  Rake::PackageTask.new(VARS.proj, VARS.version) do |p|
    # task("build/#{VARS.proj}-#{VARS.version}").invoke
    
    ENV.fetch('FMTS', 'tar.gz,zip').split(',').each{|fmt|
      if p.respond_to? "need_#{fmt.tr('.', '_')}="
        p.send("need_#{fmt.tr('.', '_')}=", true)
      else
        p.need_tar_gz = true
      end
    }
    task(:package).add_description "[FMTS=#{ENV.fetch('FMTS', 'tar.gz,zip')}]"
    task(:repackage).add_description "[FMTS=#{ENV.fetch('FMTS', 'tar.gz,zip')}]"
  end
else
  desc "[FMTS=#{ENV.fetch('FMTS', 'tar.gz,zip')}] Package project distribution"
  task :package => ["#{VARS.proj}-#{VARS.version}"] do |t|
    distdir = "#{VARS.proj}-#{VARS.version}"
    
    ENV.fetch('FMTS', 'tar.gz,zip').split(',').each{|fmt|
      case fmt
      when '7z'
        rm_rf("build/#{distdir}.7z") || true
        cd('build') {sh "7za a -t7z -mx=9 #{distdir}.7z #{distdir}" || true}
      when 'zip'
        rm_rf("build/#{distdir}.zip") || true
        cd('build') {sh "zip -9 -q -r #{distdir}.zip #{distdir}" || true}
      else
        # tarext = `echo #{fmt} | grep -e '^tar$' -e '^tar.xz$' -e '^tar.zst$' -e '^tar.bz2$' || echo tar.gz`.chomp
        tarext = fmt.match(%r{(^tar$|^tar.xz$|^tar.zst$|^tar.bz2$)}) ? fmt : 'tar.gz'
        rm_rf("build/#{distdir}.#{tarext}") || true
        cd('build') {sh "tar --posix -L -caf #{distdir}.#{tarext} #{distdir}" || true}
      end
    }
  end
end

desc 'Generate API documentation: rake doc\[opt1,opt2\]'
task :doc, [:opt1] do |t, opts|
  rm_f("build/doc_#{VARS.pkg}.txt")
#  #serve docs at http://localhost:6060/#{VARS.pkg}
#  sh "go doc -http=:6060 || true"
  `go list .../#{VARS.pkg}`.split().each { |pkgX|
    sh "go doc -all #{opts[:opt1]} #{opts.extras.join(' ')} #{pkgX} >> build/docs/#{VARS.pkg}.txt || true"
  }
end

desc 'Lint check: rake lint\[opt1,opt2\]'
task :lint, [:opt1] do |t, opts|
  rm_f("build/lint_#{VARS.pkg}.txt")
  `go list .../#{VARS.pkg}`.split().each { |pkgX|
    sh "#{ENV['GOPATH']}/bin/golint #{opts[:opt1]} #{opts.extras.join(' ')} #{pkgX} || true"
  }
end

desc 'Report code coverage'
task :report => "build/cover_#{VARS.pkg}.out" do
  sh "go tool cover -html=#{t.source} -o build/cover_#{VARS.pkg}.html || true"
  sh "go tool cover -func=#{t.source} || true"
end
