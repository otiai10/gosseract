Vagrant.configure("2") do |config|
  config.vm.guest = :freebsd
  config.vm.synced_folder "./", "/home/vagrant/app", owner: "vagrant", group: "vagrant", disabled: true
  config.vm.box = "freebsd/FreeBSD-12.2-STABLE"
  config.ssh.shell = "sh"
  config.vm.base_mac = "080027D14C66"

  config.vm.provider :virtualbox do |vb|
    vb.name = ENV["VIRTUALBOX_NAME"]
  end

  config.vm.provision :shell, :inline => '
    mkdir -p $GOPATH/src/github.com/otiai10
    cp -r /vagrant $GOPATH/src/github.com/otiai10/gosseract
    pkg install -y --quiet tesseract git go
    mv /usr/local/share/tessdata/*.traineddata /tmp
    mv /tmp/eng.traineddata /usr/local/share/tessdata/
    cd $GOPATH/src/github.com/otiai10/gosseract
    go get -t -v ./...
    go test -v -cover github.com/otiai10/gosseract
  ', :env => {"GOPATH" => "/home/vagrant/go"}
end
