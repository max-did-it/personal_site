FROM starefossen/ruby-node:2-8

RUN mkdir frontend
WORKDIR /frontend

COPY ./frontend/package.json ./
RUN gem install slim
RUN curl --compressed -o- -L https://yarnpkg.com/install.sh | bash
RUN yarn install
ENTRYPOINT [ "yarn", "build"]