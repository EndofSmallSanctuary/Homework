FROM python:3.6-slim AS complier
WORKDIR /app
COPY Flask /app

RUN apt-get update \
    && apt-get install -y build-essential patchelf \
    && pip install --upgrade pip \
    && pip install -r requirements.txt \
    && pip install staticx pyinstaller \ 
    && pyinstaller -F main.py \
    && staticx /app/dist/main  /test2

from scratch
COPY --from=complier test2 test2
COPY --from=complier tmp /tmp
# COPY /build .
CMD ["/test2"]