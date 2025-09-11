# Docker

#docker

## Start a local docker quick

```sh
sudo docker run --name studypg -p 5000:5432 -e POSTGRES_USER=sylveryte  -e POSTGRES_PASSWORD=study -d postgres:16
```

## List and Delete

- list containers

```sh
sudo docker container ls -a #list all
sudo docker container rm 78 48 39 #delete container
sudo docker container prune #rm all stopped
```

## Run Dockerfile

### Build

    `docker build --tag 'image_name'` .

---

---

link: [docker-for-web-developers](https://www.mattbutton.com/docker-for-web-developers/)
