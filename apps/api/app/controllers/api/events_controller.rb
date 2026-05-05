class Api::EventsController < ApiController
  def index
    events = Event.upcoming.map do |e|
      { name: e.name, date: e.date.utc.strftime("%Y-%m-%dT%H:%M:%SZ") }
    end
    render json: events
  end
end
