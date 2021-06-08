FROM ubuntu:20.04

WORKDIR /app/national-lottery-picker

ENV TZ=Europe/London

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
RUN apt clean all && apt-get update && apt-get install golang git wget vim xvfb default-jre -y

RUN wget -c https://www.cisco.com/security/pki/certs/ciscoumbrellaroot.pem \
  -O /usr/local/share/ca-certificates/ciscoumbrellaroot.crt \
  && update-ca-certificates

RUN apt-get install fonts-liberation libcairo2 libgbm1 libgtk-3-0 libpango-1.0-0 libxkbcommon0 xdg-utils -y
RUN wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb \
  && dpkg -i google-chrome*.deb && rm -f google-chrome*.deb

COPY . .

RUN go test ./... -v -count=1
RUN go build -o bin cmd

ENTRYPOINT [ "./build/ci/startup.sh" ]
