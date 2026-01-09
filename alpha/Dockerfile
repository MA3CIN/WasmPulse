FROM alpine:latest

LABEL description="WasmPulse platform (version alpha)"

COPY measure.sh /usr/local/bin/measure.sh
RUN chmod +x /usr/local/bin/measure.sh

ENTRYPOINT ["/usr/local/bin/measure.sh"]
