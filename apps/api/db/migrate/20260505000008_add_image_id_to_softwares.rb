class AddImageIdToSoftwares < ActiveRecord::Migration[7.1]
  def change
    add_reference :softwares, :image, null: true, foreign_key: true
  end
end
