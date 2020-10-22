#!bin/sh
sudo docker stack rm myStack
sudo docker build . -t kadlab
sudo docker stack deploy --compose-file docker-compose.yml myStack
