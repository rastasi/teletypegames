class CreateExternalLinks < ActiveRecord::Migration[8.1]
  def change
    create_table :external_links, id: { type: :bigint, unsigned: true },
                                   charset: "utf8mb4", collation: "utf8mb4_0900_ai_ci",
                                   if_not_exists: true do |t|
      t.datetime :created_at,  precision: 3
      t.datetime :deleted_at,  precision: 3
      t.string   :label,       limit: 128
      t.bigint   :software_id, unsigned: true
      t.datetime :updated_at,  precision: 3
      t.string   :url

      t.index :deleted_at,  name: "idx_external_links_deleted_at"
      t.index :software_id, name: "idx_external_links_software_id"
    end

    add_foreign_key :external_links, :softwares, name: "fk_softwares_external_links" unless foreign_key_exists?(:external_links, :softwares)
  end
end
