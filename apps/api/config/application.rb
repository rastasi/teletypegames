require_relative "boot"
require "rails"
require "active_model/railtie"
require "active_record/railtie"
require "action_controller/railtie"
require "action_view/railtie"
require "action_dispatch/railtie"
require "sprockets/railtie"

Bundler.require(*Rails.groups)

module Api
  class Application < Rails::Application
    config.load_defaults 8.0

    config.time_zone = "UTC"
    config.active_record.default_timezone = :utc

    config.assets.prefix = "/admin/assets"

    config.autoload_lib(ignore: %w[assets tasks])
  end
end
