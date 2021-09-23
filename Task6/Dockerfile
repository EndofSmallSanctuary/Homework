FROM alpine:3.14 as base
FROM base as builder


RUN mkdir /install
WORKDIR /install



COPY /requirements.txt /requirements.txt
RUN apk add --update --no-cache python3 && ln -sf python3 /usr/bin/python  && python3 -m ensurepip && python3 -m pip install --upgrade pip  &&  pip3 install -r /requirements.txt

FROM base


COPY --from=builder /install /usr/local
RUN apk add --update --no-cache python3
COPY . /Flask
WORKDIR /Flask
ENTRYPOINT [ "python3"]
CMD ["main.py"]
