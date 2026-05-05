class SoftwareShowService
  def initialize(name)
    @name = name
  end

  def call
    software = Software.includes(:external_links).find_by!(name: @name)
    releases = Release.where(software_id: software.id).to_a
    SoftwareSerializer.show_data(software, releases)
  end
end
