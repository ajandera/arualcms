FROM golang:1.17-alpine3.14

RUN apk add git
RUN apk add openssh
RUN git config --global url."git@gitlab.com:".insteadOf "https://gitlab.com/"


# wkhtmltopdf
RUN apk add xvfb wkhtmltopdf

COPY ./entrypoint.sh /
