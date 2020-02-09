# Build Stage
FROM lacion/alpine-golang-buildimage:1.13 AS build-stage

LABEL app="build-githubinfo"
LABEL REPO="https://github.com/zloeber/githubinfo"

ENV PROJPATH=/go/src/github.com/zloeber/githubinfo

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /go/src/github.com/zloeber/githubinfo
WORKDIR /go/src/github.com/zloeber/githubinfo

RUN make build-alpine

# Final Stage
FROM lacion/alpine-base-image:latest

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/zloeber/githubinfo"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/githubinfo/bin

WORKDIR /opt/githubinfo/bin

COPY --from=build-stage /go/src/github.com/zloeber/githubinfo/bin/githubinfo /opt/githubinfo/bin/
RUN chmod +x /opt/githubinfo/bin/githubinfo

# Create appuser
RUN adduser -D -g '' githubinfo
USER githubinfo

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/opt/githubinfo/bin/githubinfo"]
