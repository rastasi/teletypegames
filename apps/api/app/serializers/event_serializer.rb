class EventSerializer
  DATE_FORMAT = "%Y-%m-%dT%H:%M:%SZ"

  def self.serialize_all(events)
    events.map { |e| { name: e.name, date: e.date.utc.strftime(DATE_FORMAT) } }
  end
end
