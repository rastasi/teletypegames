class CreateReleases < ActiveRecord::Migration[8.1]
  def change
    create_table :releases, id: { type: :bigint, unsigned: true },
                             charset: "utf8mb4", collation: "utf8mb4_0900_ai_ci",
                             if_not_exists: true do |t|
      t.string   :cartridge_path
      t.datetime :created_at,      precision: 3
      t.datetime :deleted_at,      precision: 3
      t.string   :docs_folder_path
      t.string   :html_folder_path
      t.bigint   :software_id,     unsigned: true
      t.string   :source_path
      t.datetime :updated_at,      precision: 3
      t.string   :version,         limit: 64
      t.boolean  :web_playable,    default: false

      t.index :deleted_at,  name: "idx_releases_deleted_at"
      t.index :software_id, name: "idx_releases_software_id"
    end

    add_foreign_key :releases, :softwares, name: "fk_softwares_releases" unless foreign_key_exists?(:releases, :softwares)
  end
end
