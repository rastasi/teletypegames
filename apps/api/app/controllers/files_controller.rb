class FilesController < ApplicationController
  def show
    base = Pathname.new(ENV.fetch("FILE_CONTAINER_PATH", "/softwares")).realpath
    full = base.join(params[:path])

    return render plain: "Not Found", status: :not_found unless full.to_s.start_with?(base.to_s)
    return render plain: "Not Found", status: :not_found unless full.exist? && full.file?

    send_file full.to_s, disposition: "inline"
  rescue Errno::ENOENT
    render plain: "Not Found", status: :not_found
  end
end
