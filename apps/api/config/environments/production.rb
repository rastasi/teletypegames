Rails.application.configure do
config.eager_load = true
  config.consider_all_requests_local = false
  config.log_level = :info
  config.log_tags = [:request_id]
  config.active_record.verbose_query_logs = false
end
