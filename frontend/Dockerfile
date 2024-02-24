FROM node:21-alpine

ARG VITE_API_URL
ENV VITE_API_URL $VITE_API_URL

# Set the working directory in the container
WORKDIR /app/frontend

# Copy package.json and package-lock.json to the container
COPY package*.json ./

# Install app dependencies
RUN npm install

# Copy the rest of the app files to the container
COPY . .

# Expose a port (if needed)
EXPOSE 3000

# Start the app
CMD [ "npm", "run", "dev" ]
