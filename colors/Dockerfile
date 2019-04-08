FROM gcr.io/cloud-builders/go
WORKDIR /go/src/github.com/priyawadhwa/next-2019-demo/colors
COPY colors/ .
RUN GOOS=linux CGO_ENABLED=0 go build -o /colors main.go

FROM scratch
COPY --from=0 /colors /colors
CMD ["/colors"]
EXPOSE 8888
