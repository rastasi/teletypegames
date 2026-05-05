AdminUser.find_or_create_by!(email: ENV.fetch("ADMIN_EMAIL", "admin@teletype.hu")) do |user|
  user.password              = ENV.fetch("ADMIN_PASSWORD", "password123")
  user.password_confirmation = ENV.fetch("ADMIN_PASSWORD", "password123")
end
