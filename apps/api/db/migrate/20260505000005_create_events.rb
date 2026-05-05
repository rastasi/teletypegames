class CreateEvents < ActiveRecord::Migration[8.1]
  def change
    create_table :events, id: { type: :bigint, unsigned: true },
                           charset: "utf8mb4", collation: "utf8mb4_0900_ai_ci",
                           if_not_exists: true do |t|
      t.string   :name,       null: false
      t.datetime :date,       null: false
      t.datetime :deleted_at, precision: 3
      t.timestamps            precision: 3

      t.index :deleted_at, name: "idx_events_deleted_at"
      t.index :date,       name: "idx_events_date"
    end
  end
end
