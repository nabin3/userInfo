# Importing os
FROM debian:stable-slim

# Copy source destination
COPY service/bin/userservice /bin/userservice 
COPY api_server/bin/API_SERVER /bin/API_SERVER

# Run userservice and API_SERVER at image startup
CMD ["/bin/userservice", "/bin/API_SERVER"]

EXPOSE 3000