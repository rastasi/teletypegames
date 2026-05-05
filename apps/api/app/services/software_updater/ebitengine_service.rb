module SoftwareUpdater
  class EbitengineService < BaseService
    def update(name, version)
      versioned = "#{name}-#{version}"
      extract_zip_to_dir("#{versioned}.html.zip", versioned)

      metadata = parse_json_metadata(full_path("#{versioned}.metadata.json"))
      site_url = metadata.delete(:site)

      software = update_or_create_software(metadata.merge(platform: "ebitengine"))
      upsert_external_link(software.id, "Source Code", site_url) if site_url.present?

      create_release_if_not_exists(
        software_id:      software.id,
        version:          version,
        html_folder_path: full_path(versioned)
      )
    end
  end
end
