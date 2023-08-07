FROM alpine
ADD route_api /route_api

ENTRYPOINT [ "/route_api" ]