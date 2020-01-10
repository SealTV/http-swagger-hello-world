FROM scratch

COPY ./http-swagger-hw .

EXPOSE 8080

ENTRYPOINT [ "http-swagger-hw" ]