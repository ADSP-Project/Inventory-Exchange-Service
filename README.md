# Federdated-Marketplaces

## Setup that deploys remote images (i.e. no local changes will be reflected):
### Setup the Shops on your local machine OR Google Cloud VM using


    sudo apt update
    
    #install git on google vm
    sudo apt install git-all
    
    #clone our repo
    git clone https://github.com/ADSP-Project/Federdated-Marketplaces.git
    cd Federdated-Marketplaces
    
    #run the setup script, it will set up everything for you
    sudo ./setup.sh
    
    

after that you should have setup both clusters. Depending on your machine you might need up some (long) time until all services are running.
To be able to access the shops from outside the machine with the browser we have to configure the network access. For this go to the VM Firewall settings
and set up a firewall with the configuration as seen in the picture.

![Screenshot from 2023-05-12 12-27-39](https://github.com/ADSP-Project/Federdated-Marketplaces/assets/66095628/e751bc69-730d-4b10-b670-5c7f40d681c5)

Now just find out your external IP Adress in the IP-Adress Tab, or with `curl ifconfig.me.`, and then you can access the shops:

socket-shop YOUR-EXTERN-IP-ADRESS:8080

onlineboutique YOUR-EXTERN-IP-ADRESS:8081

if you are running locally:
  
sock-shop localhost:8080
  
onlineboutique localhost:8081


Get Products of shops

sock-shop: curl localhost:8083/api/getproducts

onlineboutique: localhost:8082/products

## Setup that deploys local images using skaffold (i.e. local changes will be reflected):

### Install dependencies:

For linux Debian x86-64:

Install docker (https://docs.docker.com/desktop/install/linux-install/) (https://docs.docker.com/engine/install/linux-postinstall/)

    # Install docker
    sudo apt-get update
    sudo apt-get install ./docker-desktop-<version>-<arch>.deb  
    sudo groupadd docker
    sudo usermod -aG docker $USER
    # If you’re running Linux in a virtual machine, it may be necessary to restart the virtual machine for changes to take effect.
    newgrp docker
    docker run hello-world

Install minikube (https://minikube.sigs.k8s.io/docs/start/)

    curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
    sudo install minikube-linux-amd64 /usr/local/bin/minikube

Install kubectl (https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/)

    # Install docker
    curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"

    # Validate the binary (optional)
    curl -LO "https://dl.k8s.io/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl.sha256"
    # Validate the kubectl binary against the checksum file
    echo "$(cat kubectl.sha256)  kubectl" | sha256sum --check
    # If valid, the output is:
    # kubectl: OK

Install git and repo:

    sudo apt install git-all
    
    git clone https://github.com/ADSP-Project/Federdated-Marketplaces.git

    git clone https://github.com/ADSP-Project/Federation-Hub.git

Start minikube clusters:

    minikube start --cpus=4 --memory 4096 --disk-size 32g --profile online-boutique
    
    minikube start --memory 8192 --cpus 4 --profile sock-shop

    minikube start --profile hub

    # Verify by listing profiles:
    minikube profile list

    # IMPORTANT!! When you want to change profile you have to write the command 
    kubectl config use-context <profile>

IMPORTANT: Before you run the code you have to make a change on the IP address in the checkoutservice Dockerfile, change the ENV PUBLIC_IP to the public IP address of your virtual machine, otherwise the checkoutservice won't be able to send sock-shop requests. Same for federationservice-ui.yaml, change the VITE_FEDERATION_SERVICE to the public IP address of the virtual machine.

To run the code:
    # cd into Federdated-Marketplaces
    cd Federdated-Marketplaces
    
    # cd into socks-shop
    cd socks-shop
    kubectl config use-context sock-shop
    skaffold run

    #cd into onlineboutique
    cd ../onlineboutique
    kubectl config use-context online-boutique
    skaffold run

    # cd into Federation-Hub
    kubectl config use-context hub
    cd ../../Federation-Hub/
    skaffold run
    

To stop code:

    skaffold delete

Expose ports:

    #Sock-Shop
    #Set config
    kubectl config use-context sock-shop
    kubectl port-forward deployment/front-end 8081:8079 --address 0.0.0.0 &
    kubectl port-forward deployment/nextjs-docker 3000:3000 --address 0.0.0.0 &

    #Online boutique
    #Set config
    kubectl config use-context online-boutique
    kubectl port-forward deployment/frontend --address 0.0.0.0 8080:8080 &
    kubectl port-forward deployment/apiservice --address 0.0.0.0 9090:9090 &
    kubectl port-forward deployment/productcatalogservice --address 0.0.0.0 3560:3560 &
    kubectl port-forward deployment/federationservice-ui --address 0.0.0.0 5173:5173 &
    kubectl port-forward deployment/federationservice-be --address 0.0.0.0 8091:8091 & 

To get logs from pod:

    kubectl config use-context <sock-shop/onlineboutique>
    kubectl logs deployment/<nameofservice>



