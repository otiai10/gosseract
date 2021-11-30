Vagrant.configure("2") do |config|
  config.vm.guest = :freebsd
  config.vm.synced_folder "./", "/home/vagrant/app", owner: "vagrant", group: "vagrant", disabled: true
  config.vm.box = "freebsd/FreeBSD-12.2-STABLE"
  config.ssh.shell = "sh"
  config.vm.base_mac = "080027D14C66"

  config.vm.provider :virtualbox do |vb|
    vb.name = ENV["VIRTUALBOX_NAME"]
  end

  # User: vagrant
  # Pass: vagrant
  config.vm.provision :shell, :inline => '
    mkdir -p $GOPATH/src/github.com/otiai10
    cp -r /vagrant $GOPATH/src/github.com/otiai10/gosseract
    pkg install -y --quiet tesseract tesseract-data git go
    mv /usr/local/share/tessdata/*.traineddata /tmp
    mv /tmp/eng.traineddata /usr/local/share/tessdata/
    cd $GOPATH/src/github.com/otiai10/gosseract
    go test -v -cover ./...
    echo $? > /vagrant/test/runtimes/TESTRESULT.freebsd.txt
  ', :env => {
    "GOPATH" => "/home/vagrant/go",
    "TESSDATA_PREFIX" => "/usr/local/share/tessdata",
  }
end
