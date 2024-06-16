ARGOCLI=argocd
if command -v $ARGOCLI &> /dev/null; then
	echo "argocd is already installed"
	exit 0
fi

BIN="/usr/local/bin"
ARGO_DEB=argocd-linux-amd64

VERSION=$(curl -L -s https://raw.githubusercontent.com/argoproj/argo-cd/stable/VERSION)
curl -sSL -o $ARGO_DEB https://github.com/argoproj/argo-cd/releases/download/v$VERSION/argocd-linux-amd64
sudo install -m 555 $ARGO_DEB "${BIN}/${ARGOCLI}"
rm $ARGO_DEB
