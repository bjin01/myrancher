FROM opensuse/tumbleweed:latest

RUN mkdir /myapp
WORKDIR /myapp
RUN mkdir -p /usr/share/zoneinfo/Europe/
ADD myrancher /myapp
RUN chmod +x /myapp/myrancher
ADD timezones/Berlin /usr/share/zoneinfo/Europe/
RUN ln -s /usr/share/zoneinfo/Europe/Berlin /etc/localtime

EXPOSE 3000
ENTRYPOINT ["/myapp/myrancher"]