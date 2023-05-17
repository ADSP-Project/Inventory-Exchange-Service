# Federdated-Marketplaces

### Setup the Shops on your local machine OR Google Cloud VM


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
  
socket-shop localhost:8080
  
onlineboutique localhost:8081
