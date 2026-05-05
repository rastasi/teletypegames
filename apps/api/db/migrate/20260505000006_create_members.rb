class CreateMembers < ActiveRecord::Migration[7.1]
  def change
    create_table :members do |t|
      t.string :nick,            null: false
      t.string :real_nick,       null: false
      t.text   :motto
      t.string :avatar_filename

      t.timestamps
    end
  end
end
