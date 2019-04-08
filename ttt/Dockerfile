FROM node
WORKDIR /app
COPY ttt/package*.json /app/
RUN npm install
COPY ttt/ /app
RUN npm run build

FROM golang
COPY ttt/server /go/src/github.com/priyawadhwa/next-2019-demo/ttt/server
RUN go build -o /server github.com/priyawadhwa/next-2019-demo/ttt/server

FROM debian
COPY --from=0 /app/dist /dist/
COPY --from=1 /server /
COPY ttt/index.html /dist/
CMD ["/server"]
