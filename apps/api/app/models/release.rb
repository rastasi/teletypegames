class Release < ApplicationRecord
  self.table_name = "releases"

  belongs_to :software

  default_scope { where(deleted_at: nil) }

  def self.ransackable_attributes(auth_object = nil)
    %w[cartridge_path created_at deleted_at docs_folder_path html_folder_path id software_id source_path updated_at version web_playable]
  end

  def self.ransackable_associations(auth_object = nil)
    %w[software]
  end
end
