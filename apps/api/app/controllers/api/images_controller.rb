class Api::ImagesController < ApplicationController
  def show
    image = Image.find(params[:id])
    send_file image.file_path, type: image.content_type, disposition: "inline"
  rescue ActiveRecord::RecordNotFound
    render plain: "Not Found", status: :not_found
  rescue Errno::ENOENT
    render plain: "Not Found", status: :not_found
  end
end
