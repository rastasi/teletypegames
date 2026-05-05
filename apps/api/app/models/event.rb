class Event < ApplicationRecord
  default_scope { where(deleted_at: nil) }

  scope :upcoming, -> { where("date > ?", Time.current).order(:date) }

  def self.ransackable_attributes(auth_object = nil)
    %w[id name date created_at updated_at]
  end
end
