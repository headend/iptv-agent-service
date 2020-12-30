FROM centos:7.2.1511

RUN mkdir /app
WORKDIR /app
ADD agent-service /app/agent-service
RUN rm -f etc/localtime; ln -sf /usr/share/zoneinfo/Asia/Ho_Chi_Minh /etc/localtime
EXPOSE 50006
CMD ["/app/agent-service"]