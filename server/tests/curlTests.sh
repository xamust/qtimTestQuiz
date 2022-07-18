#!/bin/sh

echo '{"str":"hello world","char":"o"}'
curl -X POST -H 'Content-Type: application/json' --data '{"str":"hello world","char":"o"}' \
      http://localhost:8080/detect && echo
echo \

echo '{"str":"Вася полетел на Луну","char":"л"}'
curl -X POST -H 'Content-Type: application/json' --data '{"str":"Вася полетел на Луну","char":"л"}' \
      http://localhost:8080/detect && echo
echo \

echo '{"str":"Вас3я полет3ел на Луну","char":"л"}'
curl -X POST -H 'Content-Type: application/json' --data '{"str":"Вас3я полет3ел на Луну","char":"л"}' \
      http://localhost:8080/detect && echo
echo \

echo '{"str":"Вася полетел на Луну","char":"ла"}'
curl -X POST -H 'Content-Type: application/json' --data '{"str":"Вася полетел на Луну","char":"ла"}' \
      http://localhost:8080/detect && echo