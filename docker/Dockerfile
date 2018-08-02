# Build binary
FROM golang:1.9.4 as builder
LABEL ken <ken.lin@infinitechain.io>
ARG ETH_AGENT_GO_SRC_CT
COPY . $ETH_AGENT_GO_SRC_CT
RUN cd $ETH_AGENT_GO_SRC_CT && \
    go get -u github.com/kardianos/govendor && \
    govendor sync && \
    cd $ETH_AGENT_GO_SRC_CT/ && \
    go build -o eth-agent ./main.go


# Build eth-agent image
FROM ubuntu:16.04
LABEL ken <ken.lin@infinitechain.io>
ARG ETH_AGENT_GO_SRC_CT
ARG WORKDIR_CT
ENV USER=eth-agent
ENV HOME=/home/$USER
WORKDIR $WORKDIR_CT
RUN apt-get update && apt-get install -y \
    ca-certificates && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*
     
# Setting eth-agent
COPY --from=builder $ETH_AGENT_GO_SRC_CT/eth-agent /usr/bin/eth-agent
COPY --from=builder $ETH_AGENT_GO_SRC_CT/sysconfig.yaml.tpl $WORKDIR_CT/sysconfig.yaml.tpl
COPY ./docker/.env $WORKDIR_CT/.env
COPY ./docker/script/eth-agent-entrypoint.sh /docker-entrypoint.sh
RUN ["chmod","+x","/docker-entrypoint.sh"]
#Setting user
#RUN useradd -ms /bin/bash $USER \
#    && mkdir -p $HOME
#RUN chown -R $USER /etc/eth-agent
#USER $USER
#WORKDIR $HOME
ENTRYPOINT ["/docker-entrypoint.sh"]
# CMD ["-e", "p"]
# ENTRYPOINT ["/usr/bin/eth-agent","--config","/usr/bin/sysconfig.yaml"]