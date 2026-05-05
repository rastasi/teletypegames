#!/bin/sh
set -e

mkdir -p tmp/pids tmp/cache tmp/sockets log
rm -f tmp/pids/server.pid

bundle check || bundle install

bundle exec rails db:migrate
bundle exec rails db:seed

bundle exec rails server -b 0.0.0.0 -p 3000
