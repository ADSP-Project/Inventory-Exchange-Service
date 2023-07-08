#!/bin/bash

# Install Docker

sudo apt-get update
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh

# Install k3s
curl -sfL https://get.k3s.io | sh -

# Install k3d
curl -s https://raw.githubusercontent.com/k3d-io/k3d/main/install.sh | bash

# Install kubectl
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl

# Create cluster & deploy sock shop
k3d cluster create sock-shop -p 8080:30001@agent:0 -p 8083:30000@agent:0 --agents 1
kubectl create namespace sock-shop
kubectl apply -f socks-shop/deploy/kubernetes/complete-demo.yaml

# Create cluster & deploy onlineboutique
k3d cluster create onlineboutique -p 8081:32541@agent:0 -p 8082:32540@agent:0 -p 8084:32545@agent:0 --agents 1
k3d kubeconfig merge onlineboutique --kubeconfig-merge-default
kubectl config use-context k3d-onlineboutique
kubectl apply -f ./onlineboutique/release/kubernetes-manifests.yaml

#kubectl config get-contexts

