FROM cockroachdb/cockroach

COPY init.sql init.sql

EXPOSE 26257:26257
EXPOSE 8080:8080
VOLUME cockroach-data /cockroach/cockroach-data

CMD ["start", "--insecure"]
