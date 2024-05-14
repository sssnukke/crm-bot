FROM node:20-alpine as build

WORKDIR /app

COPY package.json yarn.lock ./

RUN yarn

FROM node:20-alpine as development

ENV TZ='Asia/Vladivostok'

WORKDIR /app

COPY package.json yarn.lock ./

COPY --from=build /app/node_modules ./node_modules

COPY . .

CMD ["yarn", "start:dev"]

FROM node:20-alpine as build-production

ENV TZ='Asia/Vladivostok'

WORKDIR /app

COPY . .

RUN yarn --only=production && yarn cache clean --force

RUN yarn -g @nestjs/cli

RUN yarn build

FROM node:20-alpine as production

COPY --from=build-production /app/node_modules ./node_modules
COPY --from=build-production /app/dist ./dist

CMD ["node", "dist/main.js"]
