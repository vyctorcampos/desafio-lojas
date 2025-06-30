FROM node:20 AS builder

WORKDIR /app

COPY package.json package-lock.json ./
RUN npm install --frozen-lockfile

COPY . .

RUN ls -la

RUN npm run build

RUN ls -la .output && \
    ls -la .output/server && \
    ls -la .output/public

FROM node:20-slim

WORKDIR /app

COPY --from=builder /app/.output /app/.output
COPY --from=builder /app/node_modules /app/node_modules

ENV NODE_ENV=production
ENV NITRO_HOST=0.0.0.0
ENV NITRO_PORT=3000
ENV HOST=0.0.0.0
ENV PORT=3000

HEALTHCHECK --interval=30s --timeout=3s \
  CMD curl -f http://localhost:3000 || exit 1

EXPOSE 3000

CMD ["node", "/app/.output/server/index.mjs"]