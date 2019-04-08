FROM node
WORKDIR /app
COPY frontend/package*.json /app/
RUN npm install
COPY frontend/ /app
RUN npm run build

FROM golang
COPY frontend/server /go/src/github.com/dlorenc/gameoflife/server
RUN go build -o /server github.com/dlorenc/gameoflife/server

FROM debian
COPY --from=0 /app/dist /dist/
COPY --from=1 /server /
COPY frontend/index.html /dist/
COPY frontend/favicon.png /dist/
CMD ["/server"]
