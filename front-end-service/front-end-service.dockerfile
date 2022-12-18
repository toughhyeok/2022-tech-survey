FROM alpine:latest

RUN mkdir /app /template

COPY frontApp /app
COPY template /template

CMD [ "/app/frontApp" ]