class Api::SoftwareController < ApiController
  def index
    render json: SoftwareListService.new.call
  end
end
