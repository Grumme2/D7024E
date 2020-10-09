FROM golang:1.15-alpine

RUN apk update && apk upgrade && \
	apk add go git && \
	git clone https://github.com/Grumme2/D7024E.git && \
	cd D7024E && \
	git checkout m1_ping

# Add the commands needed to put your compiled go binary in the container and
# run it when the container starts.
#
# See https://docs.docker.com/engine/reference/builder/ for a reference of all
# the commands you can use in this file.
#
# In order to use this file together with the docker-compose.yml file in the
# same directory, you need to ensure the image you build gets the name
# "kadlab", which you do by using the following command:
#
# $ docker build . -t kadlab
