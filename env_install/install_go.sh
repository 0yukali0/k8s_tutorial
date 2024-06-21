VERSION=1.22.3
GOPATH=/usr/local
GOBIN=$GOPATH/go/bin
GO=go
if ! command -v go &> /dev/null; then
	filename=go$VERSION.linux-amd64.tar.gz
	if ! ls | grep -q "^go$VERSION.*\.tar\.gz"; then
		echo "$filename doesn't exist then download"
		url=https://go.dev/dl/$filename
		wget $url
	fi
	sudo tar -C $GOPATH -xzf $filename
	export PATH=$PATH:$GOBIN
	echo update .profile
	if ! cat ~/.profile |grep "PATH=\$PATH:$GOBIN"; then
		echo "PATH=\$PATH:$GOBIN" >> ~/.profile
	fi
	source ~/.profile
	echo remove $filename
	if ls | grep -q "$filename"; then
		rm $filename
	fi
	echo $GO installation success!
else
  echo $GO is installed
fi
