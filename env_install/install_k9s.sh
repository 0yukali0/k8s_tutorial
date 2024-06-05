filename=k9s_Linux_amd64.tar.gz
path=/usr/local/k9s
wget https://github.com/derailed/k9s/releases/download/v0.32.4/$filename
sudo mkdir $path
sudo tar -C $path -zxf $filename
rm $filename
echo "PATH=\$PATH:$path" >> ~/.profile
