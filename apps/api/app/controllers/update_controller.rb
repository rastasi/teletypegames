class UpdateController < ApiController
  PLATFORM_SERVICES = {
    "tic80"      => "SoftwareUpdater::Tic80Service",
    "ebitengine" => "SoftwareUpdater::EbitengineService",
    "love"       => "SoftwareUpdater::LoveService"
  }.freeze

  def update
    return render plain: "Unauthorized", status: :unauthorized unless authorized?

    platform = params[:platform]
    name     = params[:name]
    version  = params[:version]

    return render plain: "Version not provided", status: :bad_request if version.blank?

    service_class_name = PLATFORM_SERVICES[platform]
    return render plain: "Unsupported platform: #{platform}", status: :bad_request unless service_class_name

    service_class_name.constantize.new.update(name, version)
    render plain: "Updated"
  rescue => e
    Rails.logger.error("[UpdateController] #{e.class}: #{e.message}\n#{e.backtrace.first(5).join("\n")}")
    render plain: e.message, status: :internal_server_error
  end

  private

  def authorized?
    params[:secret] == ENV.fetch("UPDATE_SECRET", "")
  end
end
