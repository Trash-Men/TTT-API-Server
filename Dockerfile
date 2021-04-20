# docker build -t woochanlee/trash-men-server:1 .
# docker run -it -P trash-men-server:1
# docker push woochanlee/trash-men-server:1     
# docker run -it -p 5000:5000 ttt-server:1.0.0 
# docker run -it -p 5001:5000 ttt-server:1.0.0 <- <로컬 포트>:<컨테이너 포트>
# docker run -it -P ttt-server:1.0.0 <- 컨테이너 포트 5000(expose한것)이 로컬 컴퓨터 무작위 포트에 매핑됨


### Builder
FROM golang:1.16.2-alpine3.13 as builder
# RUN apk update && apk add git && apk add ca-certificates
# https://lynlab.co.kr/blog/dock89

ARG MODE
ARG DB_NAME
ARG DB_HOST
ARG DB_PORT
ARG DB_USER
ARG DB_PASSWORD
ARG SERVER_PORT
ARG S3_BUCKET_NAME
ARG JWT_SECRET_KEY
ARG IAM_ACCESS_KEY
ARG IAM_SECRET_ACCESS_KEY

WORKDIR /usr/src/app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 MODE=production go build -ldflags "-s -X 'github.com/Trash-Men/api-server/configs.jwtSecretKey=${JWT_SECRET_KEY}' -X 'github.com/Trash-Men/api-server/configs.s3BucketName=${S3_BUCKET_NAME}' -X 'github.com/Trash-Men/api-server/configs.mode=${MODE}' -X 'github.com/Trash-Men/api-server/configs.dbName=${DB_NAME}' -X 'github.com/Trash-Men/api-server/configs.dbHost=${DB_HOST}' -X 'github.com/Trash-Men/api-server/configs.dbPort=${DB_PORT}' -X 'github.com/Trash-Men/api-server/configs.dbUser=${DB_USER}' -X 'github.com/Trash-Men/api-server/configs.dbPassword=${DB_PASSWORD}' -X 'github.com/Trash-Men/api-server/configs.serverPort=${SERVER_PORT}' -X 'github.com/Trash-Men/api-server/configs.iamAccessKey=${IAM_ACCESS_KEY}' -X 'github.com/Trash-Men/api-server/configs.iamSecretAccessKey=${IAM_SECRET_ACCESS_KEY}'" -o main .

# https://lynlab.co.kr/blog/64

### Make executable image
FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /usr/src/app/main /main

EXPOSE 5000

ENTRYPOINT [ "/main" ]
