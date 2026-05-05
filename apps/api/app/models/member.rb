class Member < ApplicationRecord
  belongs_to :image, optional: true

  def self.ransackable_attributes(auth_object = nil)
    %w[avatar_filename created_at id image_id motto nick real_nick updated_at]
  end

  def self.ransackable_associations(auth_object = nil)
    %w[image]
  end
end
