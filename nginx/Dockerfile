FROM nginx:stable

RUN rm -f /etc/nginx/conf.d/default.conf
COPY nginx.conf /etc/nginx/
COPY conf.d/default.conf /etc/nginx/conf.d/


COPY index.html /usr/share/nginx/html/
COPY 404.html /usr/share/nginx/html/
COPY 404.html /usr/local/nginx/html/
COPY 404.html /var/www/html/

COPY 50x.html /usr/share/nginx/html/

