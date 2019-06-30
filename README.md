# SERVICE_DEMO

## คำอธิบาย

เป็นโปรเจคสร้าง REST API ด้วย GO โดยออกแบบ Structure โปรเจคให้คล้ายกับ [ardanlabs/servic](https://github.com/ardanlabs/service) โปรเจคนี้เป็นการทำงาน CRUD โดยใช้ Database เป็น PostgreSQL

## สิ่งที่ต้องติดตั้ง

- [Golang version 1.10 ขึ้นไป](https://golang.org/dl/)
- [Docker](https://docs.docker.com/install/)

## รันโปรเจค

ในโปรเจคนี้มี dockerfile สำหรับ build application และ docker-compose สำหรับ run service

เรามี makefile เพื่อให้ง่ายต่อการทำงาน โดยมีคำสั่ง build-app, up, down

### Build Application

build application ให้เป็น docer image
``` 
$ make build-app
```

### Run Application
สั่ง docker-compose up เพิ่ม start PostgreSQL, adminer และ sales-api
```
$ make up
```

### Stop Applicatio
```
$ <ctrl>C
$ make down
```