FROM "golang:1.21-alpine"

# Install dev dependencies
RUN apk add --update \
    inotify-tools bash zsh sudo git zsh-vcs \
    openssh htop \
    make build-base \
    nodejs npm \
    docker-cli postgresql-client curl

# Create and switch to a jailed admin user
RUN echo "%sudo ALL=(ALL:ALL) NOPASSWD: ALL" > /etc/sudoers.d/sudo && \
    addgroup sudo && addgroup docker && \
    adduser -D -s /bin/zsh coder && \
    addgroup coder sudo && \
    addgroup coder docker
USER coder

ENV EDITOR="code --wait"

RUN git config --global core.editor "$EDITOR" && \
    git config --global init.defaultBranch main

# Install oh my zsh
ARG ZSH_THEME_NAME="agnoster"
RUN sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)" && \
    sed -i -e "s/ZSH_THEME=.*/ZSH_THEME=\"$ZSH_THEME_NAME\"/" ~/.zshrc
