
# syntax=docker/dockerfile:1
FROM node:20 as build

WORKDIR /app
COPY frontend-react/package.json frontend-react/package-lock.json* frontend-react/yarn.lock* ./
RUN npm install
COPY frontend-react/ .
RUN npm run build

COPY --from=build /app/dist /usr/share/nginx/html

EXPOSE 80


