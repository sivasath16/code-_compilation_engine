FROM openjdk:17-jdk-slim

WORKDIR /app

# Copy the executor script
COPY executor.java /app

# Compile the executor
RUN javac executor.java

CMD ["java", "executor"]
