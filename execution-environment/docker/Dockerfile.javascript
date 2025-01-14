FROM node:18-slim

WORKDIR /app

# Copy the executor script
COPY executor.js /app

CMD ["node", "executor.js"]
