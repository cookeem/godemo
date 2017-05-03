# 在jenkins下测试：
# cd /var/jenkins_home/tools/org.jenkinsci.plugins.docker.commons.tools.DockerTool/docker_1.13.1/bin
# ./docker -H tcp://docker:2375 images
# REPOSITORY             TAG                 IMAGE ID            CREATED              SIZE
# registry:5000/godemo   0.1.0               6b9c973d9e80        About a minute ago   13.7 MB
# registry:5000/godemo   latest              6b9c973d9e80        About a minute ago   13.7 MB
# alpine                 latest              4a415e366388        2 months ago         3.98 MB
# registry:5000/alpine   latest              4a415e366388        2 months ago         3.98 MB



FROM alpine
MAINTAINER cookeem cookeem@qq.com

COPY gin_demo /gin_demo

# Commands when creating a new container
CMD /gin_demo

