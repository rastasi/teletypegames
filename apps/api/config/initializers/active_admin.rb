ActiveAdmin.setup do |config|
  config.site_title = "Teletype Games Admin"

  config.authentication_method = :authenticate_admin_user!
  config.current_user_method   = :current_admin_user
  config.logout_link_path      = :destroy_admin_user_session_path
  config.logout_link_method    = :delete

  config.comments = false
  config.batch_actions = true
  config.filter_attributes = [:password, :password_confirmation, :current_password, :encrypted_password]

  config.localize_format = :long
end
