FROM ubuntu:20.04

WORKDIR /app/national-lottery-picker

ENV TZ=Europe/London

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
RUN apt clean all && apt-get update && apt-get install golang git wget vim xvfb default-jre -y

RUN wget -c https://www.cisco.com/security/pki/certs/ciscoumbrellaroot.pem \
  -O /usr/local/share/ca-certificates/ciscoumbrellaroot.crt \
  && update-ca-certificates

RUN wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb \
  && (dpkg -i google-chrome*.deb || apt-get -f install -y) && rm -f google-chrome*.deb

COPY . .

ENTRYPOINT [ "./test/startup.sh" ]
