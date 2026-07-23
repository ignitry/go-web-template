FROM ruby:3.4-alpine
RUN apk add --no-cache build-base postgresql-dev
WORKDIR /migrate
COPY Gemfile Gemfile.lock* ./
RUN bundle install
COPY . .
ENTRYPOINT ["bundle", "exec", "rake"]
CMD ["db:migrate"]
