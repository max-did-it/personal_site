FROM golang
ADD ./backend /go/src/gitlab.com/SiberianPanda/selfcard_service/
RUN go get gopkg.in/gomail.v2
RUN go install gitlab.com/SiberianPanda/selfcard_service/
ENTRYPOINT /go/bin/selfcard_service
EXPOSE 8080