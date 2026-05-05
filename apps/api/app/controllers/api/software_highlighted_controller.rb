module Api
  class SoftwareHighlightedController < ApiController
    def index
      result = SoftwareHighlightedService.new.call
      if result.nil?
        render json: { error: "no highlighted software found" }, status: :not_found
      else
        render json: result
      end
    end
  end
end
