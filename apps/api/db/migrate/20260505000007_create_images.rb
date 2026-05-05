class CreateImages < ActiveRecord::Migration[7.1]
  def change
    create_table :images do |t|
      t.string :filename,          null: false
      t.string :original_filename, null: false
      t.string :content_type,      null: false, default: "application/octet-stream"

      t.timestamps
    end
  end
end
