module Api
  class SoftwareController < ApiController
    def index
      result = SoftwareListService.new.call
      render json: result
    end
  end
end
