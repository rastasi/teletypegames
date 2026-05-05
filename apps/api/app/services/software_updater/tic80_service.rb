module SoftwareUpdater
  class Tic80Service < BaseService
    def update(name, version)
      versioned     = "#{name}-#{version}"
      zip_file      = "#{versioned}.html.zip"
      docs_zip_file = "#{versioned}-docs.zip"
      html_dir      = versioned
      docs_dir      = "#{versioned}-docs"

      delete_dir(html_dir) if dir_exists?(html_dir)
      delete_dir(docs_dir) if dir_exists?(docs_dir)

      create_dir(html_dir)
      unzip_file(zip_file, html_dir)

      create_dir(docs_dir)
      unzip_file(docs_zip_file, docs_dir)

      metadata = parse_lua_metadata(full_path("#{versioned}.lua"))
      site_url = metadata.delete(:site)

      software = update_or_create_software(metadata.merge(platform: "tic80"))
      upsert_external_link(software.id, "Source Code", site_url) if site_url.present?

      create_release_if_not_exists(
        software_id:     software.id,
        version:         version,
        cartridge_path:  full_path("#{versioned}.tic"),
        source_path:     full_path("#{versioned}.lua"),
        html_folder_path: full_path(html_dir),
        docs_folder_path: full_path(docs_dir)
      )
    end

    private

    def parse_lua_metadata(source_path)
      metadata = {}
      File.foreach(source_path) do |line|
        break unless line.start_with?("--")
        parts = line[2..].split(":", 2)
        next if parts.length != 2
        key   = parts[0].strip.downcase.to_sym
        value = parts[1].strip
        metadata[key] = value
      end
      metadata
    end
  end
end
