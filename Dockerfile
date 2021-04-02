FROM opensuse/tumbleweed:latest

RUN mkdir /myapp
WORKDIR /myapp

# Define your additional repositories here
#RUN zypper ar http://download.opensuse.org/repositories/openSUSE:Tools/openSUSE_15.1 openSUSE:Tools

# Put additional files into container
ADD myrancher /myapp
RUN chmod +x /myapp/myrancher
#RUN zypper install -y osc
EXPOSE 3000
ENTRYPOINT ["/myapp/myrancher"]