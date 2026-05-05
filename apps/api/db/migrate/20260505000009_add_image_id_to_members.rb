class AddImageIdToMembers < ActiveRecord::Migration[7.1]
  def change
    add_reference :members, :image, null: true, foreign_key: true
  end
end
