FROM alpine:latest

RUN mkdir /app

COPY surveyApp /app

CMD ["/app/surveyApp"]