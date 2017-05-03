FROM alpine
MAINTAINER cookeem cookeem@qq.com

COPY gin_demo /gin_demo

# Commands when creating a new container
CMD /gin_demo
