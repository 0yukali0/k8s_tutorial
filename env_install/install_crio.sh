CRIO=crio
if command -v $CRIO &> /dev/null; then
	echo "CRI-O is already installed"
	exit 0
fi
OS=xUbuntu_22.04
VERSION=1.24

sudo bash -c "echo 'deb https://download.opensuse.org/repositories/devel:/kubic:/libcontainers:/stable/${OS}/ /' > /etc/apt/sources.list.d/devel:kubic:libcontainers:stable.list"
sudo bash -c "echo 'deb http://download.opensuse.org/repositories/devel:/kubic:/libcontainers:/stable:/cri-o:/${VERSION}/${OS}/ /' > /etc/apt/sources.list.d/devel:kubic:libcontainers:stable:cri-o:${VERSION}.list"

sudo bash -c "curl -L https://download.opensuse.org/repositories/devel:/kubic:/libcontainers:/stable:/cri-o:/${VERSION}/${OS}/Release.key | apt-key add -"
sudo bash -c "curl -L https://download.opensuse.org/repositories/devel:/kubic:/libcontainers:/stable/${OS}/Release.key | apt-key add -"


sudo apt-get update && sudo apt-get install cri-o cri-o-runc -y
sudo systemctl enable crio
sudo systemctl start crio
