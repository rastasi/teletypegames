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

    config.hosts << ENV["WEBAPP_DOMAIN"]           if ENV["WEBAPP_DOMAIN"].present?
    config.hosts << ENV["WEBAPP_TECHNICAL_DOMAIN"] if ENV["WEBAPP_TECHNICAL_DOMAIN"].present?
    config.hosts << "teletypegames.org"

    # Trust Traefik and other Docker-internal proxies so Rails correctly reads
    # X-Forwarded-Proto: https — required for CSRF origin check behind SSL termination
    config.action_dispatch.trusted_proxies =
      ActionDispatch::RemoteIp::TRUSTED_PROXIES +
      [IPAddr.new("172.16.0.0/12"), IPAddr.new("10.0.0.0/8"), IPAddr.new("192.168.0.0/16")]

    config.autoload_lib(ignore: %w[assets tasks])
  end
end
