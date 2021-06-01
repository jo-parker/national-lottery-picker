FROM ubuntu:20.04

WORKDIR /app/euromillions-picker

ENV TZ=Europe/London

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
RUN apt clean all && apt-get update && apt-get install golang git firefox wget vim xvfb default-jre -y

RUN wget -c https://www.cisco.com/security/pki/certs/ciscoumbrellaroot.pem \
  -O /usr/local/share/ca-certificates/ciscoumbrellaroot.crt \
  && update-ca-certificates

COPY . .

ENTRYPOINT [ "./build/ci/startup.sh" ]