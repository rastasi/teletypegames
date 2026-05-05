class Member < ApplicationRecord
  def self.ransackable_attributes(auth_object = nil)
    %w[avatar_filename created_at id motto nick real_nick updated_at]
  end
end
