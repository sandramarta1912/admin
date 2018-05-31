FROM iron/go

WORKDIR /app

CMD mkdir /app/tpl
CMD mkdir /app/static

ADD tpl/*  /app/tpl/
ADD static/* /app/static/
ADD admin /app/


ENTRYPOINT ["./admin"]

