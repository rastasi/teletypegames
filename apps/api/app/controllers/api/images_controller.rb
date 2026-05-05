class Api::ImagesController < ApplicationController
  skip_before_action :verify_authenticity_token

  def show
    image = Image.find(params[:id])
    path  = image.file_path

    return render plain: "Not Found", status: :not_found unless File.exist?(path)

    send_file path, type: image.content_type, disposition: "inline"
  rescue ActiveRecord::RecordNotFound
    render plain: "Not Found", status: :not_found
  end
end
