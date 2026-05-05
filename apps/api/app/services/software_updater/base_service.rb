require "zip"
require "json"
require "fileutils"

module SoftwareUpdater
  class BaseService
    def initialize
      @base_path = ENV.fetch("FILE_CONTAINER_PATH", "/softwares")
    end

    private

    def full_path(filename)
      File.join(@base_path, filename)
    end

    def dir_exists?(dirname)
      File.directory?(full_path(dirname))
    end

    def delete_dir(dirname)
      FileUtils.rm_rf(full_path(dirname))
    end

    def create_dir(dirname)
      FileUtils.mkdir_p(full_path(dirname))
    end

    def unzip_file(zip_filename, dest_dirname)
      dest = full_path(dest_dirname)
      Zip::File.open(full_path(zip_filename)) do |zip|
        zip.each do |entry|
          entry_dest = File.join(dest, entry.name)
          FileUtils.mkdir_p(File.dirname(entry_dest))
          entry.extract(entry_dest) { true }
        end
      end
    end

    def extract_zip_to_dir(zip_file, dir_name)
      delete_dir(dir_name) if dir_exists?(dir_name)
      create_dir(dir_name)
      unzip_file(zip_file, dir_name)
    end

    def parse_json_metadata(path)
      raw = JSON.parse(File.read(path), symbolize_names: true)
      raw.slice(:name, :title, :author, :desc, :site, :license)
    end

    def update_or_create_software(attrs)
      software = Software.unscoped.find_or_initialize_by(name: attrs[:name])
      software.assign_attributes(attrs.except(:name))
      software.deleted_at = nil
      software.save!
      software
    end

    def upsert_external_link(software_id, label, url)
      link = ExternalLink.unscoped.find_or_initialize_by(software_id: software_id, label: label)
      link.url = url
      link.deleted_at = nil
      link.save!
    end

    def create_release_if_not_exists(attrs)
      existing = Release.unscoped.find_by(software_id: attrs[:software_id], version: attrs[:version])
      return existing if existing

      Release.create!(attrs)
    end
  end
end
