KIND=kind
if command -v $KIND &> /dev/null; then
	echo "${KIND} is already installed"
	exit 0
fi

GO=go
KIND_VERSION=v0.23.0
if command -v $GO &> /dev/null; then
	echo "kind installation by go"
        $GO install "sigs.k8s.io/kind@${KIND_VERSION}"
        GOPATH=$(go env GOPATH)
        "PATH=\$PATH:${GOPATH}/bin" >> ~/.profile
        export "PATH=\$PATH:${GOPATH}/bin"
else
	echo "kind installation via shell"
        # For AMD64 / x86_64
        [ $(uname -m) = x86_64 ] && curl -Lo ./kind "https://kind.sigs.k8s.io/dl/${KIND_VERSION}/kind-linux-amd64"
        # For ARM64
        [ $(uname -m) = aarch64 ] && curl -Lo ./kind "https://kind.sigs.k8s.io/dl/${KIND_VERSION}/kind-linux-arm64"
        chmod +x ./kind
        sudo mv ./kind /usr/local/bin/kind
fi
