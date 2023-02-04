FROM golang:1.16.0-alpine3.13

ADD ./srcs /ft_auth_bot/
WORKDIR /ft_auth_bot

RUN go install .

CMD ft_auth_bot