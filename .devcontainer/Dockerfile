ARG GO_VERSION=1.21

# Used to install CRDB into the devcontainer
FROM cockroachdb/cockroach:latest-v22.2 as CRDB

FROM mcr.microsoft.com/vscode/devcontainers/go:1-${GO_VERSION}-bullseye

# Set up crdb
RUN mkdir /usr/local/lib/cockroach
COPY --from=CRDB /cockroach/cockroach /usr/local/bin
COPY --from=CRDB /usr/local/lib/cockroach/libgeos.so /usr/local/lib/cockroach/
COPY --from=CRDB /usr/local/lib/cockroach/libgeos_c.so /usr/local/lib/cockroach/

# [Optional] Uncomment this section to install additional OS packages.
RUN export DEBIAN_FRONTEND=noninteractive \
    && curl https://baltocdn.com/helm/signing.asc | gpg --dearmor | sudo tee /usr/share/keyrings/helm.gpg \
    &&  curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | gpg --dearmor > /usr/share/keyrings/yarn-archive-keyring.gpg \
    && apt-get install apt-transport-https --yes \
    && echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/helm.gpg] https://baltocdn.com/helm/stable/debian/ all main" | sudo tee /etc/apt/sources.list.d/helm-stable-debian.list \
    && apt-get update \
    && apt-get -y install --no-install-recommends \
    bash-completion \
    helm \
    uuid-runtime \
    postgresql-client

# Install NATS Tooling
RUN curl -o /tmp/install.sh  https://raw.githubusercontent.com/nats-io/nsc/main/install.sh \
    && chmod +x /tmp/install.sh \
    && /tmp/install.sh -d /usr/local/bin -s - \
    && rm -f /tmp/install.sh

USER vscode
ENV NATS_CLI_VERSION=0.0.35
RUN go install -v github.com/cweill/gotests/gotests@v1.6.0 \
    && go install github.com/nats-io/natscli/nats@v${NATS_CLI_VERSION} \
    && go install github.com/nats-io/nkeys/nk@latest

RUN go install ariga.io/atlas/cmd/atlas@latest

USER root
