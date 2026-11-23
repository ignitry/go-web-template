require "active_record"
require "dotenv/load"
require "erb"
require "yaml"

environment = ENV.fetch("APP_ENV", "development")
database_yml = ERB.new(File.read(File.join(__dir__, "migrate/config/database.yml"))).result
config = YAML.safe_load(database_yml, aliases: true).fetch(environment)
ActiveRecord::Base.configurations = { environment => config }
ActiveRecord::Base.establish_connection(config)

namespace :db do
  task :create do
    ActiveRecord::Tasks::DatabaseTasks.create_current(environment)
  end

  task setup: [:create, :migrate]

  task :migrate do
    ActiveRecord::MigrationContext.new(File.join(__dir__, "migrate/db/migrate")).migrate
  end

  task :rollback do
    ActiveRecord::MigrationContext.new(File.join(__dir__, "migrate/db/migrate")).rollback
  end

  task :status do
    ActiveRecord::MigrationContext.new(File.join(__dir__, "migrate/db/migrate")).migrations_status.each { |row| puts row.join("  ") }
  end
end
