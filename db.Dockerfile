# start with base image
FROM mysql:8.0.23

# import data into container
# All scripts in docker-entrypoint-initdb.d/ are automatically executed during container startup
COPY ./Init_SQL_Query.sql /docker-entrypoint-initdb.d/

