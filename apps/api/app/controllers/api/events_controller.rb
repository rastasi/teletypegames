class Api::EventsController < ApiController
  def index
    render json: EventSerializer.serialize_all(Event.upcoming)
  end
end
