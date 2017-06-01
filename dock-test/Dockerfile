FROM chug/ubuntu14.04x64 
MAINTAINER ASAD MEMON, OSMAN ALI

# Update the repository sources list
RUN echo "deb http://archive.ubuntu.com/ubuntu trusty main universe" > /etc/apt/sources.list
RUN apt-get update
#RUN apt-get upgrade
#Install all the languages/compilers we are supporting.
RUN apt-get install -y gcc
RUN apt-get install -y g++
