#!bin/sh
sudo apt update
sudo apt install docker.io -y
sudo docker build . -t kadlab
sudo docker swarm init
sudo docker stack deploy --compose-file docker-compose.yml myStack
