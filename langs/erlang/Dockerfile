FROM erlang:25.1.2.0-alpine

FROM alpine:3.16

COPY --from=0 /bin/echo /bin/cat /bin/sh   /bin/
COPY --from=0 /usr/local/lib/erlang        /usr/local/lib/erlang
COPY --from=0 /usr/lib/libncursesw.so.6    \
              /usr/lib/libgcc_s.so.1       \
              /usr/lib/libstdc++.so.6      /usr/lib/

COPY erlang /usr/bin

ENTRYPOINT ["erlang"]

CMD ["--version"]
