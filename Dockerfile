FROM scratch
ADD server /
ENTRYPOINT ["/server"]
EXPOSE 8080
