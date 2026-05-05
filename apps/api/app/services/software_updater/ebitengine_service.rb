require "json"

module SoftwareUpdater
  class EbitengineService < BaseService
    def update(name, version)
      versioned = "#{name}-#{version}"
      zip_file  = "#{versioned}.html.zip"
      html_dir  = versioned

      delete_dir(html_dir) if dir_exists?(html_dir)
      create_dir(html_dir)
      unzip_file(zip_file, html_dir)

      metadata = parse_metadata(full_path("#{versioned}.metadata.json"))
      site_url = metadata.delete(:site)

      software = update_or_create_software(metadata.merge(platform: "ebitengine"))
      upsert_external_link(software.id, "Source Code", site_url) if site_url.present?

      create_release_if_not_exists(
        software_id:      software.id,
        version:          version,
        html_folder_path: full_path(html_dir)
      )
    end

    private

    def parse_metadata(path)
      raw = JSON.parse(File.read(path), symbolize_names: true)
      raw.slice(:name, :title, :author, :desc, :site, :license)
    end
  end
end
