FROM nginx
# COPY ./frontend/dist /usr/share/nginx/html

# RUN apt-get update -qq && apt-get -y install apache2-utils
# RUN rm -f /etc/nginx/conf.d/*
COPY ./nginx.conf /etc/nginx/conf.d/default.conf

# RUN mkdir /data/
# RUN mkdir /data/www

CMD [ "nginx", "-g", "daemon off;" ]