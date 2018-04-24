FROM golang

ARG app_env
ENV APP_ENV $app_env

COPY ./src /go/src/app
WORKDIR /go/src/app

RUN go get ./
RUN go build

CMD app
	
#CMD if [ ${APP_ENV} == "production" ]; \
#	then \
#	app; \
#	else \
#	go get github.com/pilu/fresh && \
#	fresh; \
#	fi
EXPOSE 8080