FROM gcc:12

WORKDIR /app

# Copy the executor script
COPY executor.cpp /app

# Compile the executor script
RUN g++ executor.cpp -o executor

CMD ["./executor"]
