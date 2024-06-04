VERSION=0.9.5
arch=$(uname -m)
path=/usr/local/kubectx
filename=kubectx_v${VERSION}_linux_${arch}.tar.gz
url=https://github.com/ahmetb/kubectx/releases/download/v${VERSION}/${filename}
wget $url
if [ ! -d "${path}" ]; then
	sudo mkdir $path
fi
sudo tar -C $path -zxf $filename
rm $filename
echo "# kubectx bin" >> ~/.profile
echo "PATH=\$PATH:${path}" >> ~/.profile
