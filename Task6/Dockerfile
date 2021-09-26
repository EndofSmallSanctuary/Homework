FROM python:3.6-slim AS bride
WORKDIR /flask
COPY Flask /flask

RUN apt-get update -y && apt-get autoremove -y \
    && apt-get install -y build-essential patchelf \
    && pip install --no-cache-dir --upgrade pip \
    && pip install --no-cache-dir -r requirements.txt \
    && pip install --no-cache-dir staticx pyinstaller \ 
    && pyinstaller -F main.py \
    && staticx /flask/dist/main  /task6

FROM alpine:3.5 as follower
RUN apk add --no-cache libgcc

FROM alpine:3.5 as final
COPY --from=follower  /usr/lib/libgcc_s.so.1 /usr/lib
COPY --from=bride /flask/static /static
COPY --from=bride /flask/templates /templates
COPY --from=bride task6 .
CMD ["/task6"]