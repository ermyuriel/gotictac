# base image
FROM mhart/alpine-node:11
WORKDIR /
ENV PATH /node_modules/.bin:$PATH
COPY . /
RUN npm install --silent
EXPOSE 3000
CMD ["npm", "start"]