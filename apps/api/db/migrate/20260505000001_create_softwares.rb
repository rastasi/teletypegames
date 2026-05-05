class CreateSoftwares < ActiveRecord::Migration[8.1]
  def change
    create_table :softwares, id: { type: :bigint, unsigned: true },
                              charset: "utf8mb4", collation: "utf8mb4_0900_ai_ci",
                              if_not_exists: true do |t|
      t.string   :author
      t.datetime :created_at, precision: 3
      t.datetime :deleted_at, precision: 3
      t.text     :desc
      t.boolean  :highlighted, default: false
      t.string   :license,     limit: 128
      t.string   :name,        limit: 128
      t.string   :platform,    limit: 128
      t.string   :site
      t.string   :status,      limit: 20, default: "development"
      t.text     :story
      t.string   :title
      t.datetime :updated_at,  precision: 3

      t.index :deleted_at, name: "idx_softwares_deleted_at"
      t.index :name,       name: "idx_softwares_name", unique: true
    end
  end
end
