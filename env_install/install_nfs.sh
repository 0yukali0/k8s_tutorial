SHARE_PATH=/var/nfsshare
EXPORT_PATH=/etc/exports
sudo apt update && sudo apt install nfs-kernel-server nfs-common
sudo mkdir -p $SHARE_PATH
sudo chown nobody:nogroup $SHARE_PATH
sudo bash -c "echo '' > ${EXPORT_PATH}"
sudo bash -c "echo '${SHARE_PATH} *(rw,sync,no_root_squash)' > ${EXPORT_PATH}"
sudo exportfs -arv
sudo systemctl enable nfs-kernel-server
sudo systemctl restart nfs-kernel-server
