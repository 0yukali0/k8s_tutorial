HELM=helm
if command -v $HELM &> /dev/null; then
	echo "${HELM} is already installed"
	exit 0
fi

script=get_helm.sh
curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3
chmod 700 $script
./$script
rm $script
