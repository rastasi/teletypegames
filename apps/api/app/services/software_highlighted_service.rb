class SoftwareHighlightedService
  def call
    software = Software.includes(:external_links)
                       .where(highlighted: true)
                       .order(id: :desc)
                       .first
    return nil unless software

    releases = Release.where(software_id: software.id).to_a
    SoftwareSerializer.show_data(software, releases)
  end
end
