FROM debian:stable-slim

WORKDIR /jfrog-cli

ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update && apt-get install -y curl && rm -rf /var/lib/apt/lists/* 

RUN curl -fL https://getcli.jfrog.io | sh

ENV JFROG_CLI_OFFER_CONFIG false
ENV BINTRAY_LICENCES MIT

RUN /jfrog-cli/jfrog bt config --licenses $BINTRAY_LICENCES

RUN ln -s /jfrog-cli/jfrog /usr/local/bin/jfrog

CMD ["/jfrog-cli/jfrog"]

