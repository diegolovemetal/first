FROM library/golang
# Godep for vendoring
RUN go get github.com/beego/bee
ENV APP_DIR $GOPATH/src/BeegoD
RUN mkdir -p $APP_DIR
# Set the entrypoint
ENTRYPOINT (cd $APP_DIR && bee run > debug_log.log 2>&1)
#ADD . $APP_DIR
EXPOSE 80