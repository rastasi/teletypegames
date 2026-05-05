class ExternalLink < ApplicationRecord
  self.table_name = "external_links"

  belongs_to :software

  default_scope { where(deleted_at: nil) }

  def self.ransackable_attributes(auth_object = nil)
    %w[created_at deleted_at id label software_id updated_at url]
  end

  def self.ransackable_associations(auth_object = nil)
    %w[software]
  end
end
