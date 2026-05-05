class Image < ApplicationRecord
  UPLOAD_PATH = ENV.fetch("IMAGE_CONTAINER_PATH", "/images")

  attr_accessor :file_upload

  before_save :process_upload, if: -> { file_upload.present? }

  def self.ransackable_attributes(auth_object = nil)
    %w[content_type created_at filename id original_filename updated_at]
  end

  def file_path
    File.join(UPLOAD_PATH, filename.to_s)
  end

  private

  def process_upload
    FileUtils.mkdir_p(UPLOAD_PATH)
    self.original_filename = file_upload.original_filename
    self.content_type      = file_upload.content_type.presence || "application/octet-stream"
    ext                    = File.extname(file_upload.original_filename)
    self.filename          = "#{SecureRandom.uuid}#{ext}"
    IO.copy_stream(file_upload.to_io, file_path)
  end
end
