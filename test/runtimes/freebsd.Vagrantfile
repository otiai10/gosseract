Vagrant.configure("2") do |config|
  # https://developer.hashicorp.com/vagrant/docs/synced-folders/basic_usage#disabling
  config.vm.synced_folder ".", "/vagrant", disabled: true

  config.vm.guest = :freebsd
  config.vm.box = "freebsd/FreeBSD-12.2-STABLE"
  config.vm.base_mac = "080027D14C66"
  config.ssh.shell = "sh"
  # config.ssh.username = "vagrant"
  # config.ssh.password = "vagrant"

  config.vm.provider :virtualbox do |vb|
    vb.name = ENV["VIRTUALBOX_NAME"]
  end

  # {{{ FIXME: This is ugly.
  config.vm.provision :shell, :inline => 'rm -rf /home/vagrant/go/src/github.com/otiai10/gosseract'
  config.vm.provision :file, source: "./", destination: "/home/vagrant/go/src/github.com/otiai10/gosseract"
  # }}}

  config.vm.provision :shell, :inline => '
    pkg install -y tesseract tesseract-data git go
    cd $GOPATH/src/github.com/otiai10/gosseract
    go test -v -cover ./...
    exit $?
  ', :env => {
    "GOPATH" => "/home/vagrant/go",
    "TESSDATA_PREFIX" => "/usr/local/share/tessdata",
    "CPATH" => "/usr/local/include",
    # "LIBRARY_PATH" => "/usr/local/lib",
    # "LD_LIBRARY_PATH" => "/usr/local/lib",
  }
end
