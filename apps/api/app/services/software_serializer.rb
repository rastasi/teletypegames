class SoftwareSerializer
  GO_ZERO_TIME   = "0001-01-01T00:00:00Z"
  FILE_PATH_FROM = "/softwares/"
  FILE_PATH_TO   = "/file/"

  def self.show_data(software, releases)
    sorted = releases.sort_by { |r| r.created_at || Time.at(0) }.reverse

    latest      = sorted.reject { |r| r.version.to_s.start_with?("dev-") }.first
    web_playable = sorted.find   { |r| r.html_folder_path.present? }

    {
      software:           serialize_software(software),
      releases:           sorted.map { |r| serialize_release(r) },
      latestRelease:      latest       ? serialize_release(latest)       : nil,
      webPlayableRelease: web_playable ? serialize_release(web_playable) : nil
    }
  end

  def self.serialize_software(sw)
    {
      "ID"           => sw.id,
      "CreatedAt"    => fmt(sw.created_at),
      "UpdatedAt"    => fmt(sw.updated_at),
      "DeletedAt"    => fmt_nullable(sw.deleted_at),
      "name"         => sw.name,
      "title"        => sw.title,
      "author"       => sw.author,
      "desc"         => sw.desc.to_s,
      "story"        => sw.story.to_s,
      "license"      => sw.license.to_s,
      "platform"     => sw.platform,
      "status"       => sw.status,
      "highlighted"  => sw.highlighted ? true : false,
      "externalLinks" => sw.external_links.map { |el| serialize_external_link(el) }
    }
  end

  def self.serialize_release(r)
    {
      "ID"             => r.id,
      "CreatedAt"      => fmt(r.created_at),
      "UpdatedAt"      => fmt(r.updated_at),
      "DeletedAt"      => fmt_nullable(r.deleted_at),
      "softwareId"     => r.software_id,
      "version"        => r.version,
      "cartridgePath"  => rewrite(r.cartridge_path),
      "sourcePath"     => rewrite(r.source_path),
      "htmlFolderPath" => rewrite(r.html_folder_path),
      "docsFolderPath" => rewrite(r.docs_folder_path)
    }
  end

  def self.serialize_external_link(el)
    {
      "ID"         => el.id,
      "CreatedAt"  => fmt(el.created_at),
      "UpdatedAt"  => fmt(el.updated_at),
      "DeletedAt"  => fmt_nullable(el.deleted_at),
      "softwareId" => el.software_id,
      "label"      => el.label,
      "url"        => el.url
    }
  end

  # null → Go zero-time string (matches GORM behavior)
  def self.fmt(t)
    return GO_ZERO_TIME if t.nil?
    t.utc.strftime("%Y-%m-%dT%H:%M:%S.%3NZ")
  end

  # null → JSON null (matches gorm.DeletedAt behavior)
  def self.fmt_nullable(t)
    return nil if t.nil?
    t.utc.strftime("%Y-%m-%dT%H:%M:%S.%3NZ")
  end

  def self.rewrite(path)
    return "" if path.blank?
    path.gsub(FILE_PATH_FROM, FILE_PATH_TO)
  end
end
