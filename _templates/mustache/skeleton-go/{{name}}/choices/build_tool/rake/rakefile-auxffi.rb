# FFI auxiliary rakefile script
PREFIX = PREFIX ? PREFIX : ENV['prefix']
FFI_LIBDIR = `PKG_CONFIG_PATH=#{PREFIX}/lib/pkgconfig pkg-config --variable=libdir intro_c-practice || echo .`.chomp
FFI_INCDIR = `PKG_CONFIG_PATH=#{PREFIX}/lib/pkgconfig pkg-config --variable=includedir intro_c-practice || echo .`.chomp
ENV['LD_LIBRARY_PATH'] = "#{ENV['LD_LIBRARY_PATH']}:#{FFI_LIBDIR}"
sh "export LD_LIBRARY_PATH"

ENV['GOPATH'] = `go env GOPATH`.chomp
#ENV['PATH'] = #{ENV['PATH']}:#{ENV['GOPATH']}/bin
ENV['CGO_CFLAGS'] = "#{ENV['CGO_CFLAGS']} " + `PKG_CONFIG_PATH=#{PREFIX}/lib/pkgconfig pkg-config --cflags intro_c-practice`.chomp
ENV['CGO_LDFLAGS'] = "#{ENV['CGO_LDFLAGS']} " + `PKG_CONFIG_PATH=#{PREFIX}/lib/pkgconfig pkg-config --libs intro_c-practice`.chomp

proj_dir = Dir.pwd

file "pkg/classicc/classicc_wrap.c" => "classicc.i" do |t|
  mkdir_p("pkg/classicc")
  sh "swig -go -cgo -intgosize 32 -package classicc -v -I#{FFI_INCDIR} -outdir pkg/classicc -o #{t.name} #{t.source} || true"
end

desc 'Prepare Swig files'
task :prep_swig => ["pkg/classicc/classicc_wrap.c"]
