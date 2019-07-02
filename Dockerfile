FROM heroku/heroku:18-build as build

COPY . /app
WORKDIR /app

# Setup buildpack
RUN mkdir -p /tmp/buildpack/heroku/go /tmp/build_cache /tmp/env
RUN curl https://codon-buildpacks.s3.amazonaws.com/buildpacks/heroku/go.tgz | tar xz -C /tmp/buildpack/heroku/go

# Install dependencies
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go get -u github.com/tools/godep
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go get -u github.com/pborman/uuid
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go get -u github.com/gin-gonic/gin
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go get -u github.com/heroku/x/hmetrics/onload
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go get -u github.com/lib/pq
RUN CGO_ENABLED=0 GOOS=linxu GOARCH=amd64 godep save ./...


#Execute Buildpack
RUN STACK=heroku-18 /tmp/buildpack/heroku/go/bin/compile /app /tmp/build_cache /tmp/env

# Prepare final, minimal image
FROM heroku/heroku:18

COPY --from=build /app /app
ENV HOME /app
WORKDIR /app
RUN useradd -m heroku
USER heroku
CMD /app/bin/personalassistant
