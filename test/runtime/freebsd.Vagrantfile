# XXX: FreeBSD on Vagrant not working ;(
#
# https://forums.freebsd.org/threads/62311/
# https://github.com/vagrant-libvirt/vagrant-libvirt/issues/50
# https://github.com/freebsd/pkg/issues/1612
#
# Summary: https://gist.github.com/otiai10/b954787954797911974f6e399dfa8c8a

Vagrant.configure("2") do |config|
  config.vm.guest = :freebsd
  config.vm.synced_folder ".", "/vagrant", id: "vagrant-root", disabled: true
  config.vm.box = "freebsd/FreeBSD-10.4-STABLE"
  config.ssh.shell = "sh"
  config.vm.base_mac = "080027D14C66"

  config.vm.provider :virtualbox do |vb|
    vb.customize ["modifyvm", :id, "--memory", "1024"]
    vb.customize ["modifyvm", :id, "--cpus", "1"]
    vb.customize ["modifyvm", :id, "--hwvirtex", "on"]
    vb.customize ["modifyvm", :id, "--audio", "none"]
    vb.customize ["modifyvm", :id, "--nictype1", "virtio"]
    vb.customize ["modifyvm", :id, "--nictype2", "virtio"]

    vb.name = ENV["VIRTUALBOX_NAME"]

  end
  config.vm.provision :shell, :inline => "echo 'TODO: pkg install tesseract/tesseract-dev'"
end
