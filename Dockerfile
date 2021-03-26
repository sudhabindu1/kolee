FROM golang

WORKDIR /kolee

ADD templates/index.tmpl templates/index.tmpl

ADD bin/kolee kolee

ENTRYPOINT ["./kolee"]