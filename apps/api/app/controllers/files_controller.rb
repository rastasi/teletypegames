class FilesController < ApplicationController
  skip_forgery_protection

  BASE_PATH = Pathname.new(ENV.fetch("FILE_CONTAINER_PATH", "/softwares")).realpath
  def show
    requested = params[:path].to_s
    full_path = BASE_PATH.join(requested)

    if File.directory?(full_path)
      index_path = full_path.join('index.html')
      return head(:not_found) unless File.file?(index_path)
      return redirect_to("/file/#{requested.chomp('/')}/index.html", status: :moved_permanently)
    end

    if File.file?(full_path)
      send_file full_path, disposition: 'inline'
    else
      head :not_found
    end  
  end
end
