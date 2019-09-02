FROM golang:1.12.9

MAINTAINER Ken

RUN mkdir -p /go/src/observedcat && mkdir -p /observedcat

COPY ./ /go/src/observedcat

RUN touch /tmp/abc.log && cd /go/src/observedcat && go install \
    && mv config.yaml /config.yaml \
    && rm -rf /go/src/observedcat


#安裝supervisor
RUN apt-get update && \
  apt-get -y install supervisor vim && \
  rm -rf /var/lib/apt/lists/* && \
  mkdir -p /var/log/supervisor && \
  mkdir -p /etc/supervisor/conf.d


COPY ./run.conf /etc/supervisor/conf.d/run.conf
COPY ./supervisord.conf /etc/supervisor/supervisord.conf


CMD ["supervisord", "-c", "/etc/supervisor/supervisord.conf", "-n"]