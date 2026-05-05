# This file is auto-generated from the current state of the database. Instead
# of editing this file, please use the migrations feature of Active Record to
# incrementally modify your database, and then regenerate this schema definition.
#
# This file is the source Rails uses to define your schema when running `bin/rails
# db:schema:load`. When creating a new database, `bin/rails db:schema:load` tends to
# be faster and is potentially less error prone than running all of your
# migrations from scratch. Old migrations may fail to apply correctly if those
# migrations use external dependencies or application code.
#
# It's strongly recommended that you check this file into your version control system.

ActiveRecord::Schema[8.1].define(version: 2026_05_05_000005) do
  create_table "admin_users", charset: "utf8mb4", collation: "utf8mb4_0900_ai_ci", force: :cascade do |t|
    t.datetime "created_at", null: false
    t.string "email", default: "", null: false
    t.string "encrypted_password", default: "", null: false
    t.datetime "remember_created_at"
    t.datetime "reset_password_sent_at"
    t.string "reset_password_token"
    t.datetime "updated_at", null: false
    t.index ["email"], name: "index_admin_users_on_email", unique: true
    t.index ["reset_password_token"], name: "index_admin_users_on_reset_password_token", unique: true
  end

  create_table "articles", id: :integer, charset: "utf8mb4", collation: "utf8mb4_0900_ai_ci", force: :cascade do |t|
    t.text "content"
    t.timestamp "created_at", default: -> { "CURRENT_TIMESTAMP" }
    t.string "slug"
    t.string "title"
    t.timestamp "updated_at", default: -> { "CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" }
  end

  create_table "events", id: { type: :bigint, unsigned: true }, charset: "utf8mb4", collation: "utf8mb4_0900_ai_ci", force: :cascade do |t|
    t.datetime "created_at", precision: 3, null: false
    t.datetime "date", null: false
    t.datetime "deleted_at", precision: 3
    t.string "name", null: false
    t.datetime "updated_at", precision: 3, null: false
    t.index ["date"], name: "idx_events_date"
    t.index ["deleted_at"], name: "idx_events_deleted_at"
  end

  create_table "external_links", id: { type: :bigint, unsigned: true }, charset: "utf8mb4", collation: "utf8mb4_0900_ai_ci", force: :cascade do |t|
    t.datetime "created_at", precision: 3
    t.datetime "deleted_at", precision: 3
    t.string "label", limit: 128
    t.bigint "software_id", unsigned: true
    t.datetime "updated_at", precision: 3
    t.string "url"
    t.index ["deleted_at"], name: "idx_external_links_deleted_at"
    t.index ["software_id"], name: "idx_external_links_software_id"
  end

  create_table "goadmin_menu", id: { type: :integer, unsigned: true }, charset: "utf8mb4", collation: "utf8mb4_0900_ai_ci", force: :cascade do |t|
    t.timestamp "created_at", default: -> { "CURRENT_TIMESTAMP" }
    t.string "header", limit: 150
    t.string "icon", limit: 50, null: false
    t.integer "order", default: 0, null: false
    t.integer "parent_id", default: 0, null: false
    t.string "plugin_name", limit: 150, default: "", null: false
    t.string "title", limit: 50, null: false
    t.integer "type", default: 0, null: false
    t.timestamp "updated_at", default: -> { "CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" }
    t.string "uri", null: false
    t.string "uuid", limit: 150
  end

  create_table "goadmin_operation_log", id: { type: :integer, unsigned: true }, charset: "utf8mb4", collation: "utf8mb4_0900_ai_ci", force: :cascade do |t|
    t.timestamp "created_at", default: -> { "CURRENT_TIMESTAMP" }
    t.text "input", null: false
    t.string "ip", limit: 15, null: false
    t.string "method", limit: 10, null: false
    t.string "path", null: false
    t.timestamp "updated_at", default: -> { "CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" }
    t.integer "user_id", null: false
  end

  create_table "goadmin_permissions", id: { type: :integer, unsigned: true }, charset: "utf8mb4", collation: "utf8mb4_0900_ai_ci", force: :cascade do |t|
    t.timestamp "created_at", default: -> { "CURRENT_TIMESTAMP" }
    t.string "http_method"
    t.text "http_path"
    t.string "name", limit: 50, null: false
    t.string "slug", limit: 50, null: false
    t.timestamp "updated_at", default: -> { "CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" }
    t.index ["name"], name: "permissions_name_unique", unique: true
    t.index ["slug"], name: "permissions_slug_unique", unique: true
  end

  create_table "goadmin_role_menu", id: false, charset: "utf8mb4", collation: "utf8mb4_0900_ai_ci", force: :cascade do |t|
    t.timestamp "created_at", default: -> { "CURRENT_TIMESTAMP" }
    t.integer "menu_id", null: false
    t.integer "role_id", null: false
    t.timestamp "updated_at", default: -> { "CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" }
  end

  create_table "goadmin_role_permissions", id: false, charset: "utf8mb4", collation: "utf8mb4_0900_ai_ci", force: :cascade do |t|
    t.timestamp "created_at", default: -> { "CURRENT_TIMESTAMP" }
    t.integer "permission_id", null: false
    t.integer "role_id", null: false
    t.timestamp "updated_at", default: -> { "CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" }
  end

  create_table "goadmin_role_users", id: false, charset: "utf8mb4", collation: "utf8mb4_0900_ai_ci", force: :cascade do |t|
    t.timestamp "created_at", default: -> { "CURRENT_TIMESTAMP" }
    t.integer "role_id", null: false
    t.timestamp "updated_at", default: -> { "CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" }
    t.integer "user_id", null: false
  end

  create_table "goadmin_roles", id: { type: :integer, unsigned: true }, charset: "utf8mb4", collation: "utf8mb4_0900_ai_ci", force: :cascade do |t|
    t.timestamp "created_at", default: -> { "CURRENT_TIMESTAMP" }
    t.string "name", limit: 100, null: false
    t.string "slug", limit: 100, null: false
    t.timestamp "updated_at", default: -> { "CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" }
    t.index ["name"], name: "roles_name_unique", unique: true
    t.index ["slug"], name: "roles_slug_unique", unique: true
  end

  create_table "goadmin_site", id: { type: :integer, unsigned: true }, charset: "utf8mb4", collation: "utf8mb4_0900_ai_ci", force: :cascade do |t|
    t.timestamp "created_at", default: -> { "CURRENT_TIMESTAMP" }
    t.string "description", limit: 3000, null: false
    t.string "key_", limit: 100, null: false
    t.integer "state", limit: 1, default: 0, null: false, unsigned: true
    t.timestamp "updated_at", default: -> { "CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" }
    t.text "value", null: false
    t.index ["key_"], name: "site_key_unique", unique: true
  end

  create_table "goadmin_user_permissions", id: false, charset: "utf8mb4", collation: "utf8mb4_0900_ai_ci", force: :cascade do |t|
    t.timestamp "created_at", default: -> { "CURRENT_TIMESTAMP" }
    t.integer "permission_id", null: false
    t.timestamp "updated_at", default: -> { "CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" }
    t.integer "user_id", null: false
  end

  create_table "goadmin_users", id: { type: :integer, unsigned: true }, charset: "utf8mb4", collation: "utf8mb4_0900_ai_ci", force: :cascade do |t|
    t.string "avatar"
    t.timestamp "created_at", default: -> { "CURRENT_TIMESTAMP" }
    t.string "name", limit: 100, null: false
    t.string "password", limit: 100, null: false
    t.string "remember_token", limit: 100
    t.timestamp "updated_at", default: -> { "CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" }
    t.string "username", limit: 100, null: false
    t.index ["username"], name: "users_username_unique", unique: true
  end

  create_table "releases", id: { type: :bigint, unsigned: true }, charset: "utf8mb4", collation: "utf8mb4_0900_ai_ci", force: :cascade do |t|
    t.string "cartridge_path"
    t.datetime "created_at", precision: 3
    t.datetime "deleted_at", precision: 3
    t.string "docs_folder_path"
    t.string "html_folder_path"
    t.bigint "software_id", unsigned: true
    t.string "source_path"
    t.datetime "updated_at", precision: 3
    t.string "version", limit: 64
    t.boolean "web_playable", default: false
    t.index ["deleted_at"], name: "idx_releases_deleted_at"
    t.index ["software_id"], name: "idx_releases_software_id"
  end

  create_table "softwares", id: { type: :bigint, unsigned: true }, charset: "utf8mb4", collation: "utf8mb4_0900_ai_ci", force: :cascade do |t|
    t.string "author"
    t.datetime "created_at", precision: 3
    t.datetime "deleted_at", precision: 3
    t.text "desc"
    t.boolean "highlighted", default: false
    t.string "license", limit: 128
    t.string "name", limit: 128
    t.string "platform", limit: 128
    t.string "site"
    t.string "status", limit: 20, default: "development"
    t.text "story"
    t.string "title"
    t.datetime "updated_at", precision: 3
    t.index ["deleted_at"], name: "idx_softwares_deleted_at"
    t.index ["name"], name: "idx_softwares_name", unique: true
  end

  add_foreign_key "external_links", "softwares", name: "fk_softwares_external_links"
  add_foreign_key "releases", "softwares", name: "fk_softwares_releases"
end
