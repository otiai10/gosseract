Vagrant.configure("2") do |config|
  config.vm.guest = :freebsd
  config.vm.synced_folder ".", "/vagrant", id: "vagrant-root", disabled: true
  config.vm.box = "freebsd/FreeBSD-10.4-STABLE"
  config.ssh.shell = "sh"
  config.vm.base_mac = "080027D14C66"

  config.vm.provider :virtualbox do |vb|
    vb.name = ENV["VIRTUALBOX_NAME"]
  end

  config.vm.provision :shell, :inline => '
    pkg install -y --quiet tesseract git go
    export GOPATH=~/go
    go get -t github.com/otiai10/gosseract
  '
  config.vm.provision :shell, :inline => "go test github.com/otiai10/gosseract"
end
