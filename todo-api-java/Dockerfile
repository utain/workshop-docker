FROM maven:3.6.1-jdk-8-alpine
RUN mkdir /app/
ADD . /app/
WORKDIR /app/
RUN mvn clean install

FROM openjdk:8-jdk-alpine

RUN mkdir /oneweb/
RUN mkdir /app/
# ONEWEB_HOME
ENV ONEWEB_HOME=/oneweb/

# Make port 8080 available to the world outside container
EXPOSE 8080

WORKDIR /app/
# Add the application's JAR file to the container
COPY --from=0 /app/target/todo-api-0.0.1-SNAPSHOT.war /app/app.jar

# Run the JAR file
ENTRYPOINT ["java", "-Djava.security.egd=file:/dev/./urandom", "-jar", "/app/app.jar"]