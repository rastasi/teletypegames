class CreateAdminUsers < ActiveRecord::Migration[8.1]
  def change
    create_table :admin_users, charset: "utf8mb4", collation: "utf8mb4_0900_ai_ci",
                                if_not_exists: true do |t|
      t.string   :email,                  null: false, default: ""
      t.string   :encrypted_password,     null: false, default: ""
      t.string   :reset_password_token
      t.datetime :reset_password_sent_at
      t.datetime :remember_created_at
      t.datetime :created_at,             null: false
      t.datetime :updated_at,             null: false

      t.index :email,                name: "index_admin_users_on_email", unique: true
      t.index :reset_password_token, name: "index_admin_users_on_reset_password_token", unique: true
    end
  end
end
