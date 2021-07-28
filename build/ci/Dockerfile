FROM ubuntu:20.04

WORKDIR /app/national-lottery-picker

ENV TZ=Europe/London

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
RUN apt clean all && apt-get update && apt-get install golang wget xvfb default-jre -y

RUN wget -c https://www.cisco.com/security/pki/certs/ciscoumbrellaroot.pem \
  -O /usr/local/share/ca-certificates/ciscoumbrellaroot.crt \
  && update-ca-certificates

RUN wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb \
  && (dpkg -i google-chrome*.deb || apt-get -f install -y) && rm -f google-chrome*.deb

COPY . .

RUN go build -o bin/national-lottery-picker cmd/main.go && mv bin/* /usr/local/bin/.
RUN mv build/ci/startup.sh /app/. \
  && mkdir /etc/national-lottery-picker \
  && mv build/ci/config.yml /etc/national-lottery-picker/. \
  && rm -rf /app/national-lottery-picker

CMD [ "/usr/local/bin/national-lottery-picker" ]
