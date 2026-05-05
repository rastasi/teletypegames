class Api::SoftwareHighlightedController < ApiController
  def index
    result = SoftwareHighlightedService.new.call
    if result
      render json: result
    else
      render json: { error: "no highlighted software found" }, status: :not_found
    end
  end
end
