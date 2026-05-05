class SoftwareListService
  def call
    softwares = Software.includes(:releases, :external_links).all
    items = softwares.map do |sw|
      SoftwareSerializer.show_data(sw, sw.releases.to_a)
    end
    { softwares: items }
  end
end
