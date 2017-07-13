#Using same OS as my ssvm for my image inside the container
FROM centos:centos7
MAINTAINER Mark Magaling magaling.markizm@gmail.com 
RUN yum -y update
RUN yum -y install mailx-12.5-12.el7_0.x86_64 && yum -y clean all
COPY . /
COPY ./run /
RUN chmod -v +x /run
WORKDIR /
EXPOSE 8085

CMD ["./run"]
