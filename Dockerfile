FROM golang:1.14

RUN mkdir -p /work/markdown-toc
WORKDIR /work/markdown-toc
ADD . /work/markdown-toc

RUN make install

ENTRYPOINT ["markdown-toc"]
